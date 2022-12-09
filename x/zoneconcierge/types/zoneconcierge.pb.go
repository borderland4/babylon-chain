// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/zoneconcierge/zoneconcierge.proto

package types

import (
	fmt "fmt"
	types "github.com/babylonchain/babylon/x/checkpointing/types"
	proto "github.com/gogo/protobuf/proto"
	crypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
	_ "github.com/tendermint/tendermint/proto/tendermint/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// IndexedHeader is the metadata of a CZ header
type IndexedHeader struct {
	// chain_id is the unique ID of the chain
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// hash is the hash of this header
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// height is the height of this header on CZ ledger
	// (hash, height) jointly provides the position of the header on CZ ledger
	Height uint64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	// babylon_block_height is the height of the Babylon block that includes this header
	BabylonBlockHeight uint64 `protobuf:"varint,4,opt,name=babylon_block_height,json=babylonBlockHeight,proto3" json:"babylon_block_height,omitempty"`
	// epoch is the epoch number of this header on Babylon ledger
	BabylonEpoch uint64 `protobuf:"varint,5,opt,name=babylon_epoch,json=babylonEpoch,proto3" json:"babylon_epoch,omitempty"`
	// babylon_tx_hash is the hash of the tx that includes this header
	// (babylon_block_height, babylon_tx_hash) jointly provides the position of the header on Babylon ledger
	BabylonTxHash []byte `protobuf:"bytes,6,opt,name=babylon_tx_hash,json=babylonTxHash,proto3" json:"babylon_tx_hash,omitempty"`
}

func (m *IndexedHeader) Reset()         { *m = IndexedHeader{} }
func (m *IndexedHeader) String() string { return proto.CompactTextString(m) }
func (*IndexedHeader) ProtoMessage()    {}
func (*IndexedHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76d28ce8dde4532, []int{0}
}
func (m *IndexedHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexedHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexedHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexedHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexedHeader.Merge(m, src)
}
func (m *IndexedHeader) XXX_Size() int {
	return m.Size()
}
func (m *IndexedHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexedHeader.DiscardUnknown(m)
}

var xxx_messageInfo_IndexedHeader proto.InternalMessageInfo

func (m *IndexedHeader) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *IndexedHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *IndexedHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *IndexedHeader) GetBabylonBlockHeight() uint64 {
	if m != nil {
		return m.BabylonBlockHeight
	}
	return 0
}

func (m *IndexedHeader) GetBabylonEpoch() uint64 {
	if m != nil {
		return m.BabylonEpoch
	}
	return 0
}

func (m *IndexedHeader) GetBabylonTxHash() []byte {
	if m != nil {
		return m.BabylonTxHash
	}
	return nil
}

// Forks is a list of non-canonical `IndexedHeader`s at the same height.
// For example, assuming the following blockchain
// ```
// A <- B <- C <- D <- E
//            \ -- D1
//            \ -- D2
// ```
// Then the fork will be {[D1, D2]} where each item is in struct `IndexedBlock`.
//
// Note that each `IndexedHeader` in the fork should have a valid quorum certificate.
// Such forks exist since Babylon considers CZs might have dishonest majority.
// Also note that the IBC-Go implementation will only consider the first header in a fork valid, since
// the subsequent headers cannot be verified without knowing the validator set in the previous header.
type Forks struct {
	// blocks is the list of non-canonical indexed headers at the same height
	Headers []*IndexedHeader `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (m *Forks) Reset()         { *m = Forks{} }
func (m *Forks) String() string { return proto.CompactTextString(m) }
func (*Forks) ProtoMessage()    {}
func (*Forks) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76d28ce8dde4532, []int{1}
}
func (m *Forks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Forks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Forks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Forks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forks.Merge(m, src)
}
func (m *Forks) XXX_Size() int {
	return m.Size()
}
func (m *Forks) XXX_DiscardUnknown() {
	xxx_messageInfo_Forks.DiscardUnknown(m)
}

var xxx_messageInfo_Forks proto.InternalMessageInfo

func (m *Forks) GetHeaders() []*IndexedHeader {
	if m != nil {
		return m.Headers
	}
	return nil
}

// ChainInfo is the information of a CZ
type ChainInfo struct {
	// chain_id is the ID of the chain
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// latest_header is the latest header in the canonical chain of CZ
	LatestHeader *IndexedHeader `protobuf:"bytes,2,opt,name=latest_header,json=latestHeader,proto3" json:"latest_header,omitempty"`
	// latest_forks is the latest forks, formed as a series of IndexedHeader (from low to high)
	LatestForks *Forks `protobuf:"bytes,3,opt,name=latest_forks,json=latestForks,proto3" json:"latest_forks,omitempty"`
}

func (m *ChainInfo) Reset()         { *m = ChainInfo{} }
func (m *ChainInfo) String() string { return proto.CompactTextString(m) }
func (*ChainInfo) ProtoMessage()    {}
func (*ChainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76d28ce8dde4532, []int{2}
}
func (m *ChainInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainInfo.Merge(m, src)
}
func (m *ChainInfo) XXX_Size() int {
	return m.Size()
}
func (m *ChainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChainInfo proto.InternalMessageInfo

func (m *ChainInfo) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *ChainInfo) GetLatestHeader() *IndexedHeader {
	if m != nil {
		return m.LatestHeader
	}
	return nil
}

func (m *ChainInfo) GetLatestForks() *Forks {
	if m != nil {
		return m.LatestForks
	}
	return nil
}

// ProofEpochSealed is the proof that an epoch is sealed by the sealer header, i.e., the 2nd header of the next epoch
// With the access of metadata
// - Metadata of this epoch, which includes the sealer header
// - Raw checkpoint of this epoch
// The verifier can perform the following verification rules:
// - The raw checkpoint's `last_commit_hash` is same as in the sealer header
// - More than 1/3 (in voting power) validators in the validator set of this epoch have signed `last_commit_hash` of the sealer header
// - The epoch medatata is committed to the `app_hash` of the sealer header
// - The validator set is committed to the `app_hash` of the sealer header
type ProofEpochSealed struct {
	// validator_set is the validator set of the sealed epoch
	// This validator set has generated a BLS multisig on `last_commit_hash` of the sealer header
	ValidatorSet []*types.ValidatorWithBlsKey `protobuf:"bytes,1,rep,name=validator_set,json=validatorSet,proto3" json:"validator_set,omitempty"`
	// proof_epoch_info is the Merkle proof that the epoch's metadata is committed to `app_hash` of the sealer header
	ProofEpochInfo *crypto.ProofOps `protobuf:"bytes,2,opt,name=proof_epoch_info,json=proofEpochInfo,proto3" json:"proof_epoch_info,omitempty"`
	// proof_epoch_info is the Merkle proof that the epoch's validator set is committed to `app_hash` of the sealer header
	ProofEpochValSet *crypto.ProofOps `protobuf:"bytes,3,opt,name=proof_epoch_val_set,json=proofEpochValSet,proto3" json:"proof_epoch_val_set,omitempty"`
}

func (m *ProofEpochSealed) Reset()         { *m = ProofEpochSealed{} }
func (m *ProofEpochSealed) String() string { return proto.CompactTextString(m) }
func (*ProofEpochSealed) ProtoMessage()    {}
func (*ProofEpochSealed) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76d28ce8dde4532, []int{3}
}
func (m *ProofEpochSealed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofEpochSealed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofEpochSealed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofEpochSealed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofEpochSealed.Merge(m, src)
}
func (m *ProofEpochSealed) XXX_Size() int {
	return m.Size()
}
func (m *ProofEpochSealed) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofEpochSealed.DiscardUnknown(m)
}

var xxx_messageInfo_ProofEpochSealed proto.InternalMessageInfo

func (m *ProofEpochSealed) GetValidatorSet() []*types.ValidatorWithBlsKey {
	if m != nil {
		return m.ValidatorSet
	}
	return nil
}

func (m *ProofEpochSealed) GetProofEpochInfo() *crypto.ProofOps {
	if m != nil {
		return m.ProofEpochInfo
	}
	return nil
}

func (m *ProofEpochSealed) GetProofEpochValSet() *crypto.ProofOps {
	if m != nil {
		return m.ProofEpochValSet
	}
	return nil
}

func init() {
	proto.RegisterType((*IndexedHeader)(nil), "babylon.zoneconcierge.v1.IndexedHeader")
	proto.RegisterType((*Forks)(nil), "babylon.zoneconcierge.v1.Forks")
	proto.RegisterType((*ChainInfo)(nil), "babylon.zoneconcierge.v1.ChainInfo")
	proto.RegisterType((*ProofEpochSealed)(nil), "babylon.zoneconcierge.v1.ProofEpochSealed")
}

func init() {
	proto.RegisterFile("babylon/zoneconcierge/zoneconcierge.proto", fileDescriptor_c76d28ce8dde4532)
}

var fileDescriptor_c76d28ce8dde4532 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0xad, 0x69, 0xd7, 0x31, 0xb7, 0x85, 0xca, 0x20, 0x14, 0x06, 0x84, 0xaa, 0x93, 0x46, 0x79,
	0x20, 0x81, 0x21, 0x3e, 0x80, 0xa2, 0xa1, 0x6e, 0x20, 0x0d, 0xa5, 0x68, 0x48, 0xbc, 0x44, 0x4e,
	0xe2, 0xd6, 0x56, 0x33, 0x3b, 0x4a, 0x4c, 0xd5, 0xf2, 0x15, 0x7c, 0x0e, 0x9f, 0xc0, 0xe3, 0x5e,
	0x90, 0x78, 0x44, 0xed, 0x57, 0xf0, 0x86, 0x72, 0xe3, 0xb4, 0x2b, 0x62, 0xa0, 0xbd, 0x54, 0x3d,
	0xbe, 0xe7, 0x9c, 0xdc, 0x7b, 0xae, 0x8d, 0x1f, 0x07, 0x34, 0x98, 0xc7, 0x4a, 0xba, 0x9f, 0x95,
	0x64, 0xa1, 0x92, 0xa1, 0x60, 0xe9, 0x98, 0x6d, 0x22, 0x27, 0x49, 0x95, 0x56, 0xc4, 0x32, 0x54,
	0x67, 0xb3, 0x38, 0x7d, 0xb6, 0xbb, 0x57, 0x9a, 0x84, 0x9c, 0x85, 0x93, 0x44, 0x09, 0xa9, 0x85,
	0x1c, 0xbb, 0x41, 0x9c, 0xf9, 0x13, 0x36, 0x2f, 0xe4, 0xbb, 0xfb, 0x7f, 0x27, 0xad, 0x91, 0xe1,
	0xdd, 0xd7, 0x4c, 0x46, 0x2c, 0x3d, 0x13, 0x52, 0xbb, 0x7a, 0x9e, 0xb0, 0xac, 0xf8, 0x35, 0xd5,
	0x07, 0x17, 0xaa, 0x61, 0x3a, 0x4f, 0xb4, 0x72, 0x93, 0x54, 0xa9, 0x51, 0x51, 0xee, 0x7e, 0x47,
	0xb8, 0x75, 0x24, 0x23, 0x36, 0x63, 0xd1, 0x80, 0xd1, 0x88, 0xa5, 0xe4, 0x2e, 0xbe, 0x1e, 0x72,
	0x2a, 0xa4, 0x2f, 0x22, 0x0b, 0x75, 0x50, 0x6f, 0xc7, 0xdb, 0x06, 0x7c, 0x14, 0x11, 0x82, 0x6b,
	0x9c, 0x66, 0xdc, 0xba, 0xd6, 0x41, 0xbd, 0xa6, 0x07, 0xff, 0xc9, 0x1d, 0x5c, 0xe7, 0x4c, 0x8c,
	0xb9, 0xb6, 0xaa, 0x1d, 0xd4, 0xab, 0x79, 0x06, 0x91, 0xa7, 0xf8, 0xb6, 0xe9, 0xdf, 0x0f, 0x62,
	0x15, 0x4e, 0x7c, 0xc3, 0xaa, 0x01, 0x8b, 0x98, 0x5a, 0x3f, 0x2f, 0x0d, 0x0a, 0xc5, 0x1e, 0x6e,
	0x95, 0x0a, 0x96, 0xa8, 0x90, 0x5b, 0x5b, 0x40, 0x6d, 0x9a, 0xc3, 0xc3, 0xfc, 0x8c, 0xec, 0xe3,
	0x9b, 0x25, 0x49, 0xcf, 0x7c, 0xe8, 0xa6, 0x0e, 0xdd, 0x94, 0xda, 0xf7, 0xb3, 0x01, 0xcd, 0x78,
	0xf7, 0x18, 0x6f, 0xbd, 0x56, 0xe9, 0x24, 0x23, 0x2f, 0xf1, 0x36, 0x87, 0xc1, 0x32, 0xab, 0xda,
	0xa9, 0xf6, 0x1a, 0x07, 0x8f, 0x9c, 0xcb, 0xd6, 0xe2, 0x6c, 0x04, 0xe1, 0x95, 0xba, 0xee, 0x57,
	0x84, 0x77, 0x5e, 0x41, 0x04, 0x72, 0xa4, 0xfe, 0x95, 0xcf, 0x5b, 0xdc, 0x8a, 0xa9, 0x66, 0x99,
	0xf6, 0x0b, 0x29, 0x04, 0x75, 0x85, 0x2f, 0x36, 0x0b, 0xb5, 0x59, 0x44, 0x1f, 0x1b, 0xec, 0x8f,
	0xf2, 0x49, 0x20, 0xdf, 0xc6, 0xc1, 0xc3, 0xcb, 0xcd, 0x60, 0x60, 0xaf, 0x51, 0x88, 0x00, 0x74,
	0x7f, 0x21, 0xdc, 0x7e, 0x97, 0xaf, 0x1b, 0xd2, 0x1b, 0x32, 0x1a, 0xb3, 0x88, 0x78, 0xb8, 0x35,
	0xa5, 0xb1, 0x88, 0xa8, 0x56, 0xa9, 0x9f, 0x31, 0x6d, 0x21, 0x08, 0xe6, 0xc9, 0xca, 0x79, 0xe3,
	0xc2, 0xe5, 0xce, 0xa7, 0x25, 0xfd, 0x83, 0xd0, 0xbc, 0x1f, 0x67, 0x6f, 0xd8, 0xdc, 0x6b, 0xae,
	0x3c, 0x86, 0x4c, 0x93, 0x43, 0xdc, 0x86, 0x6b, 0x55, 0xac, 0xce, 0x17, 0x72, 0xa4, 0xcc, 0xf4,
	0xf7, 0x9c, 0xf5, 0x0d, 0x74, 0x8a, 0x1b, 0xe8, 0x40, 0x4b, 0x27, 0x49, 0xe6, 0xdd, 0x48, 0x56,
	0xcd, 0x41, 0xb8, 0xc7, 0xf8, 0xd6, 0x45, 0x9b, 0x29, 0x8d, 0xa1, 0xc1, 0xea, 0xff, 0x9d, 0xda,
	0x6b, 0xa7, 0x53, 0x1a, 0x0f, 0x99, 0xee, 0x9f, 0x7c, 0x5b, 0xd8, 0xe8, 0x7c, 0x61, 0xa3, 0x9f,
	0x0b, 0x1b, 0x7d, 0x59, 0xda, 0x95, 0xf3, 0xa5, 0x5d, 0xf9, 0xb1, 0xb4, 0x2b, 0x1f, 0x5f, 0x8c,
	0x85, 0xe6, 0x9f, 0x02, 0x27, 0x54, 0x67, 0xae, 0x99, 0x19, 0x56, 0x58, 0x02, 0x77, 0xf6, 0xc7,
	0xeb, 0x86, 0x07, 0x15, 0xd4, 0xe1, 0xc9, 0x3c, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x6c,
	0x87, 0xd2, 0x03, 0x04, 0x00, 0x00,
}

func (m *IndexedHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexedHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexedHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BabylonTxHash) > 0 {
		i -= len(m.BabylonTxHash)
		copy(dAtA[i:], m.BabylonTxHash)
		i = encodeVarintZoneconcierge(dAtA, i, uint64(len(m.BabylonTxHash)))
		i--
		dAtA[i] = 0x32
	}
	if m.BabylonEpoch != 0 {
		i = encodeVarintZoneconcierge(dAtA, i, uint64(m.BabylonEpoch))
		i--
		dAtA[i] = 0x28
	}
	if m.BabylonBlockHeight != 0 {
		i = encodeVarintZoneconcierge(dAtA, i, uint64(m.BabylonBlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.Height != 0 {
		i = encodeVarintZoneconcierge(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintZoneconcierge(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintZoneconcierge(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Forks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Forks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Forks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Headers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintZoneconcierge(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	return len(dAtA) - i, nil
}

func (m *ChainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LatestForks != nil {
		{
			size, err := m.LatestForks.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintZoneconcierge(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.LatestHeader != nil {
		{
			size, err := m.LatestHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintZoneconcierge(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintZoneconcierge(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProofEpochSealed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofEpochSealed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProofEpochSealed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProofEpochValSet != nil {
		{
			size, err := m.ProofEpochValSet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintZoneconcierge(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ProofEpochInfo != nil {
		{
			size, err := m.ProofEpochInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintZoneconcierge(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ValidatorSet) > 0 {
		for iNdEx := len(m.ValidatorSet) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorSet[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintZoneconcierge(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintZoneconcierge(dAtA []byte, offset int, v uint64) int {
	offset -= sovZoneconcierge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IndexedHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovZoneconcierge(uint64(m.Height))
	}
	if m.BabylonBlockHeight != 0 {
		n += 1 + sovZoneconcierge(uint64(m.BabylonBlockHeight))
	}
	if m.BabylonEpoch != 0 {
		n += 1 + sovZoneconcierge(uint64(m.BabylonEpoch))
	}
	l = len(m.BabylonTxHash)
	if l > 0 {
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	return n
}

func (m *Forks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovZoneconcierge(uint64(l))
		}
	}
	return n
}

func (m *ChainInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	if m.LatestHeader != nil {
		l = m.LatestHeader.Size()
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	if m.LatestForks != nil {
		l = m.LatestForks.Size()
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	return n
}

func (m *ProofEpochSealed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ValidatorSet) > 0 {
		for _, e := range m.ValidatorSet {
			l = e.Size()
			n += 1 + l + sovZoneconcierge(uint64(l))
		}
	}
	if m.ProofEpochInfo != nil {
		l = m.ProofEpochInfo.Size()
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	if m.ProofEpochValSet != nil {
		l = m.ProofEpochValSet.Size()
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	return n
}

func sovZoneconcierge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozZoneconcierge(x uint64) (n int) {
	return sovZoneconcierge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexedHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZoneconcierge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IndexedHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexedHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BabylonBlockHeight", wireType)
			}
			m.BabylonBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BabylonBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BabylonEpoch", wireType)
			}
			m.BabylonEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BabylonEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BabylonTxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BabylonTxHash = append(m.BabylonTxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BabylonTxHash == nil {
				m.BabylonTxHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZoneconcierge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Forks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZoneconcierge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Forks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Forks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &IndexedHeader{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZoneconcierge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChainInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZoneconcierge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LatestHeader == nil {
				m.LatestHeader = &IndexedHeader{}
			}
			if err := m.LatestHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestForks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LatestForks == nil {
				m.LatestForks = &Forks{}
			}
			if err := m.LatestForks.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZoneconcierge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProofEpochSealed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZoneconcierge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProofEpochSealed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofEpochSealed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorSet = append(m.ValidatorSet, &types.ValidatorWithBlsKey{})
			if err := m.ValidatorSet[len(m.ValidatorSet)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofEpochInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ProofEpochInfo == nil {
				m.ProofEpochInfo = &crypto.ProofOps{}
			}
			if err := m.ProofEpochInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofEpochValSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ProofEpochValSet == nil {
				m.ProofEpochValSet = &crypto.ProofOps{}
			}
			if err := m.ProofEpochValSet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZoneconcierge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipZoneconcierge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowZoneconcierge
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthZoneconcierge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupZoneconcierge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthZoneconcierge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthZoneconcierge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowZoneconcierge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupZoneconcierge = fmt.Errorf("proto: unexpected end of group")
)
