package signer

import (
	"context"
	"fmt"

	"github.com/strangelove-ventures/horcrux/pkg/nodes"

	"github.com/strangelove-ventures/horcrux/pkg/types"

	"github.com/google/uuid"
	"github.com/hashicorp/raft"
	"github.com/strangelove-ventures/horcrux/signer/proto"
)

var _ proto.CosignerServer = &CosignerGRPCServer{}

type CosignerGRPCServer struct {
	cosigner           *nodes.LocalCosigner
	thresholdValidator *ThresholdValidator
	raftStore          *RaftStore
	// TODO: add logger and not rely on raftStore.logger
	proto.UnimplementedCosignerServer
}

func NewCosignerGRPCServer(
	cosigner *nodes.LocalCosigner,
	thresholdValidator *ThresholdValidator,
	raftStore *RaftStore,
) *CosignerGRPCServer {
	return &CosignerGRPCServer{
		cosigner:           cosigner,
		thresholdValidator: thresholdValidator,
		raftStore:          raftStore,
	}
}

func (rpc *CosignerGRPCServer) SignBlock(
	ctx context.Context,
	req *proto.SignBlockRequest,
) (*proto.SignBlockResponse, error) {
	res, _, err := rpc.thresholdValidator.Sign(ctx, req.ChainID, types.BlockFromProto(req.Block))
	if err != nil {
		return nil, err
	}
	return &proto.SignBlockResponse{
		Signature: res,
	}, nil
}

func (rpc *CosignerGRPCServer) SetNoncesAndSign(
	ctx context.Context,
	req *proto.SetNoncesAndSignRequest,
) (*proto.SetNoncesAndSignResponse, error) {
	res, err := rpc.cosigner.SetNoncesAndSign(ctx, nodes.CosignerSetNoncesAndSignRequest{
		ChainID: req.ChainID,
		Nonces: &nodes.CosignerUUIDNonces{
			UUID:   uuid.UUID(req.Uuid),
			Nonces: nodes.CosignerNoncesFromProto(req.GetNonces()),
		},
		HRST:      types.HRSTKeyFromProto(req.GetHrst()),
		SignBytes: req.GetSignBytes(),
	})
	if err != nil {
		rpc.raftStore.logger.Error(
			"Failed to sign with shard",
			"chain_id", req.ChainID,
			"height", req.Hrst.Height,
			"round", req.Hrst.Round,
			"step", req.Hrst.Step,
			"error", err,
		)
		return nil, err
	}
	rpc.raftStore.logger.Info(
		"Signed with shard",
		"chain_id", req.ChainID,
		"height", req.Hrst.Height,
		"round", req.Hrst.Round,
		"step", req.Hrst.Step,
	)
	return &proto.SetNoncesAndSignResponse{
		NoncePublic: res.NoncePublic,
		Timestamp:   res.Timestamp.UnixNano(),
		Signature:   res.Signature,
	}, nil
}

func (rpc *CosignerGRPCServer) GetNonces(
	ctx context.Context,
	req *proto.GetNoncesRequest,
) (*proto.GetNoncesResponse, error) {
	uuids := make([]uuid.UUID, len(req.Uuids))
	for i, uuidBytes := range req.Uuids {
		uuids[i] = uuid.UUID(uuidBytes)
	}
	res, err := rpc.cosigner.GetNonces(
		ctx,
		uuids,
	)
	if err != nil {
		return nil, err
	}

	return &proto.GetNoncesResponse{
		Nonces: res.ToProto(),
	}, nil
}

// TODO: // TransferLeadership should not be a CosignerGRPCServer method?
func (rpc *CosignerGRPCServer) TransferLeadership(
	_ context.Context,
	req *proto.TransferLeadershipRequest,
) (*proto.TransferLeadershipResponse, error) {
	if rpc.raftStore.raft.State() != raft.Leader {
		return &proto.TransferLeadershipResponse{}, nil
	}
	leaderID := req.GetLeaderID()
	if leaderID != "" {
		for _, c := range rpc.raftStore.Cosigners {
			shardID := fmt.Sprint(c.GetIndex())
			if shardID == leaderID {
				raftAddress := p2pURLToRaftAddress(c.GetAddress())
				fmt.Printf("Transferring leadership to Index: %s - Address: %s\n", shardID, raftAddress)
				rpc.raftStore.raft.LeadershipTransferToServer(raft.ServerID(shardID), raft.ServerAddress(raftAddress))
				return &proto.TransferLeadershipResponse{LeaderID: shardID, LeaderAddress: raftAddress}, nil
			}
		}
	}
	fmt.Printf("Transferring leadership to next candidate\n")
	rpc.raftStore.raft.LeadershipTransfer()
	return &proto.TransferLeadershipResponse{}, nil
}

func (rpc *CosignerGRPCServer) GetLeader(
	context.Context,
	*proto.GetLeaderRequest,
) (*proto.GetLeaderResponse, error) {
	leader := rpc.raftStore.GetLeader()
	return &proto.GetLeaderResponse{Leader: int32(leader)}, nil
}

func (rpc *CosignerGRPCServer) Ping(context.Context, *proto.PingRequest) (*proto.PingResponse, error) {
	return &proto.PingResponse{}, nil
}
