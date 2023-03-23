// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/monitor/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryEndedEpochBtcHeightRequest defines a query type for EndedEpochBtcHeight
// RPC method
type QueryEndedEpochBtcHeightRequest struct {
	EpochNum uint64 `protobuf:"varint,1,opt,name=epoch_num,json=epochNum,proto3" json:"epoch_num,omitempty"`
}

func (m *QueryEndedEpochBtcHeightRequest) Reset()         { *m = QueryEndedEpochBtcHeightRequest{} }
func (m *QueryEndedEpochBtcHeightRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEndedEpochBtcHeightRequest) ProtoMessage()    {}
func (*QueryEndedEpochBtcHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8aafb034c55a8f2, []int{0}
}
func (m *QueryEndedEpochBtcHeightRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEndedEpochBtcHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEndedEpochBtcHeightRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEndedEpochBtcHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEndedEpochBtcHeightRequest.Merge(m, src)
}
func (m *QueryEndedEpochBtcHeightRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEndedEpochBtcHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEndedEpochBtcHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEndedEpochBtcHeightRequest proto.InternalMessageInfo

func (m *QueryEndedEpochBtcHeightRequest) GetEpochNum() uint64 {
	if m != nil {
		return m.EpochNum
	}
	return 0
}

// QueryEndedEpochBtcHeightResponse defines a response type for
// EndedEpochBtcHeight RPC method
type QueryEndedEpochBtcHeightResponse struct {
	// height of btc light client when epoch ended
	BtcLightClientHeight uint64 `protobuf:"varint,1,opt,name=btc_light_client_height,json=btcLightClientHeight,proto3" json:"btc_light_client_height,omitempty"`
}

func (m *QueryEndedEpochBtcHeightResponse) Reset()         { *m = QueryEndedEpochBtcHeightResponse{} }
func (m *QueryEndedEpochBtcHeightResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEndedEpochBtcHeightResponse) ProtoMessage()    {}
func (*QueryEndedEpochBtcHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8aafb034c55a8f2, []int{1}
}
func (m *QueryEndedEpochBtcHeightResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEndedEpochBtcHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEndedEpochBtcHeightResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEndedEpochBtcHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEndedEpochBtcHeightResponse.Merge(m, src)
}
func (m *QueryEndedEpochBtcHeightResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEndedEpochBtcHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEndedEpochBtcHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEndedEpochBtcHeightResponse proto.InternalMessageInfo

func (m *QueryEndedEpochBtcHeightResponse) GetBtcLightClientHeight() uint64 {
	if m != nil {
		return m.BtcLightClientHeight
	}
	return 0
}

// QueryReportedCheckpointBtcHeightRequest defines a query type for
// ReportedCheckpointBtcHeight RPC method
type QueryReportedCheckpointBtcHeightRequest struct {
	// ckpt_hash is hex encoded byte string of the hash of the checkpoint
	CkptHash string `protobuf:"bytes,1,opt,name=ckpt_hash,json=ckptHash,proto3" json:"ckpt_hash,omitempty"`
}

func (m *QueryReportedCheckpointBtcHeightRequest) Reset() {
	*m = QueryReportedCheckpointBtcHeightRequest{}
}
func (m *QueryReportedCheckpointBtcHeightRequest) String() string { return proto.CompactTextString(m) }
func (*QueryReportedCheckpointBtcHeightRequest) ProtoMessage()    {}
func (*QueryReportedCheckpointBtcHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8aafb034c55a8f2, []int{2}
}
func (m *QueryReportedCheckpointBtcHeightRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryReportedCheckpointBtcHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryReportedCheckpointBtcHeightRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryReportedCheckpointBtcHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReportedCheckpointBtcHeightRequest.Merge(m, src)
}
func (m *QueryReportedCheckpointBtcHeightRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryReportedCheckpointBtcHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReportedCheckpointBtcHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReportedCheckpointBtcHeightRequest proto.InternalMessageInfo

func (m *QueryReportedCheckpointBtcHeightRequest) GetCkptHash() string {
	if m != nil {
		return m.CkptHash
	}
	return ""
}

// QueryReportedCheckpointBtcHeightResponse defines a response type for
// ReportedCheckpointBtcHeight RPC method
type QueryReportedCheckpointBtcHeightResponse struct {
	// height of btc light client when checkpoint is reported
	BtcLightClientHeight uint64 `protobuf:"varint,1,opt,name=btc_light_client_height,json=btcLightClientHeight,proto3" json:"btc_light_client_height,omitempty"`
}

func (m *QueryReportedCheckpointBtcHeightResponse) Reset() {
	*m = QueryReportedCheckpointBtcHeightResponse{}
}
func (m *QueryReportedCheckpointBtcHeightResponse) String() string { return proto.CompactTextString(m) }
func (*QueryReportedCheckpointBtcHeightResponse) ProtoMessage()    {}
func (*QueryReportedCheckpointBtcHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8aafb034c55a8f2, []int{3}
}
func (m *QueryReportedCheckpointBtcHeightResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryReportedCheckpointBtcHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryReportedCheckpointBtcHeightResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryReportedCheckpointBtcHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReportedCheckpointBtcHeightResponse.Merge(m, src)
}
func (m *QueryReportedCheckpointBtcHeightResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryReportedCheckpointBtcHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReportedCheckpointBtcHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReportedCheckpointBtcHeightResponse proto.InternalMessageInfo

func (m *QueryReportedCheckpointBtcHeightResponse) GetBtcLightClientHeight() uint64 {
	if m != nil {
		return m.BtcLightClientHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryEndedEpochBtcHeightRequest)(nil), "babylon.monitor.v1.QueryEndedEpochBtcHeightRequest")
	proto.RegisterType((*QueryEndedEpochBtcHeightResponse)(nil), "babylon.monitor.v1.QueryEndedEpochBtcHeightResponse")
	proto.RegisterType((*QueryReportedCheckpointBtcHeightRequest)(nil), "babylon.monitor.v1.QueryReportedCheckpointBtcHeightRequest")
	proto.RegisterType((*QueryReportedCheckpointBtcHeightResponse)(nil), "babylon.monitor.v1.QueryReportedCheckpointBtcHeightResponse")
}

func init() { proto.RegisterFile("babylon/monitor/v1/query.proto", fileDescriptor_a8aafb034c55a8f2) }

var fileDescriptor_a8aafb034c55a8f2 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x1b, 0xff, 0xb1, 0x3b, 0xc7, 0x71, 0x41, 0xc9, 0x4a, 0x5c, 0x72, 0xd0, 0x82, 0x98,
	0x61, 0xad, 0x9e, 0x14, 0x0f, 0xbb, 0xac, 0x54, 0x10, 0xc1, 0xdc, 0xf4, 0x12, 0x26, 0xd3, 0x21,
	0x33, 0x34, 0x99, 0x77, 0x9a, 0x99, 0x14, 0x4b, 0xe9, 0xc5, 0xa3, 0x27, 0xc1, 0x2f, 0xe2, 0xc7,
	0xf0, 0x22, 0x14, 0xbc, 0x78, 0x94, 0xd6, 0x0f, 0x22, 0x99, 0xc4, 0x5c, 0x4c, 0x2d, 0xea, 0x2d,
	0x79, 0x9f, 0x79, 0x7e, 0x79, 0xdf, 0xf7, 0xc9, 0xa0, 0x20, 0xa5, 0xe9, 0x22, 0x07, 0x45, 0x0a,
	0x50, 0xd2, 0x42, 0x49, 0xe6, 0xa7, 0x64, 0x56, 0xf1, 0x72, 0x11, 0xe9, 0x12, 0x2c, 0x60, 0xdc,
	0xea, 0x51, 0xab, 0x47, 0xf3, 0x53, 0xff, 0x28, 0x83, 0x0c, 0x9c, 0x4c, 0xea, 0xa7, 0xe6, 0xa4,
	0x7f, 0x2b, 0x03, 0xc8, 0x72, 0x4e, 0xa8, 0x96, 0x84, 0x2a, 0x05, 0x96, 0x5a, 0x09, 0xca, 0x34,
	0x6a, 0xf8, 0x14, 0xdd, 0x7e, 0x55, 0x63, 0x2f, 0xd4, 0x84, 0x4f, 0x2e, 0x34, 0x30, 0x71, 0x66,
	0xd9, 0x98, 0xcb, 0x4c, 0xd8, 0x98, 0xcf, 0x2a, 0x6e, 0x2c, 0x3e, 0x46, 0x87, 0xbc, 0x16, 0x12,
	0x55, 0x15, 0x37, 0xbd, 0x13, 0x6f, 0x78, 0x25, 0x3e, 0x70, 0x85, 0x97, 0x55, 0x11, 0xbe, 0x46,
	0x27, 0xbb, 0xfd, 0x46, 0x83, 0x32, 0x1c, 0x3f, 0x42, 0x37, 0x52, 0xcb, 0x92, 0xbc, 0x2e, 0x26,
	0x2c, 0x97, 0x5c, 0xd9, 0x44, 0xb8, 0x23, 0x2d, 0xee, 0x28, 0xb5, 0xec, 0x45, 0xfd, 0x7e, 0xee,
	0xc4, 0xc6, 0x1e, 0x3e, 0x43, 0x77, 0x1d, 0x3a, 0xe6, 0x1a, 0x4a, 0xcb, 0x27, 0xe7, 0x82, 0xb3,
	0xa9, 0x06, 0xa9, 0x6c, 0x5f, 0x8b, 0x6c, 0xaa, 0x6d, 0x22, 0xa8, 0x11, 0x8e, 0x79, 0x18, 0x1f,
	0xd4, 0x85, 0x31, 0x35, 0x22, 0xa4, 0x68, 0xb8, 0x9f, 0xf3, 0x5f, 0xad, 0x3e, 0x78, 0x7f, 0x19,
	0x5d, 0x75, 0xdf, 0xc0, 0x9f, 0x3c, 0x74, 0xbd, 0x67, 0x17, 0x78, 0x14, 0xfd, 0x1e, 0x58, 0xb4,
	0x67, 0xf3, 0xfe, 0xc3, 0xbf, 0x33, 0x35, 0x33, 0x84, 0xd1, 0xbb, 0xaf, 0x3f, 0x3e, 0x5e, 0x1a,
	0xe2, 0x3b, 0xa4, 0xe7, 0x1f, 0x72, 0xc1, 0x19, 0xb2, 0xec, 0x12, 0x5d, 0xe1, 0x2f, 0x1e, 0x3a,
	0xfe, 0xc3, 0x6e, 0xf0, 0xe3, 0x9d, 0x5d, 0xec, 0x4f, 0xc6, 0x7f, 0xf2, 0x6f, 0xe6, 0x76, 0x94,
	0x91, 0x1b, 0xe5, 0x3e, 0xbe, 0xd7, 0x37, 0x0a, 0xeb, 0x8c, 0x86, 0x2c, 0xbb, 0xf8, 0x57, 0x67,
	0xcf, 0x3f, 0x6f, 0x02, 0x6f, 0xbd, 0x09, 0xbc, 0xef, 0x9b, 0xc0, 0xfb, 0xb0, 0x0d, 0x06, 0xeb,
	0x6d, 0x30, 0xf8, 0xb6, 0x0d, 0x06, 0x6f, 0x48, 0x26, 0xad, 0xa8, 0xd2, 0x88, 0x41, 0xf1, 0x0b,
	0xc8, 0x04, 0x95, 0xaa, 0xa3, 0xbf, 0xed, 0xf8, 0x76, 0xa1, 0xb9, 0x49, 0xaf, 0xb9, 0x4b, 0x32,
	0xfa, 0x19, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x9f, 0x08, 0x3d, 0x8e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// EndedEpochBtcHeight returns the BTC light client height at provided epoch
	// finish
	EndedEpochBtcHeight(ctx context.Context, in *QueryEndedEpochBtcHeightRequest, opts ...grpc.CallOption) (*QueryEndedEpochBtcHeightResponse, error)
	// ReportedCheckpointBtcHeight returns the BTC light client height at which
	// the checkpoint with the given hash is reported back to Babylon
	ReportedCheckpointBtcHeight(ctx context.Context, in *QueryReportedCheckpointBtcHeightRequest, opts ...grpc.CallOption) (*QueryReportedCheckpointBtcHeightResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) EndedEpochBtcHeight(ctx context.Context, in *QueryEndedEpochBtcHeightRequest, opts ...grpc.CallOption) (*QueryEndedEpochBtcHeightResponse, error) {
	out := new(QueryEndedEpochBtcHeightResponse)
	err := c.cc.Invoke(ctx, "/babylon.monitor.v1.Query/EndedEpochBtcHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ReportedCheckpointBtcHeight(ctx context.Context, in *QueryReportedCheckpointBtcHeightRequest, opts ...grpc.CallOption) (*QueryReportedCheckpointBtcHeightResponse, error) {
	out := new(QueryReportedCheckpointBtcHeightResponse)
	err := c.cc.Invoke(ctx, "/babylon.monitor.v1.Query/ReportedCheckpointBtcHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// EndedEpochBtcHeight returns the BTC light client height at provided epoch
	// finish
	EndedEpochBtcHeight(context.Context, *QueryEndedEpochBtcHeightRequest) (*QueryEndedEpochBtcHeightResponse, error)
	// ReportedCheckpointBtcHeight returns the BTC light client height at which
	// the checkpoint with the given hash is reported back to Babylon
	ReportedCheckpointBtcHeight(context.Context, *QueryReportedCheckpointBtcHeightRequest) (*QueryReportedCheckpointBtcHeightResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) EndedEpochBtcHeight(ctx context.Context, req *QueryEndedEpochBtcHeightRequest) (*QueryEndedEpochBtcHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndedEpochBtcHeight not implemented")
}
func (*UnimplementedQueryServer) ReportedCheckpointBtcHeight(ctx context.Context, req *QueryReportedCheckpointBtcHeightRequest) (*QueryReportedCheckpointBtcHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportedCheckpointBtcHeight not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_EndedEpochBtcHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEndedEpochBtcHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EndedEpochBtcHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/babylon.monitor.v1.Query/EndedEpochBtcHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EndedEpochBtcHeight(ctx, req.(*QueryEndedEpochBtcHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ReportedCheckpointBtcHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReportedCheckpointBtcHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ReportedCheckpointBtcHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/babylon.monitor.v1.Query/ReportedCheckpointBtcHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ReportedCheckpointBtcHeight(ctx, req.(*QueryReportedCheckpointBtcHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babylon.monitor.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EndedEpochBtcHeight",
			Handler:    _Query_EndedEpochBtcHeight_Handler,
		},
		{
			MethodName: "ReportedCheckpointBtcHeight",
			Handler:    _Query_ReportedCheckpointBtcHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "babylon/monitor/v1/query.proto",
}

func (m *QueryEndedEpochBtcHeightRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEndedEpochBtcHeightRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEndedEpochBtcHeightRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochNum != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.EpochNum))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryEndedEpochBtcHeightResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEndedEpochBtcHeightResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEndedEpochBtcHeightResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BtcLightClientHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BtcLightClientHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryReportedCheckpointBtcHeightRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryReportedCheckpointBtcHeightRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryReportedCheckpointBtcHeightRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CkptHash) > 0 {
		i -= len(m.CkptHash)
		copy(dAtA[i:], m.CkptHash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CkptHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryReportedCheckpointBtcHeightResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryReportedCheckpointBtcHeightResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryReportedCheckpointBtcHeightResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BtcLightClientHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BtcLightClientHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryEndedEpochBtcHeightRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochNum != 0 {
		n += 1 + sovQuery(uint64(m.EpochNum))
	}
	return n
}

func (m *QueryEndedEpochBtcHeightResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BtcLightClientHeight != 0 {
		n += 1 + sovQuery(uint64(m.BtcLightClientHeight))
	}
	return n
}

func (m *QueryReportedCheckpointBtcHeightRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CkptHash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryReportedCheckpointBtcHeightResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BtcLightClientHeight != 0 {
		n += 1 + sovQuery(uint64(m.BtcLightClientHeight))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryEndedEpochBtcHeightRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEndedEpochBtcHeightRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEndedEpochBtcHeightRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNum", wireType)
			}
			m.EpochNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEndedEpochBtcHeightResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEndedEpochBtcHeightResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEndedEpochBtcHeightResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcLightClientHeight", wireType)
			}
			m.BtcLightClientHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BtcLightClientHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryReportedCheckpointBtcHeightRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryReportedCheckpointBtcHeightRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryReportedCheckpointBtcHeightRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CkptHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CkptHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryReportedCheckpointBtcHeightResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryReportedCheckpointBtcHeightResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryReportedCheckpointBtcHeightResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcLightClientHeight", wireType)
			}
			m.BtcLightClientHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BtcLightClientHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
