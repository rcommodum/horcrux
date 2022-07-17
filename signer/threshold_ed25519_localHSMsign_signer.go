package signer

import (
	"errors"
)

// Placeholder for Local LocalHSMsignThresholdEd25519SignatureConfig
type LocalHSMsignThresholdEd25519SignatureConfig struct {
}

type LocalHSMsignThresholdEd25519Signature struct {
	// panic("Not Implemented") // TODO: Implement HSM Signer
	// UnimplementedThresholdEd25519Signature // embedding UnimplementedCosignerGRPCServer
}

func (cfg LocalHSMsignThresholdEd25519SignatureConfig) NewThresholdEd25519Signature() *LocalHSMsignThresholdEd25519Signature {
	panic("Not Implemented")
}

func (localsigner LocalHSMsignThresholdEd25519Signature) DealShares(req CosignerGetEphemeralSecretPartRequest) (HrsMetadata, error) {
	panic("Not Implemented")
	return HrsMetadata{}, errors.New("method dealShares not implemented")
}
func (localsigner LocalHSMsignThresholdEd25519Signature) SetEphemeralSecretPart(req CosignerSetEphemeralSecretPartRequest, m *LastSignStateStruct) error {
	panic("Not Implemented")
	return errors.New("method setEphemeralSecretPart not implemented")
}
func (localsigner LocalHSMsignThresholdEd25519Signature) GetEphemeralSecretPart(req CosignerGetEphemeralSecretPartRequest, m *LastSignStateStruct) (CosignerEphemeralSecretPart, error) {
	panic("Not Implemented")
	return CosignerEphemeralSecretPart{}, errors.New("method getEphemeralSecretPart")
}
func (localsigner LocalHSMsignThresholdEd25519Signature) Sign(req CosignerSignRequest, m *LastSignStateStruct) (CosignerSignResponse, error) {
	panic("Not Implemented")
	return CosignerSignResponse{}, errors.New("method sign not implemented")
}

/*

// UnimplementedThresholdEd25519Signature must be embedded to have forward compatible implementations.
type UnimplementedThresholdEd25519Signature struct {
}

func (UnimplementedThresholdEd25519Signature) dealShares(req CosignerGetEphemeralSecretPartRequest) (HrsMetadata, error) {
	return HrsMetadata{}, errors.New("method dealShares not implemented")
}
func (UnimplementedThresholdEd25519Signature) setEphemeralSecretPart(req CosignerSetEphemeralSecretPartRequest) error {
	return errors.New("method setEphemeralSecretPart not implemented")
}
func (UnimplementedThresholdEd25519Signature) getEphemeralSecretPart(req CosignerGetEphemeralSecretPartRequest) (CosignerEphemeralSecretPart, error) {
	return CosignerEphemeralSecretPart{}, errors.New("method getEphemeralSecretPart")
}
func (UnimplementedThresholdEd25519Signature) sign(req CosignerSignRequest) (CosignerSignResponse, error) {
	return CosignerSignResponse{}, errors.New("method sign not implemented")
}

*/
