// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/finality/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_babylonchain_babylon_types "github.com/babylonchain/babylon/types"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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

// MsgAddFinalitySig defines a message for adding a finality vote
type MsgAddFinalitySig struct {
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	// fp_btc_pk is the BTC PK of the finality provider that casts this vote
	FpBtcPk *github_com_babylonchain_babylon_types.BIP340PubKey `protobuf:"bytes,2,opt,name=fp_btc_pk,json=fpBtcPk,proto3,customtype=github.com/babylonchain/babylon/types.BIP340PubKey" json:"fp_btc_pk,omitempty"`
	// block_height is the height of the voted block
	BlockHeight uint64 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// block_app_hash is the AppHash of the voted block
	BlockAppHash []byte `protobuf:"bytes,4,opt,name=block_app_hash,json=blockAppHash,proto3" json:"block_app_hash,omitempty"`
	// finality_sig is the finality signature to this block
	// where finality signature is an EOTS signature, i.e.,
	// the `s` in a Schnorr signature `(r, s)`
	// `r` is the public randomness that is already committed by the finality provider
	FinalitySig *github_com_babylonchain_babylon_types.SchnorrEOTSSig `protobuf:"bytes,5,opt,name=finality_sig,json=finalitySig,proto3,customtype=github.com/babylonchain/babylon/types.SchnorrEOTSSig" json:"finality_sig,omitempty"`
}

func (m *MsgAddFinalitySig) Reset()         { *m = MsgAddFinalitySig{} }
func (m *MsgAddFinalitySig) String() string { return proto.CompactTextString(m) }
func (*MsgAddFinalitySig) ProtoMessage()    {}
func (*MsgAddFinalitySig) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dd6da066b6baf1d, []int{0}
}
func (m *MsgAddFinalitySig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddFinalitySig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddFinalitySig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddFinalitySig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddFinalitySig.Merge(m, src)
}
func (m *MsgAddFinalitySig) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddFinalitySig) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddFinalitySig.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddFinalitySig proto.InternalMessageInfo

func (m *MsgAddFinalitySig) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgAddFinalitySig) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *MsgAddFinalitySig) GetBlockAppHash() []byte {
	if m != nil {
		return m.BlockAppHash
	}
	return nil
}

// MsgAddFinalitySigResponse is the response to the MsgAddFinalitySig message
type MsgAddFinalitySigResponse struct {
}

func (m *MsgAddFinalitySigResponse) Reset()         { *m = MsgAddFinalitySigResponse{} }
func (m *MsgAddFinalitySigResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddFinalitySigResponse) ProtoMessage()    {}
func (*MsgAddFinalitySigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dd6da066b6baf1d, []int{1}
}
func (m *MsgAddFinalitySigResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddFinalitySigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddFinalitySigResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddFinalitySigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddFinalitySigResponse.Merge(m, src)
}
func (m *MsgAddFinalitySigResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddFinalitySigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddFinalitySigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddFinalitySigResponse proto.InternalMessageInfo

// MsgUpdateParams defines a message for updating finality module parameters.
type MsgUpdateParams struct {
	// authority is the address of the governance account.
	// just FYI: cosmos.AddressString marks that this field should use type alias
	// for AddressString instead of string, but the functionality is not yet implemented
	// in cosmos-proto
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the finality parameters to update.
	//
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dd6da066b6baf1d, []int{2}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse is the response to the MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dd6da066b6baf1d, []int{3}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddFinalitySig)(nil), "babylon.finality.v1.MsgAddFinalitySig")
	proto.RegisterType((*MsgAddFinalitySigResponse)(nil), "babylon.finality.v1.MsgAddFinalitySigResponse")
	proto.RegisterType((*MsgUpdateParams)(nil), "babylon.finality.v1.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "babylon.finality.v1.MsgUpdateParamsResponse")
}

func init() { proto.RegisterFile("babylon/finality/v1/tx.proto", fileDescriptor_2dd6da066b6baf1d) }

var fileDescriptor_2dd6da066b6baf1d = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0x6d, 0xf6, 0xa7, 0xa8, 0x6e, 0x55, 0x44, 0x98, 0x58, 0xda, 0xa1, 0xb4, 0x54, 0x13, 0xaa,
	0x26, 0x48, 0xb6, 0x6e, 0x9a, 0x60, 0xb7, 0x46, 0x02, 0x0d, 0x50, 0x45, 0x95, 0xc2, 0x05, 0x0e,
	0x91, 0x93, 0xa6, 0xb6, 0xd5, 0x36, 0xb6, 0x62, 0x77, 0x5a, 0x6e, 0x88, 0x4f, 0xc0, 0x81, 0x0f,
	0x32, 0x21, 0x3e, 0xc4, 0x8e, 0x13, 0x27, 0xd4, 0x43, 0x85, 0xda, 0xc3, 0xbe, 0x06, 0x6a, 0xe2,
	0xa8, 0xb0, 0x15, 0xd1, 0x9b, 0x7f, 0x7e, 0xef, 0xe7, 0xf7, 0xfc, 0xfc, 0x33, 0x78, 0xe8, 0x42,
	0x37, 0x1a, 0xd0, 0xc0, 0xec, 0x91, 0x00, 0x0e, 0x88, 0x88, 0xcc, 0xb3, 0x03, 0x53, 0x9c, 0x1b,
	0x2c, 0xa4, 0x82, 0xaa, 0xf7, 0x25, 0x6a, 0xa4, 0xa8, 0x71, 0x76, 0x50, 0xde, 0x42, 0x14, 0xd1,
	0x18, 0x37, 0xe7, 0xab, 0x84, 0x5a, 0x2e, 0x79, 0x94, 0x0f, 0x29, 0x77, 0x12, 0x20, 0x29, 0x24,
	0xb4, 0x9d, 0x54, 0xe6, 0x90, 0xa3, 0xf9, 0xe9, 0x43, 0x8e, 0x24, 0x50, 0x5d, 0x26, 0xce, 0x60,
	0x08, 0x87, 0xb2, 0xb5, 0xf6, 0x6d, 0x0d, 0xdc, 0x6b, 0x71, 0xd4, 0xec, 0x76, 0x5f, 0x4a, 0x4a,
	0x87, 0x20, 0xf5, 0x01, 0xc8, 0x72, 0x82, 0x02, 0x3f, 0xd4, 0x94, 0xaa, 0x52, 0xcf, 0xd9, 0xb2,
	0x52, 0x6d, 0x90, 0xeb, 0x31, 0xc7, 0x15, 0x9e, 0xc3, 0xfa, 0xda, 0x5a, 0x55, 0xa9, 0x17, 0xac,
	0xe3, 0xf1, 0xa4, 0xd2, 0x40, 0x44, 0xe0, 0x91, 0x6b, 0x78, 0x74, 0x68, 0x4a, 0x45, 0x0f, 0x43,
	0x12, 0xa4, 0x85, 0x29, 0x22, 0xe6, 0x73, 0xc3, 0x7a, 0xd5, 0x3e, 0x3c, 0xda, 0x6f, 0x8f, 0xdc,
	0x37, 0x7e, 0x64, 0xdf, 0xe9, 0x31, 0x4b, 0x78, 0xed, 0xbe, 0xfa, 0x08, 0x14, 0xdc, 0x01, 0xf5,
	0xfa, 0x0e, 0xf6, 0x09, 0xc2, 0x42, 0x5b, 0xaf, 0x2a, 0xf5, 0x0d, 0x3b, 0x1f, 0xef, 0x9d, 0xc6,
	0x5b, 0xea, 0x2e, 0x28, 0x26, 0x14, 0xc8, 0x98, 0x83, 0x21, 0xc7, 0xda, 0xc6, 0x5c, 0xdb, 0x4e,
	0x1a, 0x9b, 0x8c, 0x9d, 0x42, 0x8e, 0xd5, 0x8f, 0xa0, 0x90, 0x5e, 0xd3, 0xe1, 0x04, 0x69, 0x9b,
	0xb1, 0xbf, 0x67, 0xe3, 0x49, 0xe5, 0x68, 0x35, 0x7f, 0x1d, 0x0f, 0x07, 0x34, 0x0c, 0x5f, 0xbc,
	0x7d, 0xd7, 0xe9, 0x10, 0x64, 0xe7, 0x7b, 0x8b, 0x44, 0x4e, 0xf2, 0x9f, 0xaf, 0x2f, 0xf6, 0x64,
	0x0c, 0xb5, 0x1d, 0x50, 0xba, 0x95, 0x99, 0xed, 0x73, 0x46, 0x03, 0xee, 0xd7, 0xbe, 0x2a, 0xe0,
	0x6e, 0x8b, 0xa3, 0xf7, 0xac, 0x0b, 0x85, 0xdf, 0x8e, 0xb3, 0x56, 0x8f, 0x41, 0x0e, 0x8e, 0x04,
	0xa6, 0x21, 0x11, 0x51, 0x12, 0xa9, 0xa5, 0xfd, 0xf8, 0xfe, 0x74, 0x4b, 0xbe, 0x62, 0xb3, 0xdb,
	0x0d, 0x7d, 0xce, 0x3b, 0x22, 0x24, 0x01, 0xb2, 0x17, 0x54, 0xf5, 0x39, 0xc8, 0x26, 0xaf, 0x15,
	0x87, 0x9d, 0x6f, 0xec, 0x18, 0x4b, 0xe6, 0xc5, 0x48, 0x44, 0xac, 0x8d, 0xcb, 0x49, 0x25, 0x63,
	0xcb, 0x86, 0x93, 0xe2, 0xdc, 0xf0, 0xe2, 0xa8, 0x5a, 0x09, 0x6c, 0xdf, 0x70, 0x95, 0x3a, 0x6e,
	0x8c, 0x15, 0xb0, 0xde, 0xe2, 0x48, 0xc5, 0xa0, 0x78, 0x63, 0x0e, 0x1e, 0x2f, 0xd5, 0xbb, 0x75,
	0xf7, 0xb2, 0xb1, 0x1a, 0x2f, 0x55, 0x54, 0x5d, 0x50, 0xf8, 0x2b, 0x9f, 0xdd, 0x7f, 0xf5, 0xff,
	0xc9, 0x2a, 0x3f, 0x59, 0x85, 0x95, 0x6a, 0x94, 0x37, 0x3f, 0x5d, 0x5f, 0xec, 0x29, 0xd6, 0xeb,
	0xcb, 0xa9, 0xae, 0x5c, 0x4d, 0x75, 0xe5, 0xd7, 0x54, 0x57, 0xbe, 0xcc, 0xf4, 0xcc, 0xd5, 0x4c,
	0xcf, 0xfc, 0x9c, 0xe9, 0x99, 0x0f, 0xfb, 0xff, 0x9b, 0x8a, 0xf3, 0xc5, 0xb7, 0x89, 0x07, 0xc4,
	0xcd, 0xc6, 0x7f, 0xe6, 0xf0, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x9d, 0x41, 0x5a, 0xd4,
	0x03, 0x00, 0x00,
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
	// AddFinalitySig adds a finality signature to a given block
	AddFinalitySig(ctx context.Context, in *MsgAddFinalitySig, opts ...grpc.CallOption) (*MsgAddFinalitySigResponse, error)
	// TODO: msg for evidence of equivocation. this is not specified yet
	// UpdateParams updates the finality module parameters.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddFinalitySig(ctx context.Context, in *MsgAddFinalitySig, opts ...grpc.CallOption) (*MsgAddFinalitySigResponse, error) {
	out := new(MsgAddFinalitySigResponse)
	err := c.cc.Invoke(ctx, "/babylon.finality.v1.Msg/AddFinalitySig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/babylon.finality.v1.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// AddFinalitySig adds a finality signature to a given block
	AddFinalitySig(context.Context, *MsgAddFinalitySig) (*MsgAddFinalitySigResponse, error)
	// TODO: msg for evidence of equivocation. this is not specified yet
	// UpdateParams updates the finality module parameters.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddFinalitySig(ctx context.Context, req *MsgAddFinalitySig) (*MsgAddFinalitySigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFinalitySig not implemented")
}
func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddFinalitySig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddFinalitySig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddFinalitySig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/babylon.finality.v1.Msg/AddFinalitySig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddFinalitySig(ctx, req.(*MsgAddFinalitySig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/babylon.finality.v1.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babylon.finality.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFinalitySig",
			Handler:    _Msg_AddFinalitySig_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "babylon/finality/v1/tx.proto",
}

func (m *MsgAddFinalitySig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddFinalitySig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddFinalitySig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.BlockAppHash) > 0 {
		i -= len(m.BlockAppHash)
		copy(dAtA[i:], m.BlockAppHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BlockAppHash)))
		i--
		dAtA[i] = 0x22
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.FpBtcPk != nil {
		{
			size := m.FpBtcPk.Size()
			i -= size
			if _, err := m.FpBtcPk.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddFinalitySigResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddFinalitySigResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddFinalitySigResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgAddFinalitySig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.FpBtcPk != nil {
		l = m.FpBtcPk.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTx(uint64(m.BlockHeight))
	}
	l = len(m.BlockAppHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.FinalitySig != nil {
		l = m.FinalitySig.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddFinalitySigResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
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
func (m *MsgAddFinalitySig) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddFinalitySig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddFinalitySig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FpBtcPk", wireType)
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
			var v github_com_babylonchain_babylon_types.BIP340PubKey
			m.FpBtcPk = &v
			if err := m.FpBtcPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockAppHash", wireType)
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
			m.BlockAppHash = append(m.BlockAppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockAppHash == nil {
				m.BlockAppHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalitySig", wireType)
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
			var v github_com_babylonchain_babylon_types.SchnorrEOTSSig
			m.FinalitySig = &v
			if err := m.FinalitySig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgAddFinalitySigResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddFinalitySigResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddFinalitySigResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
