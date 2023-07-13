// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/finality/v1/finality.proto

package types

import (
	fmt "fmt"
	github_com_babylonchain_babylon_types "github.com/babylonchain/babylon/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// IndexedBlock is the block with some indexed info
type IndexedBlock struct {
	// height is the height of the block
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// last_commit_hash is the last_commit_hash of the block
	LastCommitHash []byte `protobuf:"bytes,2,opt,name=last_commit_hash,json=lastCommitHash,proto3" json:"last_commit_hash,omitempty"`
	// finalized indicates whether the IndexedBlock is finalised by 2/3
	// BTC validators or not
	Finalized bool `protobuf:"varint,3,opt,name=finalized,proto3" json:"finalized,omitempty"`
}

func (m *IndexedBlock) Reset()         { *m = IndexedBlock{} }
func (m *IndexedBlock) String() string { return proto.CompactTextString(m) }
func (*IndexedBlock) ProtoMessage()    {}
func (*IndexedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca5b87e52e3e6d02, []int{0}
}
func (m *IndexedBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexedBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexedBlock.Merge(m, src)
}
func (m *IndexedBlock) XXX_Size() int {
	return m.Size()
}
func (m *IndexedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_IndexedBlock proto.InternalMessageInfo

func (m *IndexedBlock) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *IndexedBlock) GetLastCommitHash() []byte {
	if m != nil {
		return m.LastCommitHash
	}
	return nil
}

func (m *IndexedBlock) GetFinalized() bool {
	if m != nil {
		return m.Finalized
	}
	return false
}

// Evidence is the evidence that a BTC validator has signed a finality
// signature with correct public randomness on a fork header
// It includes all fields of MsgAddFinalitySig, such that anyone seeing
// the evidence and a signature on the canonical fork can extract the
// BTC validator's BTC secret key.
type Evidence struct {
	// val_btc_pk is the BTC Pk of the validator that casts this vote
	ValBtcPk *github_com_babylonchain_babylon_types.BIP340PubKey `protobuf:"bytes,1,opt,name=val_btc_pk,json=valBtcPk,proto3,customtype=github.com/babylonchain/babylon/types.BIP340PubKey" json:"val_btc_pk,omitempty"`
	// block_height is the height of the voted block
	BlockHeight uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// block_last_commit_hash is the last_commit_hash of the voted block
	BlockLastCommitHash []byte `protobuf:"bytes,3,opt,name=block_last_commit_hash,json=blockLastCommitHash,proto3" json:"block_last_commit_hash,omitempty"`
	// finality_sig is the finality signature to this block
	// where finality signature is an EOTS signature, i.e.,
	// the `s` in a Schnorr signature `(r, s)`
	// `r` is the public randomness that is already committed by the validator
	FinalitySig *github_com_babylonchain_babylon_types.SchnorrEOTSSig `protobuf:"bytes,4,opt,name=finality_sig,json=finalitySig,proto3,customtype=github.com/babylonchain/babylon/types.SchnorrEOTSSig" json:"finality_sig,omitempty"`
}

func (m *Evidence) Reset()         { *m = Evidence{} }
func (m *Evidence) String() string { return proto.CompactTextString(m) }
func (*Evidence) ProtoMessage()    {}
func (*Evidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca5b87e52e3e6d02, []int{1}
}
func (m *Evidence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Evidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Evidence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Evidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evidence.Merge(m, src)
}
func (m *Evidence) XXX_Size() int {
	return m.Size()
}
func (m *Evidence) XXX_DiscardUnknown() {
	xxx_messageInfo_Evidence.DiscardUnknown(m)
}

var xxx_messageInfo_Evidence proto.InternalMessageInfo

func (m *Evidence) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Evidence) GetBlockLastCommitHash() []byte {
	if m != nil {
		return m.BlockLastCommitHash
	}
	return nil
}

func init() {
	proto.RegisterType((*IndexedBlock)(nil), "babylon.finality.v1.IndexedBlock")
	proto.RegisterType((*Evidence)(nil), "babylon.finality.v1.Evidence")
}

func init() {
	proto.RegisterFile("babylon/finality/v1/finality.proto", fileDescriptor_ca5b87e52e3e6d02)
}

var fileDescriptor_ca5b87e52e3e6d02 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0xe2, 0x40,
	0x14, 0xc7, 0x8d, 0x8a, 0xb8, 0x63, 0x58, 0x96, 0xb8, 0x48, 0x58, 0x96, 0xe8, 0x7a, 0xca, 0x29,
	0xd1, 0x55, 0x96, 0x3d, 0xa7, 0x08, 0xda, 0x16, 0x2a, 0x89, 0xa7, 0xf6, 0x10, 0x26, 0x93, 0x69,
	0x66, 0x30, 0x66, 0xc4, 0x8c, 0xc1, 0xf4, 0x53, 0xf4, 0xd2, 0xef, 0xd4, 0xa3, 0xc7, 0xe2, 0x41,
	0x8a, 0x7e, 0x91, 0xe2, 0x24, 0x2a, 0xa5, 0x87, 0xf6, 0xf6, 0xde, 0xff, 0xfd, 0x79, 0xff, 0xf7,
	0x4b, 0x06, 0xb4, 0x3d, 0xe8, 0xa5, 0x21, 0x8b, 0xcc, 0x7b, 0x1a, 0xc1, 0x90, 0xf2, 0xd4, 0x4c,
	0xba, 0xa7, 0xda, 0x98, 0x2f, 0x18, 0x67, 0x4a, 0x3d, 0xf7, 0x18, 0x27, 0x3d, 0xe9, 0xfe, 0xfa,
	0x19, 0xb0, 0x80, 0x89, 0xb9, 0x79, 0xa8, 0x32, 0x6b, 0x3b, 0x02, 0xf2, 0x28, 0xf2, 0xf1, 0x0a,
	0xfb, 0x56, 0xc8, 0xd0, 0x54, 0x69, 0x80, 0x0a, 0xc1, 0x34, 0x20, 0x5c, 0x95, 0x5a, 0x92, 0x5e,
	0xb6, 0xf3, 0x4e, 0xd1, 0xc1, 0x8f, 0x10, 0xc6, 0xdc, 0x45, 0x6c, 0x36, 0xa3, 0xdc, 0x25, 0x30,
	0x26, 0x6a, 0xb1, 0x25, 0xe9, 0xb2, 0xfd, 0xfd, 0xa0, 0x5f, 0x08, 0x79, 0x08, 0x63, 0xa2, 0xfc,
	0x06, 0xdf, 0xb2, 0xd8, 0x07, 0xec, 0xab, 0xa5, 0x96, 0xa4, 0x57, 0xed, 0xb3, 0xd0, 0x7e, 0x2a,
	0x82, 0xea, 0x20, 0xa1, 0x3e, 0x8e, 0x10, 0x56, 0x26, 0x00, 0x24, 0x30, 0x74, 0x3d, 0x8e, 0xdc,
	0xf9, 0x54, 0x04, 0xca, 0xd6, 0xbf, 0xcd, 0xb6, 0xf9, 0x37, 0xa0, 0x9c, 0x2c, 0x3d, 0x03, 0xb1,
	0x99, 0x99, 0xa3, 0x20, 0x02, 0x69, 0x74, 0x6c, 0x4c, 0x9e, 0xce, 0x71, 0x6c, 0x58, 0xa3, 0x71,
	0xaf, 0xdf, 0x19, 0x2f, 0xbd, 0x2b, 0x9c, 0xda, 0xd5, 0x04, 0x86, 0x16, 0x47, 0xe3, 0xa9, 0xf2,
	0x07, 0xc8, 0xde, 0x81, 0xc5, 0xcd, 0x41, 0x8a, 0x02, 0xa4, 0x26, 0xb4, 0x61, 0x46, 0xd3, 0x03,
	0x8d, 0xcc, 0xf2, 0x81, 0xa9, 0x24, 0x98, 0xea, 0x62, 0x7a, 0xfd, 0x1e, 0xec, 0x0e, 0xc8, 0xc7,
	0xef, 0xe9, 0xc6, 0x34, 0x50, 0xcb, 0xe2, 0xde, 0xff, 0x9b, 0x6d, 0xb3, 0xff, 0xb5, 0x7b, 0x1d,
	0x44, 0x22, 0xb6, 0x58, 0x0c, 0x6e, 0x26, 0x8e, 0x43, 0x03, 0xbb, 0x76, 0xdc, 0xe6, 0xd0, 0xc0,
	0xba, 0x7c, 0xde, 0x69, 0xd2, 0x7a, 0xa7, 0x49, 0xaf, 0x3b, 0x4d, 0x7a, 0xdc, 0x6b, 0x85, 0xf5,
	0x5e, 0x2b, 0xbc, 0xec, 0xb5, 0xc2, 0x6d, 0xe7, 0xb3, 0xe5, 0xab, 0xf3, 0x53, 0x10, 0x39, 0x5e,
	0x45, 0xfc, 0xda, 0xde, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0x87, 0x93, 0x2b, 0x2b, 0x02,
	0x00, 0x00,
}

func (m *IndexedBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexedBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexedBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Finalized {
		i--
		if m.Finalized {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.LastCommitHash) > 0 {
		i -= len(m.LastCommitHash)
		copy(dAtA[i:], m.LastCommitHash)
		i = encodeVarintFinality(dAtA, i, uint64(len(m.LastCommitHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintFinality(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Evidence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Evidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Evidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinalitySig != nil {
		{
			size := m.FinalitySig.Size()
			i -= size
			if _, err := m.FinalitySig.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintFinality(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.BlockLastCommitHash) > 0 {
		i -= len(m.BlockLastCommitHash)
		copy(dAtA[i:], m.BlockLastCommitHash)
		i = encodeVarintFinality(dAtA, i, uint64(len(m.BlockLastCommitHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintFinality(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.ValBtcPk != nil {
		{
			size := m.ValBtcPk.Size()
			i -= size
			if _, err := m.ValBtcPk.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintFinality(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFinality(dAtA []byte, offset int, v uint64) int {
	offset -= sovFinality(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IndexedBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovFinality(uint64(m.Height))
	}
	l = len(m.LastCommitHash)
	if l > 0 {
		n += 1 + l + sovFinality(uint64(l))
	}
	if m.Finalized {
		n += 2
	}
	return n
}

func (m *Evidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValBtcPk != nil {
		l = m.ValBtcPk.Size()
		n += 1 + l + sovFinality(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovFinality(uint64(m.BlockHeight))
	}
	l = len(m.BlockLastCommitHash)
	if l > 0 {
		n += 1 + l + sovFinality(uint64(l))
	}
	if m.FinalitySig != nil {
		l = m.FinalitySig.Size()
		n += 1 + l + sovFinality(uint64(l))
	}
	return n
}

func sovFinality(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFinality(x uint64) (n int) {
	return sovFinality(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexedBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFinality
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
			return fmt.Errorf("proto: IndexedBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexedBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastCommitHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastCommitHash = append(m.LastCommitHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastCommitHash == nil {
				m.LastCommitHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Finalized", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Finalized = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFinality(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFinality
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
func (m *Evidence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFinality
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
			return fmt.Errorf("proto: Evidence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Evidence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValBtcPk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.BIP340PubKey
			m.ValBtcPk = &v
			if err := m.ValBtcPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockLastCommitHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockLastCommitHash = append(m.BlockLastCommitHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockLastCommitHash == nil {
				m.BlockLastCommitHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalitySig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.SchnorrEOTSSig
			m.FinalitySig = &v
			if err := m.FinalitySig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFinality(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFinality
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
func skipFinality(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFinality
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
					return 0, ErrIntOverflowFinality
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
					return 0, ErrIntOverflowFinality
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
				return 0, ErrInvalidLengthFinality
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFinality
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFinality
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFinality        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFinality          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFinality = fmt.Errorf("proto: unexpected end of group")
)