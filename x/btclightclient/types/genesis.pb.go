// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/btclightclient/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// GenesisState defines the btclightclient module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_723ed9409b965050, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// this structure can go on a separate file
type BitcoinHeader struct {
	Version    int32                  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	ParentHash []byte                 `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	MerkleRoot []byte                 `protobuf:"bytes,3,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	Time       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Nbits      uint32                 `protobuf:"varint,5,opt,name=nbits,proto3" json:"nbits,omitempty"`
	Nonce      uint32                 `protobuf:"varint,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *BitcoinHeader) Reset()         { *m = BitcoinHeader{} }
func (m *BitcoinHeader) String() string { return proto.CompactTextString(m) }
func (*BitcoinHeader) ProtoMessage()    {}
func (*BitcoinHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_723ed9409b965050, []int{1}
}
func (m *BitcoinHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BitcoinHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BitcoinHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BitcoinHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BitcoinHeader.Merge(m, src)
}
func (m *BitcoinHeader) XXX_Size() int {
	return m.Size()
}
func (m *BitcoinHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BitcoinHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BitcoinHeader proto.InternalMessageInfo

func (m *BitcoinHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BitcoinHeader) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *BitcoinHeader) GetMerkleRoot() []byte {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

func (m *BitcoinHeader) GetTime() *timestamppb.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *BitcoinHeader) GetNbits() uint32 {
	if m != nil {
		return m.Nbits
	}
	return 0
}

func (m *BitcoinHeader) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "babylon.btclightclient.v1.GenesisState")
	proto.RegisterType((*BitcoinHeader)(nil), "babylon.btclightclient.v1.BitcoinHeader")
}

func init() {
	proto.RegisterFile("babylon/btclightclient/genesis.proto", fileDescriptor_723ed9409b965050)
}

var fileDescriptor_723ed9409b965050 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xb1, 0x4e, 0xe3, 0x40,
	0x18, 0x84, 0xbd, 0x77, 0x49, 0x4e, 0xda, 0x24, 0x8d, 0x95, 0xc2, 0xe7, 0xc2, 0xc9, 0xe5, 0x28,
	0x52, 0xad, 0x05, 0x48, 0xb4, 0x48, 0x69, 0x48, 0x47, 0x64, 0xa8, 0x68, 0xa2, 0xb5, 0x59, 0xec,
	0x15, 0xf6, 0xfe, 0x96, 0xf7, 0x4f, 0x44, 0xde, 0x82, 0xb7, 0x22, 0x65, 0x4a, 0x2a, 0x84, 0x92,
	0x17, 0x41, 0xde, 0xb5, 0x0b, 0x90, 0xd2, 0x58, 0x9e, 0xd1, 0xf7, 0x8f, 0x66, 0x6c, 0x7a, 0x16,
	0xf3, 0x78, 0x9b, 0x83, 0x0a, 0x63, 0x4c, 0x72, 0x99, 0x66, 0xf5, 0x53, 0x28, 0x0c, 0x53, 0xa1,
	0x84, 0x96, 0x9a, 0x95, 0x15, 0x20, 0xb8, 0x7f, 0x1b, 0x8a, 0x7d, 0xa7, 0xd8, 0xe6, 0xdc, 0x1f,
	0xa5, 0x90, 0x82, 0xa1, 0xc2, 0xfa, 0xcd, 0x1e, 0xf8, 0xff, 0x4f, 0xc4, 0x96, 0xbc, 0xe2, 0x45,
	0x93, 0xea, 0x8f, 0x53, 0x80, 0x34, 0x17, 0xa1, 0x51, 0xf1, 0xfa, 0x29, 0x44, 0x59, 0x08, 0x8d,
	0xbc, 0x28, 0x2d, 0x30, 0xbd, 0xa5, 0x83, 0x1b, 0xdb, 0xe3, 0x0e, 0x39, 0x0a, 0xf7, 0x9a, 0xf6,
	0x6c, 0x80, 0x47, 0x26, 0x64, 0xd6, 0xbf, 0xf8, 0xc7, 0x4e, 0xf6, 0x62, 0x4b, 0x03, 0xce, 0x3b,
	0xbb, 0x8f, 0xb1, 0x13, 0x35, 0x67, 0xd3, 0x37, 0x42, 0x87, 0x73, 0x89, 0x09, 0x48, 0xb5, 0x10,
	0xfc, 0x51, 0x54, 0xae, 0x47, 0xff, 0x6c, 0x44, 0xa5, 0x25, 0x28, 0x93, 0xd9, 0x8d, 0x5a, 0xe9,
	0x8e, 0x69, 0xbf, 0xe4, 0x95, 0x50, 0xb8, 0xca, 0xb8, 0xce, 0xbc, 0x5f, 0x13, 0x32, 0x1b, 0x44,
	0xd4, 0x5a, 0x0b, 0xae, 0xb3, 0x1a, 0x28, 0x44, 0xf5, 0x9c, 0x8b, 0x55, 0x05, 0x80, 0xde, 0x6f,
	0x0b, 0x58, 0x2b, 0x02, 0x40, 0x97, 0xd1, 0x4e, 0xbd, 0xc8, 0xeb, 0x98, 0xb2, 0x3e, 0xb3, 0x73,
	0x59, 0x3b, 0x97, 0xdd, 0xb7, 0x73, 0x23, 0xc3, 0xb9, 0x23, 0xda, 0x55, 0xb1, 0x44, 0xed, 0x75,
	0x27, 0x64, 0x36, 0x8c, 0xac, 0x30, 0x2e, 0xa8, 0x44, 0x78, 0xbd, 0xc6, 0xad, 0xc5, 0x7c, 0xb9,
	0x3b, 0x04, 0x64, 0x7f, 0x08, 0xc8, 0xe7, 0x21, 0x20, 0xaf, 0xc7, 0xc0, 0xd9, 0x1f, 0x03, 0xe7,
	0xfd, 0x18, 0x38, 0x0f, 0x57, 0xa9, 0xc4, 0x6c, 0x1d, 0xb3, 0x04, 0x8a, 0xb0, 0xf9, 0x3c, 0x49,
	0xc6, 0xa5, 0x6a, 0x45, 0xf8, 0xf2, 0xf3, 0xa7, 0xe0, 0xb6, 0x14, 0x3a, 0xee, 0x99, 0x5e, 0x97,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x0c, 0x89, 0x99, 0x12, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BitcoinHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BitcoinHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BitcoinHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nonce != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x30
	}
	if m.Nbits != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Nbits))
		i--
		dAtA[i] = 0x28
	}
	if m.Time != nil {
		{
			size, err := m.Time.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.MerkleRoot) > 0 {
		i -= len(m.MerkleRoot)
		copy(dAtA[i:], m.MerkleRoot)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.MerkleRoot)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ParentHash) > 0 {
		i -= len(m.ParentHash)
		copy(dAtA[i:], m.ParentHash)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ParentHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Version != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *BitcoinHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovGenesis(uint64(m.Version))
	}
	l = len(m.ParentHash)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.MerkleRoot)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Time != nil {
		l = m.Time.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Nbits != 0 {
		n += 1 + sovGenesis(uint64(m.Nbits))
	}
	if m.Nonce != 0 {
		n += 1 + sovGenesis(uint64(m.Nonce))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *BitcoinHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: BitcoinHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BitcoinHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentHash = append(m.ParentHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ParentHash == nil {
				m.ParentHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerkleRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MerkleRoot = append(m.MerkleRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.MerkleRoot == nil {
				m.MerkleRoot = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Time == nil {
				m.Time = &timestamppb.Timestamp{}
			}
			if err := m.Time.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nbits", wireType)
			}
			m.Nbits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nbits |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
