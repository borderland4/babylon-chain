// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/checkpointing/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgAddBlsSig defines a message to add a bls signature from a
// validator
type MsgAddBlsSig struct {
	BlsSig *BlsSig `protobuf:"bytes,1,opt,name=bls_sig,json=blsSig,proto3" json:"bls_sig,omitempty"`
}

func (m *MsgAddBlsSig) Reset()         { *m = MsgAddBlsSig{} }
func (m *MsgAddBlsSig) String() string { return proto.CompactTextString(m) }
func (*MsgAddBlsSig) ProtoMessage()    {}
func (*MsgAddBlsSig) Descriptor() ([]byte, []int) {
	return fileDescriptor_24b023a97b92daa6, []int{0}
}
func (m *MsgAddBlsSig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddBlsSig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddBlsSig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddBlsSig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddBlsSig.Merge(m, src)
}
func (m *MsgAddBlsSig) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddBlsSig) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddBlsSig.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddBlsSig proto.InternalMessageInfo

// MsgAddBlsSigResponse defines the MsgAddBlsSig response type.
type MsgAddBlsSigResponse struct {
}

func (m *MsgAddBlsSigResponse) Reset()         { *m = MsgAddBlsSigResponse{} }
func (m *MsgAddBlsSigResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddBlsSigResponse) ProtoMessage()    {}
func (*MsgAddBlsSigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_24b023a97b92daa6, []int{1}
}
func (m *MsgAddBlsSigResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddBlsSigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddBlsSigResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddBlsSigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddBlsSigResponse.Merge(m, src)
}
func (m *MsgAddBlsSigResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddBlsSigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddBlsSigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddBlsSigResponse proto.InternalMessageInfo

// MsgWrappedCreateValidator defines a wrapped message to create a validator
type MsgWrappedCreateValidator struct {
	Key                *BlsKey                   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	MsgCreateValidator *types.MsgCreateValidator `protobuf:"bytes,2,opt,name=msg_create_validator,json=msgCreateValidator,proto3" json:"msg_create_validator,omitempty"`
}

func (m *MsgWrappedCreateValidator) Reset()         { *m = MsgWrappedCreateValidator{} }
func (m *MsgWrappedCreateValidator) String() string { return proto.CompactTextString(m) }
func (*MsgWrappedCreateValidator) ProtoMessage()    {}
func (*MsgWrappedCreateValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_24b023a97b92daa6, []int{2}
}
func (m *MsgWrappedCreateValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWrappedCreateValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWrappedCreateValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWrappedCreateValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWrappedCreateValidator.Merge(m, src)
}
func (m *MsgWrappedCreateValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgWrappedCreateValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWrappedCreateValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWrappedCreateValidator proto.InternalMessageInfo

// MsgWrappedCreateValidatorResponse defines the MsgWrappedCreateValidator response type
type MsgWrappedCreateValidatorResponse struct {
}

func (m *MsgWrappedCreateValidatorResponse) Reset()         { *m = MsgWrappedCreateValidatorResponse{} }
func (m *MsgWrappedCreateValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWrappedCreateValidatorResponse) ProtoMessage()    {}
func (*MsgWrappedCreateValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_24b023a97b92daa6, []int{3}
}
func (m *MsgWrappedCreateValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWrappedCreateValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWrappedCreateValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWrappedCreateValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWrappedCreateValidatorResponse.Merge(m, src)
}
func (m *MsgWrappedCreateValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWrappedCreateValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWrappedCreateValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWrappedCreateValidatorResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddBlsSig)(nil), "babylon.checkpointing.v1.MsgAddBlsSig")
	proto.RegisterType((*MsgAddBlsSigResponse)(nil), "babylon.checkpointing.v1.MsgAddBlsSigResponse")
	proto.RegisterType((*MsgWrappedCreateValidator)(nil), "babylon.checkpointing.v1.MsgWrappedCreateValidator")
	proto.RegisterType((*MsgWrappedCreateValidatorResponse)(nil), "babylon.checkpointing.v1.MsgWrappedCreateValidatorResponse")
}

func init() { proto.RegisterFile("babylon/checkpointing/tx.proto", fileDescriptor_24b023a97b92daa6) }

var fileDescriptor_24b023a97b92daa6 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x4a, 0x4c, 0xaa,
	0xcc, 0xc9, 0xcf, 0xd3, 0x4f, 0xce, 0x48, 0x4d, 0xce, 0x2e, 0xc8, 0xcf, 0xcc, 0x2b, 0xc9, 0xcc,
	0x4b, 0xd7, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x80, 0xca, 0xeb, 0xa1,
	0xc8, 0xeb, 0x95, 0x19, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x15, 0xe9, 0x83, 0x58, 0x10,
	0xf5, 0x52, 0x6a, 0xd8, 0xcd, 0x43, 0xf0, 0xa0, 0xea, 0x94, 0xb1, 0xab, 0x4b, 0xca, 0x29, 0x8e,
	0xcf, 0x4e, 0xad, 0x84, 0x2a, 0x92, 0x4f, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0x2f, 0x2e, 0x49,
	0xcc, 0x06, 0xc9, 0x96, 0x19, 0x26, 0xa5, 0x96, 0x24, 0x1a, 0xc2, 0x5d, 0xa7, 0x14, 0xcc, 0xc5,
	0xe3, 0x5b, 0x9c, 0xee, 0x98, 0x92, 0xe2, 0x94, 0x53, 0x1c, 0x9c, 0x99, 0x2e, 0x64, 0xc9, 0xc5,
	0x0e, 0x32, 0xa1, 0x38, 0x33, 0x5d, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x41, 0x0f, 0x97,
	0xfb, 0xf5, 0x20, 0x5a, 0x82, 0xd8, 0x92, 0xc0, 0xb4, 0x15, 0x47, 0xc7, 0x02, 0x79, 0x86, 0x17,
	0x0b, 0xe4, 0x19, 0x94, 0xc4, 0xb8, 0x44, 0x90, 0x0d, 0x0d, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b,
	0x4e, 0x55, 0xda, 0xce, 0xc8, 0x25, 0xe9, 0x5b, 0x9c, 0x1e, 0x5e, 0x94, 0x58, 0x50, 0x90, 0x9a,
	0xe2, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x1a, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58, 0x92, 0x5f, 0x24,
	0x64, 0xc4, 0xc5, 0x9c, 0x9d, 0x5a, 0x49, 0x94, 0xb5, 0xde, 0xa9, 0x95, 0x41, 0x20, 0xc5, 0x42,
	0x31, 0x5c, 0x22, 0xb9, 0xc5, 0xe9, 0xf1, 0xc9, 0x60, 0xa3, 0xe2, 0xcb, 0x60, 0x66, 0x49, 0x30,
	0x81, 0x0d, 0xd1, 0xd2, 0x83, 0x78, 0x5f, 0x0f, 0xea, 0x7d, 0x3d, 0xa8, 0xf7, 0xf5, 0x7c, 0x8b,
	0xd3, 0xd1, 0x6c, 0x0f, 0x12, 0xca, 0xc5, 0x10, 0x43, 0xf2, 0x91, 0x32, 0x97, 0x22, 0x4e, 0x87,
	0xc3, 0xbc, 0x67, 0xf4, 0x9d, 0x91, 0x8b, 0xd9, 0xb7, 0x38, 0x5d, 0x28, 0x99, 0x8b, 0x13, 0x11,
	0xa0, 0x6a, 0xb8, 0x3d, 0x82, 0x1c, 0x46, 0x52, 0x7a, 0xc4, 0xa9, 0x83, 0x59, 0x26, 0xd4, 0xc5,
	0xc8, 0x25, 0x86, 0x23, 0x20, 0x8d, 0xf1, 0x1a, 0x85, 0x5d, 0x93, 0x94, 0x35, 0x19, 0x9a, 0x60,
	0x8e, 0x71, 0xf2, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xd3, 0xf4,
	0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xa8, 0x05, 0xc9, 0x19, 0x89, 0x99,
	0x79, 0x30, 0x8e, 0x7e, 0x05, 0x7a, 0xbe, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xa7, 0x4e,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xe9, 0x81, 0x51, 0x5d, 0x03, 0x00, 0x00,
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
	// AddBlsSig defines a method for accumulating BLS signatures
	AddBlsSig(ctx context.Context, in *MsgAddBlsSig, opts ...grpc.CallOption) (*MsgAddBlsSigResponse, error)
	// WrappedCreateValidator defines a method for registering a new validator
	WrappedCreateValidator(ctx context.Context, in *MsgWrappedCreateValidator, opts ...grpc.CallOption) (*MsgWrappedCreateValidatorResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddBlsSig(ctx context.Context, in *MsgAddBlsSig, opts ...grpc.CallOption) (*MsgAddBlsSigResponse, error) {
	out := new(MsgAddBlsSigResponse)
	err := c.cc.Invoke(ctx, "/babylon.checkpointing.v1.Msg/AddBlsSig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WrappedCreateValidator(ctx context.Context, in *MsgWrappedCreateValidator, opts ...grpc.CallOption) (*MsgWrappedCreateValidatorResponse, error) {
	out := new(MsgWrappedCreateValidatorResponse)
	err := c.cc.Invoke(ctx, "/babylon.checkpointing.v1.Msg/WrappedCreateValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// AddBlsSig defines a method for accumulating BLS signatures
	AddBlsSig(context.Context, *MsgAddBlsSig) (*MsgAddBlsSigResponse, error)
	// WrappedCreateValidator defines a method for registering a new validator
	WrappedCreateValidator(context.Context, *MsgWrappedCreateValidator) (*MsgWrappedCreateValidatorResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddBlsSig(ctx context.Context, req *MsgAddBlsSig) (*MsgAddBlsSigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlsSig not implemented")
}
func (*UnimplementedMsgServer) WrappedCreateValidator(ctx context.Context, req *MsgWrappedCreateValidator) (*MsgWrappedCreateValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WrappedCreateValidator not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddBlsSig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddBlsSig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddBlsSig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/babylon.checkpointing.v1.Msg/AddBlsSig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddBlsSig(ctx, req.(*MsgAddBlsSig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WrappedCreateValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWrappedCreateValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WrappedCreateValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/babylon.checkpointing.v1.Msg/WrappedCreateValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WrappedCreateValidator(ctx, req.(*MsgWrappedCreateValidator))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babylon.checkpointing.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBlsSig",
			Handler:    _Msg_AddBlsSig_Handler,
		},
		{
			MethodName: "WrappedCreateValidator",
			Handler:    _Msg_WrappedCreateValidator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "babylon/checkpointing/tx.proto",
}

func (m *MsgAddBlsSig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddBlsSig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddBlsSig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlsSig != nil {
		{
			size, err := m.BlsSig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddBlsSigResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddBlsSigResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddBlsSigResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgWrappedCreateValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWrappedCreateValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWrappedCreateValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MsgCreateValidator != nil {
		{
			size, err := m.MsgCreateValidator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Key != nil {
		{
			size, err := m.Key.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWrappedCreateValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWrappedCreateValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWrappedCreateValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgAddBlsSig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlsSig != nil {
		l = m.BlsSig.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddBlsSigResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgWrappedCreateValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.MsgCreateValidator != nil {
		l = m.MsgCreateValidator.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgWrappedCreateValidatorResponse) Size() (n int) {
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
func (m *MsgAddBlsSig) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddBlsSig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddBlsSig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlsSig", wireType)
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
			if m.BlsSig == nil {
				m.BlsSig = &BlsSig{}
			}
			if err := m.BlsSig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgAddBlsSigResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddBlsSigResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddBlsSigResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgWrappedCreateValidator) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWrappedCreateValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWrappedCreateValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
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
			if m.Key == nil {
				m.Key = &BlsKey{}
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgCreateValidator", wireType)
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
			if m.MsgCreateValidator == nil {
				m.MsgCreateValidator = &types.MsgCreateValidator{}
			}
			if err := m.MsgCreateValidator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgWrappedCreateValidatorResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWrappedCreateValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWrappedCreateValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
