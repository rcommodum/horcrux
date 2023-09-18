// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.13.0
// source: raft_service/raft.proto

package raftService

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height    int64  `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round     int64  `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Step      int32  `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	SignBytes []byte `protobuf:"bytes,4,opt,name=signBytes,proto3" json:"signBytes,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_raft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_raft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_raft_service_raft_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Block) GetRound() int64 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *Block) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *Block) GetSignBytes() []byte {
	if x != nil {
		return x.SignBytes
	}
	return nil
}

func (x *Block) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type RaftGRPCSignBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainID string `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	Block   *Block `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *RaftGRPCSignBlockRequest) Reset() {
	*x = RaftGRPCSignBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_raft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftGRPCSignBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftGRPCSignBlockRequest) ProtoMessage() {}

func (x *RaftGRPCSignBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_raft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftGRPCSignBlockRequest.ProtoReflect.Descriptor instead.
func (*RaftGRPCSignBlockRequest) Descriptor() ([]byte, []int) {
	return file_raft_service_raft_proto_rawDescGZIP(), []int{1}
}

func (x *RaftGRPCSignBlockRequest) GetChainID() string {
	if x != nil {
		return x.ChainID
	}
	return ""
}

func (x *RaftGRPCSignBlockRequest) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

type RaftGRPCSignBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *RaftGRPCSignBlockResponse) Reset() {
	*x = RaftGRPCSignBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_raft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftGRPCSignBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftGRPCSignBlockResponse) ProtoMessage() {}

func (x *RaftGRPCSignBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_raft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftGRPCSignBlockResponse.ProtoReflect.Descriptor instead.
func (*RaftGRPCSignBlockResponse) Descriptor() ([]byte, []int) {
	return file_raft_service_raft_proto_rawDescGZIP(), []int{2}
}

func (x *RaftGRPCSignBlockResponse) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type RaftGRPCTransferLeadershipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaderID string `protobuf:"bytes,1,opt,name=leaderID,proto3" json:"leaderID,omitempty"`
}

func (x *RaftGRPCTransferLeadershipRequest) Reset() {
	*x = RaftGRPCTransferLeadershipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_raft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftGRPCTransferLeadershipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftGRPCTransferLeadershipRequest) ProtoMessage() {}

func (x *RaftGRPCTransferLeadershipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_raft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftGRPCTransferLeadershipRequest.ProtoReflect.Descriptor instead.
func (*RaftGRPCTransferLeadershipRequest) Descriptor() ([]byte, []int) {
	return file_raft_service_raft_proto_rawDescGZIP(), []int{3}
}

func (x *RaftGRPCTransferLeadershipRequest) GetLeaderID() string {
	if x != nil {
		return x.LeaderID
	}
	return ""
}

type RaftGRPCTransferLeadershipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaderID      string `protobuf:"bytes,1,opt,name=leaderID,proto3" json:"leaderID,omitempty"`
	LeaderAddress string `protobuf:"bytes,2,opt,name=leaderAddress,proto3" json:"leaderAddress,omitempty"`
}

func (x *RaftGRPCTransferLeadershipResponse) Reset() {
	*x = RaftGRPCTransferLeadershipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_raft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftGRPCTransferLeadershipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftGRPCTransferLeadershipResponse) ProtoMessage() {}

func (x *RaftGRPCTransferLeadershipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_raft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftGRPCTransferLeadershipResponse.ProtoReflect.Descriptor instead.
func (*RaftGRPCTransferLeadershipResponse) Descriptor() ([]byte, []int) {
	return file_raft_service_raft_proto_rawDescGZIP(), []int{4}
}

func (x *RaftGRPCTransferLeadershipResponse) GetLeaderID() string {
	if x != nil {
		return x.LeaderID
	}
	return ""
}

func (x *RaftGRPCTransferLeadershipResponse) GetLeaderAddress() string {
	if x != nil {
		return x.LeaderAddress
	}
	return ""
}

type RaftGRPCGetLeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RaftGRPCGetLeaderRequest) Reset() {
	*x = RaftGRPCGetLeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_raft_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftGRPCGetLeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftGRPCGetLeaderRequest) ProtoMessage() {}

func (x *RaftGRPCGetLeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_raft_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftGRPCGetLeaderRequest.ProtoReflect.Descriptor instead.
func (*RaftGRPCGetLeaderRequest) Descriptor() ([]byte, []int) {
	return file_raft_service_raft_proto_rawDescGZIP(), []int{5}
}

type RaftGRPCGetLeaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Leader string `protobuf:"bytes,1,opt,name=leader,proto3" json:"leader,omitempty"`
}

func (x *RaftGRPCGetLeaderResponse) Reset() {
	*x = RaftGRPCGetLeaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_raft_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftGRPCGetLeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftGRPCGetLeaderResponse) ProtoMessage() {}

func (x *RaftGRPCGetLeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_raft_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftGRPCGetLeaderResponse.ProtoReflect.Descriptor instead.
func (*RaftGRPCGetLeaderResponse) Descriptor() ([]byte, []int) {
	return file_raft_service_raft_proto_rawDescGZIP(), []int{6}
}

func (x *RaftGRPCGetLeaderResponse) GetLeader() string {
	if x != nil {
		return x.Leader
	}
	return ""
}

var File_raft_service_raft_proto protoreflect.FileDescriptor

var file_raft_service_raft_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72,
	0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x61, 0x66, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x74,
	0x65, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x5e,
	0x0a, 0x18, 0x52, 0x61, 0x66, 0x74, 0x47, 0x52, 0x50, 0x43, 0x53, 0x69, 0x67, 0x6e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x39,
	0x0a, 0x19, 0x52, 0x61, 0x66, 0x74, 0x47, 0x52, 0x50, 0x43, 0x53, 0x69, 0x67, 0x6e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x3f, 0x0a, 0x21, 0x52, 0x61, 0x66,
	0x74, 0x47, 0x52, 0x50, 0x43, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x66, 0x0a, 0x22, 0x52, 0x61,
	0x66, 0x74, 0x47, 0x52, 0x50, 0x43, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x61, 0x66, 0x74, 0x47, 0x52, 0x50, 0x43, 0x47, 0x65,
	0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33,
	0x0a, 0x19, 0x52, 0x61, 0x66, 0x74, 0x47, 0x52, 0x50, 0x43, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x32, 0xc0, 0x02, 0x0a, 0x09, 0x49, 0x52, 0x61, 0x66, 0x74, 0x47, 0x52, 0x50,
	0x43, 0x12, 0x5c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x25,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x61, 0x66,
	0x74, 0x47, 0x52, 0x50, 0x43, 0x53, 0x69, 0x67, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x47, 0x52, 0x50, 0x43, 0x53, 0x69, 0x67, 0x6e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x77, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x2e, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x47, 0x52, 0x50, 0x43, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x47, 0x52, 0x50, 0x43, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x47, 0x52, 0x50, 0x43, 0x47, 0x65, 0x74, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x47,
	0x52, 0x50, 0x43, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x76, 0x65,
	0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x68, 0x6f, 0x72, 0x63, 0x72, 0x75,
	0x78, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x66, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raft_service_raft_proto_rawDescOnce sync.Once
	file_raft_service_raft_proto_rawDescData = file_raft_service_raft_proto_rawDesc
)

func file_raft_service_raft_proto_rawDescGZIP() []byte {
	file_raft_service_raft_proto_rawDescOnce.Do(func() {
		file_raft_service_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_raft_service_raft_proto_rawDescData)
	})
	return file_raft_service_raft_proto_rawDescData
}

var file_raft_service_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_raft_service_raft_proto_goTypes = []interface{}{
	(*Block)(nil),                              // 0: raftService.Block
	(*RaftGRPCSignBlockRequest)(nil),           // 1: raftService.RaftGRPCSignBlockRequest
	(*RaftGRPCSignBlockResponse)(nil),          // 2: raftService.RaftGRPCSignBlockResponse
	(*RaftGRPCTransferLeadershipRequest)(nil),  // 3: raftService.RaftGRPCTransferLeadershipRequest
	(*RaftGRPCTransferLeadershipResponse)(nil), // 4: raftService.RaftGRPCTransferLeadershipResponse
	(*RaftGRPCGetLeaderRequest)(nil),           // 5: raftService.RaftGRPCGetLeaderRequest
	(*RaftGRPCGetLeaderResponse)(nil),          // 6: raftService.RaftGRPCGetLeaderResponse
}
var file_raft_service_raft_proto_depIdxs = []int32{
	0, // 0: raftService.RaftGRPCSignBlockRequest.block:type_name -> raftService.Block
	1, // 1: raftService.IRaftGRPC.SignBlock:input_type -> raftService.RaftGRPCSignBlockRequest
	3, // 2: raftService.IRaftGRPC.TransferLeadership:input_type -> raftService.RaftGRPCTransferLeadershipRequest
	5, // 3: raftService.IRaftGRPC.GetLeader:input_type -> raftService.RaftGRPCGetLeaderRequest
	2, // 4: raftService.IRaftGRPC.SignBlock:output_type -> raftService.RaftGRPCSignBlockResponse
	4, // 5: raftService.IRaftGRPC.TransferLeadership:output_type -> raftService.RaftGRPCTransferLeadershipResponse
	6, // 6: raftService.IRaftGRPC.GetLeader:output_type -> raftService.RaftGRPCGetLeaderResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_raft_service_raft_proto_init() }
func file_raft_service_raft_proto_init() {
	if File_raft_service_raft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_raft_service_raft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_service_raft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftGRPCSignBlockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_service_raft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftGRPCSignBlockResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_service_raft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftGRPCTransferLeadershipRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_service_raft_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftGRPCTransferLeadershipResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_service_raft_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftGRPCGetLeaderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_service_raft_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftGRPCGetLeaderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_raft_service_raft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raft_service_raft_proto_goTypes,
		DependencyIndexes: file_raft_service_raft_proto_depIdxs,
		MessageInfos:      file_raft_service_raft_proto_msgTypes,
	}.Build()
	File_raft_service_raft_proto = out.File
	file_raft_service_raft_proto_rawDesc = nil
	file_raft_service_raft_proto_goTypes = nil
	file_raft_service_raft_proto_depIdxs = nil
}
