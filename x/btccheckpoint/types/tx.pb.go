// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/btccheckpoint/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Consider we have a Merkle tree with following structure:
//            ROOT
//           /    \
//      H1234      H5555
//     /     \       \
//   H12     H34      H55
//  /  \    /  \     /
// H1  H2  H3  H4  H5
// L1  L2  L3  L4  L5
// To prove L3 was part of ROOT we need:
// - btc_transaction_index = 2 which in binary is 010
// (where 0 means going left, 1 means going right in the tree)
// - merkle_nodes we'd have H4 || H12 || H5555
// By looking at 010 we would know that H4 is a right sibling,
// H12 is left, H5555 is right again.
type BTCSpvProof struct {
	// Valid bitcoin transaction containing OP_RETURN opcode.
	BtcTransaction []byte `protobuf:"bytes,1,opt,name=btc_transaction,json=btcTransaction,proto3" json:"btc_transaction,omitempty"`
	// Index of transaction within the block. Index is needed to determine if
	// currently hashed node is left or right.
	BtcTransactionIndex uint32 `protobuf:"varint,2,opt,name=btc_transaction_index,json=btcTransactionIndex,proto3" json:"btc_transaction_index,omitempty"`
	// List of concatenated intermediate merkle tree nodes, without root node and leaf node
	// against which we calculate the proof.
	// Each node has 32 byte length.
	// Example proof can look like: 32_bytes_of_node1 || 32_bytes_of_node2 ||  32_bytes_of_node3
	// so the length of the proof will always be divisible by 32.
	MerkleNodes []byte `protobuf:"bytes,3,opt,name=merkle_nodes,json=merkleNodes,proto3" json:"merkle_nodes,omitempty"`
	// Valid btc header which confirms btc_transaction.
	// Should have exactly 80 bytes
	ConfirmingBtcHeader []byte `protobuf:"bytes,4,opt,name=confirming_btc_header,json=confirmingBtcHeader,proto3" json:"confirming_btc_header,omitempty"`
}

func (m *BTCSpvProof) Reset()         { *m = BTCSpvProof{} }
func (m *BTCSpvProof) String() string { return proto.CompactTextString(m) }
func (*BTCSpvProof) ProtoMessage()    {}
func (*BTCSpvProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeec89810b39ea83, []int{0}
}
func (m *BTCSpvProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BTCSpvProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BTCSpvProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BTCSpvProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BTCSpvProof.Merge(m, src)
}
func (m *BTCSpvProof) XXX_Size() int {
	return m.Size()
}
func (m *BTCSpvProof) XXX_DiscardUnknown() {
	xxx_messageInfo_BTCSpvProof.DiscardUnknown(m)
}

var xxx_messageInfo_BTCSpvProof proto.InternalMessageInfo

func (m *BTCSpvProof) GetBtcTransaction() []byte {
	if m != nil {
		return m.BtcTransaction
	}
	return nil
}

func (m *BTCSpvProof) GetBtcTransactionIndex() uint32 {
	if m != nil {
		return m.BtcTransactionIndex
	}
	return 0
}

func (m *BTCSpvProof) GetMerkleNodes() []byte {
	if m != nil {
		return m.MerkleNodes
	}
	return nil
}

func (m *BTCSpvProof) GetConfirmingBtcHeader() []byte {
	if m != nil {
		return m.ConfirmingBtcHeader
	}
	return nil
}

type MsgInsertBTCSpvProof struct {
	Proofs []*BTCSpvProof `protobuf:"bytes,1,rep,name=proofs,proto3" json:"proofs,omitempty"`
}

func (m *MsgInsertBTCSpvProof) Reset()         { *m = MsgInsertBTCSpvProof{} }
func (m *MsgInsertBTCSpvProof) String() string { return proto.CompactTextString(m) }
func (*MsgInsertBTCSpvProof) ProtoMessage()    {}
func (*MsgInsertBTCSpvProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeec89810b39ea83, []int{1}
}
func (m *MsgInsertBTCSpvProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInsertBTCSpvProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInsertBTCSpvProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInsertBTCSpvProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInsertBTCSpvProof.Merge(m, src)
}
func (m *MsgInsertBTCSpvProof) XXX_Size() int {
	return m.Size()
}
func (m *MsgInsertBTCSpvProof) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInsertBTCSpvProof.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInsertBTCSpvProof proto.InternalMessageInfo

func (m *MsgInsertBTCSpvProof) GetProofs() []*BTCSpvProof {
	if m != nil {
		return m.Proofs
	}
	return nil
}

type MsgInsertBTCSpvProofResponse struct {
}

func (m *MsgInsertBTCSpvProofResponse) Reset()         { *m = MsgInsertBTCSpvProofResponse{} }
func (m *MsgInsertBTCSpvProofResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInsertBTCSpvProofResponse) ProtoMessage()    {}
func (*MsgInsertBTCSpvProofResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeec89810b39ea83, []int{2}
}
func (m *MsgInsertBTCSpvProofResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInsertBTCSpvProofResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInsertBTCSpvProofResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInsertBTCSpvProofResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInsertBTCSpvProofResponse.Merge(m, src)
}
func (m *MsgInsertBTCSpvProofResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInsertBTCSpvProofResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInsertBTCSpvProofResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInsertBTCSpvProofResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BTCSpvProof)(nil), "babylon.btccheckpoint.BTCSpvProof")
	proto.RegisterType((*MsgInsertBTCSpvProof)(nil), "babylon.btccheckpoint.MsgInsertBTCSpvProof")
	proto.RegisterType((*MsgInsertBTCSpvProofResponse)(nil), "babylon.btccheckpoint.MsgInsertBTCSpvProofResponse")
}

func init() { proto.RegisterFile("babylon/btccheckpoint/tx.proto", fileDescriptor_aeec89810b39ea83) }

var fileDescriptor_aeec89810b39ea83 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4f, 0xfa, 0x30,
	0x1c, 0xc6, 0xe9, 0x8f, 0x5f, 0x38, 0x14, 0xd4, 0x38, 0x24, 0x59, 0x8c, 0x69, 0x70, 0x17, 0x49,
	0x4c, 0xb6, 0x04, 0xe2, 0xc5, 0x23, 0x5e, 0xe4, 0x80, 0x9a, 0xc9, 0xc9, 0xcb, 0xb2, 0x96, 0xb2,
	0x35, 0x40, 0xdb, 0xb4, 0xc5, 0x40, 0xe2, 0x8b, 0xf0, 0x2d, 0x79, 0xf3, 0xc8, 0xd1, 0xa3, 0x61,
	0x6f, 0xc4, 0x6c, 0x8e, 0xf0, 0xc7, 0x79, 0xf0, 0xd6, 0x3e, 0xcf, 0xf3, 0xfd, 0x3c, 0x69, 0xf3,
	0x85, 0x08, 0x87, 0x78, 0x31, 0x11, 0xdc, 0xc3, 0x86, 0x90, 0x98, 0x92, 0xb1, 0x14, 0x8c, 0x1b,
	0xcf, 0xcc, 0x5d, 0xa9, 0x84, 0x11, 0x56, 0x23, 0xf7, 0xdd, 0x1d, 0xdf, 0x79, 0x03, 0xb0, 0xda,
	0x1d, 0xdc, 0x3c, 0xca, 0xe7, 0x07, 0x25, 0xc4, 0xc8, 0xba, 0x80, 0x47, 0xd8, 0x90, 0xc0, 0xa8,
	0x90, 0xeb, 0x90, 0x18, 0x26, 0xb8, 0x0d, 0x9a, 0xa0, 0x55, 0xf3, 0x0f, 0xb1, 0x21, 0x83, 0x8d,
	0x6a, 0xb5, 0x61, 0x63, 0x2f, 0x18, 0x30, 0x3e, 0xa4, 0x73, 0xfb, 0x5f, 0x13, 0xb4, 0x0e, 0xfc,
	0xfa, 0x6e, 0xbc, 0x97, 0x5a, 0xd6, 0x39, 0xac, 0x4d, 0xa9, 0x1a, 0x4f, 0x68, 0xc0, 0xc5, 0x90,
	0x6a, 0xbb, 0x9c, 0x91, 0xab, 0xdf, 0xda, 0x5d, 0x2a, 0xa5, 0x58, 0x22, 0xf8, 0x88, 0xa9, 0x29,
	0xe3, 0x51, 0x90, 0x36, 0xc4, 0x34, 0x1c, 0x52, 0x65, 0xff, 0xcf, 0xb2, 0xf5, 0x8d, 0xd9, 0x35,
	0xe4, 0x36, 0xb3, 0x1c, 0x1f, 0x9e, 0xf4, 0x75, 0xd4, 0xe3, 0x9a, 0x2a, 0xb3, 0xfd, 0x96, 0x6b,
	0x58, 0x91, 0xe9, 0x41, 0xdb, 0xa0, 0x59, 0x6e, 0x55, 0xdb, 0x8e, 0x5b, 0xf8, 0x07, 0xee, 0xd6,
	0x8c, 0x9f, 0x4f, 0x38, 0x08, 0x9e, 0x15, 0x31, 0x7d, 0xaa, 0xa5, 0xe0, 0x9a, 0xb6, 0x5f, 0x60,
	0xb9, 0xaf, 0x23, 0x6b, 0x06, 0x8f, 0x7f, 0xf6, 0x5e, 0xfe, 0xd2, 0x53, 0x04, 0x3c, 0xed, 0xfc,
	0x21, 0xbc, 0x6e, 0xef, 0xde, 0xbf, 0xaf, 0x10, 0x58, 0xae, 0x10, 0xf8, 0x5c, 0x21, 0xf0, 0x9a,
	0xa0, 0xd2, 0x32, 0x41, 0xa5, 0x8f, 0x04, 0x95, 0x9e, 0xae, 0x22, 0x66, 0xe2, 0x19, 0x76, 0x89,
	0x98, 0x7a, 0x39, 0x98, 0xc4, 0x21, 0xe3, 0xeb, 0x8b, 0x37, 0xdf, 0x5f, 0x90, 0x85, 0xa4, 0x1a,
	0x57, 0xb2, 0x25, 0xe9, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0xca, 0x99, 0xa8, 0x2a, 0x46, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	InsertBTCSpvProof(ctx context.Context, in *MsgInsertBTCSpvProof, opts ...grpc.CallOption) (*MsgInsertBTCSpvProofResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) InsertBTCSpvProof(ctx context.Context, in *MsgInsertBTCSpvProof, opts ...grpc.CallOption) (*MsgInsertBTCSpvProofResponse, error) {
	out := new(MsgInsertBTCSpvProofResponse)
	err := c.cc.Invoke(ctx, "/babylon.btccheckpoint.Msg/InsertBTCSpvProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	InsertBTCSpvProof(context.Context, *MsgInsertBTCSpvProof) (*MsgInsertBTCSpvProofResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) InsertBTCSpvProof(ctx context.Context, req *MsgInsertBTCSpvProof) (*MsgInsertBTCSpvProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertBTCSpvProof not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_InsertBTCSpvProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInsertBTCSpvProof)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InsertBTCSpvProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/babylon.btccheckpoint.Msg/InsertBTCSpvProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InsertBTCSpvProof(ctx, req.(*MsgInsertBTCSpvProof))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babylon.btccheckpoint.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertBTCSpvProof",
			Handler:    _Msg_InsertBTCSpvProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "babylon/btccheckpoint/tx.proto",
}

func (m *BTCSpvProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BTCSpvProof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BTCSpvProof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConfirmingBtcHeader) > 0 {
		i -= len(m.ConfirmingBtcHeader)
		copy(dAtA[i:], m.ConfirmingBtcHeader)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConfirmingBtcHeader)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MerkleNodes) > 0 {
		i -= len(m.MerkleNodes)
		copy(dAtA[i:], m.MerkleNodes)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MerkleNodes)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BtcTransactionIndex != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BtcTransactionIndex))
		i--
		dAtA[i] = 0x10
	}
	if len(m.BtcTransaction) > 0 {
		i -= len(m.BtcTransaction)
		copy(dAtA[i:], m.BtcTransaction)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BtcTransaction)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInsertBTCSpvProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInsertBTCSpvProof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInsertBTCSpvProof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proofs) > 0 {
		for iNdEx := len(m.Proofs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proofs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgInsertBTCSpvProofResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInsertBTCSpvProofResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInsertBTCSpvProofResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BTCSpvProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BtcTransaction)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.BtcTransactionIndex != 0 {
		n += 1 + sovTx(uint64(m.BtcTransactionIndex))
	}
	l = len(m.MerkleNodes)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ConfirmingBtcHeader)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgInsertBTCSpvProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Proofs) > 0 {
		for _, e := range m.Proofs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgInsertBTCSpvProofResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BTCSpvProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: BTCSpvProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BTCSpvProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcTransaction", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BtcTransaction = append(m.BtcTransaction[:0], dAtA[iNdEx:postIndex]...)
			if m.BtcTransaction == nil {
				m.BtcTransaction = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcTransactionIndex", wireType)
			}
			m.BtcTransactionIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BtcTransactionIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerkleNodes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MerkleNodes = append(m.MerkleNodes[:0], dAtA[iNdEx:postIndex]...)
			if m.MerkleNodes == nil {
				m.MerkleNodes = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmingBtcHeader", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConfirmingBtcHeader = append(m.ConfirmingBtcHeader[:0], dAtA[iNdEx:postIndex]...)
			if m.ConfirmingBtcHeader == nil {
				m.ConfirmingBtcHeader = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgInsertBTCSpvProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInsertBTCSpvProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInsertBTCSpvProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proofs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proofs = append(m.Proofs, &BTCSpvProof{})
			if err := m.Proofs[len(m.Proofs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgInsertBTCSpvProofResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInsertBTCSpvProofResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInsertBTCSpvProofResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
