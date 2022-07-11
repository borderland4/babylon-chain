// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/epoching/v1/epoching.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Epoch struct {
	EpochNumber          uint64 `protobuf:"varint,1,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
	CurrentEpochInterval uint64 `protobuf:"varint,2,opt,name=current_epoch_interval,json=currentEpochInterval,proto3" json:"current_epoch_interval,omitempty"`
	FirstBlockHeight     uint64 `protobuf:"varint,3,opt,name=first_block_height,json=firstBlockHeight,proto3" json:"first_block_height,omitempty"`
}

func (m *Epoch) Reset()         { *m = Epoch{} }
func (m *Epoch) String() string { return proto.CompactTextString(m) }
func (*Epoch) ProtoMessage()    {}
func (*Epoch) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2f209d5311f84c, []int{0}
}
func (m *Epoch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Epoch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Epoch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Epoch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Epoch.Merge(m, src)
}
func (m *Epoch) XXX_Size() int {
	return m.Size()
}
func (m *Epoch) XXX_DiscardUnknown() {
	xxx_messageInfo_Epoch.DiscardUnknown(m)
}

var xxx_messageInfo_Epoch proto.InternalMessageInfo

func (m *Epoch) GetEpochNumber() uint64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

func (m *Epoch) GetCurrentEpochInterval() uint64 {
	if m != nil {
		return m.CurrentEpochInterval
	}
	return 0
}

func (m *Epoch) GetFirstBlockHeight() uint64 {
	if m != nil {
		return m.FirstBlockHeight
	}
	return 0
}

// QueuedMessage is a message that can change the validator set and is delayed to the epoch boundary
type QueuedMessage struct {
	// tx_id is the ID of the tx that contains the message
	TxId []byte `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// msg_id is the original message ID, i.e., hash of the marshaled message
	MsgId []byte `protobuf:"bytes,2,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	// msg is the actual message that is sent by a user and is queued by the epoching module
	//
	// Types that are valid to be assigned to Msg:
	//	*QueuedMessage_MsgCreateValidator
	//	*QueuedMessage_MsgDelegate
	//	*QueuedMessage_MsgUndelegate
	//	*QueuedMessage_MsgBeginRedelegate
	Msg isQueuedMessage_Msg `protobuf_oneof:"msg"`
}

func (m *QueuedMessage) Reset()         { *m = QueuedMessage{} }
func (m *QueuedMessage) String() string { return proto.CompactTextString(m) }
func (*QueuedMessage) ProtoMessage()    {}
func (*QueuedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2f209d5311f84c, []int{1}
}
func (m *QueuedMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueuedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueuedMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueuedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueuedMessage.Merge(m, src)
}
func (m *QueuedMessage) XXX_Size() int {
	return m.Size()
}
func (m *QueuedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_QueuedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_QueuedMessage proto.InternalMessageInfo

type isQueuedMessage_Msg interface {
	isQueuedMessage_Msg()
	MarshalTo([]byte) (int, error)
	Size() int
}

type QueuedMessage_MsgCreateValidator struct {
	MsgCreateValidator *types.MsgCreateValidator `protobuf:"bytes,3,opt,name=msg_create_validator,json=msgCreateValidator,proto3,oneof" json:"msg_create_validator,omitempty"`
}
type QueuedMessage_MsgDelegate struct {
	MsgDelegate *types.MsgDelegate `protobuf:"bytes,4,opt,name=msg_delegate,json=msgDelegate,proto3,oneof" json:"msg_delegate,omitempty"`
}
type QueuedMessage_MsgUndelegate struct {
	MsgUndelegate *types.MsgUndelegate `protobuf:"bytes,5,opt,name=msg_undelegate,json=msgUndelegate,proto3,oneof" json:"msg_undelegate,omitempty"`
}
type QueuedMessage_MsgBeginRedelegate struct {
	MsgBeginRedelegate *types.MsgBeginRedelegate `protobuf:"bytes,6,opt,name=msg_begin_redelegate,json=msgBeginRedelegate,proto3,oneof" json:"msg_begin_redelegate,omitempty"`
}

func (*QueuedMessage_MsgCreateValidator) isQueuedMessage_Msg() {}
func (*QueuedMessage_MsgDelegate) isQueuedMessage_Msg()        {}
func (*QueuedMessage_MsgUndelegate) isQueuedMessage_Msg()      {}
func (*QueuedMessage_MsgBeginRedelegate) isQueuedMessage_Msg() {}

func (m *QueuedMessage) GetMsg() isQueuedMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *QueuedMessage) GetTxId() []byte {
	if m != nil {
		return m.TxId
	}
	return nil
}

func (m *QueuedMessage) GetMsgId() []byte {
	if m != nil {
		return m.MsgId
	}
	return nil
}

func (m *QueuedMessage) GetMsgCreateValidator() *types.MsgCreateValidator {
	if x, ok := m.GetMsg().(*QueuedMessage_MsgCreateValidator); ok {
		return x.MsgCreateValidator
	}
	return nil
}

func (m *QueuedMessage) GetMsgDelegate() *types.MsgDelegate {
	if x, ok := m.GetMsg().(*QueuedMessage_MsgDelegate); ok {
		return x.MsgDelegate
	}
	return nil
}

func (m *QueuedMessage) GetMsgUndelegate() *types.MsgUndelegate {
	if x, ok := m.GetMsg().(*QueuedMessage_MsgUndelegate); ok {
		return x.MsgUndelegate
	}
	return nil
}

func (m *QueuedMessage) GetMsgBeginRedelegate() *types.MsgBeginRedelegate {
	if x, ok := m.GetMsg().(*QueuedMessage_MsgBeginRedelegate); ok {
		return x.MsgBeginRedelegate
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*QueuedMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*QueuedMessage_MsgCreateValidator)(nil),
		(*QueuedMessage_MsgDelegate)(nil),
		(*QueuedMessage_MsgUndelegate)(nil),
		(*QueuedMessage_MsgBeginRedelegate)(nil),
	}
}

func init() {
	proto.RegisterType((*Epoch)(nil), "babylon.epoching.v1.Epoch")
	proto.RegisterType((*QueuedMessage)(nil), "babylon.epoching.v1.QueuedMessage")
}

func init() {
	proto.RegisterFile("babylon/epoching/v1/epoching.proto", fileDescriptor_2f2f209d5311f84c)
}

var fileDescriptor_2f2f209d5311f84c = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x93, 0x6d, 0xd3, 0xc3, 0xb4, 0x2b, 0x32, 0x5b, 0xa5, 0xec, 0x21, 0x6a, 0x45, 0x10,
	0x91, 0xc4, 0xaa, 0x9f, 0xa0, 0x2a, 0xa4, 0xc2, 0x2e, 0x18, 0xd0, 0x83, 0x07, 0xc3, 0x24, 0x79,
	0x4e, 0x86, 0x4d, 0x32, 0x65, 0x66, 0x52, 0xba, 0x5f, 0x42, 0xfc, 0x00, 0x7e, 0x20, 0x8f, 0x7b,
	0xf4, 0x28, 0xed, 0x17, 0x91, 0x79, 0x49, 0xb3, 0x8b, 0xc2, 0x7a, 0x9b, 0xf9, 0xbd, 0xff, 0xfb,
	0x25, 0xf3, 0x78, 0x64, 0x9e, 0xb2, 0xf4, 0xb2, 0x94, 0x75, 0x08, 0x6b, 0x99, 0x15, 0xa2, 0xe6,
	0xe1, 0x66, 0xd1, 0x9f, 0x83, 0xb5, 0x92, 0x46, 0xd2, 0x93, 0x2e, 0x13, 0xf4, 0x7c, 0xb3, 0x38,
	0x9d, 0x72, 0xc9, 0x25, 0xd6, 0x43, 0x7b, 0x6a, 0xa3, 0xa7, 0x0f, 0x32, 0xa9, 0x2b, 0xa9, 0x43,
	0x6d, 0xd8, 0x45, 0x2b, 0x4b, 0xc1, 0xb0, 0x45, 0x68, 0xb6, 0x6d, 0x60, 0xfe, 0xcd, 0x25, 0xde,
	0x3b, 0xab, 0xa1, 0x8f, 0xc8, 0x04, 0x7d, 0x49, 0xdd, 0x54, 0x29, 0xa8, 0x99, 0xfb, 0xd0, 0x7d,
	0x3a, 0x8c, 0xc7, 0xc8, 0xce, 0x11, 0xd1, 0xd7, 0xe4, 0x7e, 0xd6, 0x28, 0x05, 0xb5, 0x49, 0xda,
	0xa8, 0xa8, 0x0d, 0xa8, 0x0d, 0x2b, 0x67, 0x47, 0x18, 0x9e, 0x76, 0x55, 0x14, 0xae, 0xba, 0x1a,
	0x7d, 0x4e, 0xe8, 0x57, 0xa1, 0xb4, 0x49, 0xd2, 0x52, 0x66, 0x17, 0x49, 0x01, 0x82, 0x17, 0x66,
	0x36, 0xc0, 0x8e, 0xbb, 0x58, 0x59, 0xda, 0x42, 0x84, 0x7c, 0xfe, 0x63, 0x40, 0x8e, 0x3f, 0x34,
	0xd0, 0x40, 0x7e, 0x06, 0x5a, 0x33, 0x0e, 0xf4, 0x84, 0x78, 0x66, 0x9b, 0x88, 0x1c, 0xff, 0x68,
	0x12, 0x0f, 0xcd, 0x76, 0x95, 0xd3, 0x7b, 0x64, 0x54, 0x69, 0x6e, 0xe9, 0x11, 0x52, 0xaf, 0xd2,
	0x7c, 0x95, 0xd3, 0x2f, 0x64, 0x6a, 0x71, 0xa6, 0x80, 0x19, 0x48, 0x36, 0xac, 0x14, 0x39, 0x33,
	0x52, 0xe1, 0xd7, 0xc6, 0x2f, 0x9f, 0x05, 0xed, 0x38, 0x82, 0x6e, 0x1c, 0x41, 0x37, 0x8e, 0xe0,
	0x4c, 0xf3, 0x37, 0xd8, 0xf2, 0xe9, 0xd0, 0x11, 0x39, 0x31, 0xad, 0xfe, 0xa1, 0x34, 0x22, 0x13,
	0xeb, 0xcf, 0xa1, 0x04, 0xce, 0x0c, 0xcc, 0x86, 0xe8, 0x7d, 0x7c, 0x8b, 0xf7, 0x6d, 0x17, 0x8d,
	0x9c, 0x78, 0x5c, 0x5d, 0x5f, 0xe9, 0x39, 0xb9, 0x63, 0x4d, 0x4d, 0xdd, 0xbb, 0x3c, 0x74, 0x3d,
	0xb9, 0xc5, 0xf5, 0xb1, 0x0f, 0x47, 0x4e, 0x7c, 0x5c, 0xdd, 0x04, 0x87, 0x97, 0xa7, 0xc0, 0x45,
	0x9d, 0x28, 0xe8, 0xad, 0xa3, 0xff, 0xbe, 0x7c, 0x69, 0x5b, 0x62, 0xb8, 0xa1, 0xb6, 0x2f, 0xff,
	0x8b, 0x2e, 0x3d, 0x32, 0xb0, 0xf4, 0xfd, 0xcf, 0x9d, 0xef, 0x5e, 0xed, 0x7c, 0xf7, 0xf7, 0xce,
	0x77, 0xbf, 0xef, 0x7d, 0xe7, 0x6a, 0xef, 0x3b, 0xbf, 0xf6, 0xbe, 0xf3, 0xf9, 0x05, 0x17, 0xa6,
	0x68, 0xd2, 0x20, 0x93, 0x55, 0xd8, 0x2d, 0x68, 0x56, 0x30, 0x51, 0x1f, 0x2e, 0xe1, 0xf6, 0x7a,
	0xa7, 0xcd, 0xe5, 0x1a, 0x74, 0x3a, 0xc2, 0x15, 0x7c, 0xf5, 0x27, 0x00, 0x00, 0xff, 0xff, 0xbf,
	0x9c, 0x0b, 0xc3, 0xf4, 0x02, 0x00, 0x00,
}

func (m *Epoch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Epoch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Epoch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FirstBlockHeight != 0 {
		i = encodeVarintEpoching(dAtA, i, uint64(m.FirstBlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.CurrentEpochInterval != 0 {
		i = encodeVarintEpoching(dAtA, i, uint64(m.CurrentEpochInterval))
		i--
		dAtA[i] = 0x10
	}
	if m.EpochNumber != 0 {
		i = encodeVarintEpoching(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueuedMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueuedMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueuedMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Msg != nil {
		{
			size := m.Msg.Size()
			i -= size
			if _, err := m.Msg.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.MsgId) > 0 {
		i -= len(m.MsgId)
		copy(dAtA[i:], m.MsgId)
		i = encodeVarintEpoching(dAtA, i, uint64(len(m.MsgId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxId) > 0 {
		i -= len(m.TxId)
		copy(dAtA[i:], m.TxId)
		i = encodeVarintEpoching(dAtA, i, uint64(len(m.TxId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueuedMessage_MsgCreateValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueuedMessage_MsgCreateValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MsgCreateValidator != nil {
		{
			size, err := m.MsgCreateValidator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEpoching(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *QueuedMessage_MsgDelegate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueuedMessage_MsgDelegate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MsgDelegate != nil {
		{
			size, err := m.MsgDelegate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEpoching(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *QueuedMessage_MsgUndelegate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueuedMessage_MsgUndelegate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MsgUndelegate != nil {
		{
			size, err := m.MsgUndelegate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEpoching(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *QueuedMessage_MsgBeginRedelegate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueuedMessage_MsgBeginRedelegate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MsgBeginRedelegate != nil {
		{
			size, err := m.MsgBeginRedelegate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEpoching(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func encodeVarintEpoching(dAtA []byte, offset int, v uint64) int {
	offset -= sovEpoching(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Epoch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochNumber != 0 {
		n += 1 + sovEpoching(uint64(m.EpochNumber))
	}
	if m.CurrentEpochInterval != 0 {
		n += 1 + sovEpoching(uint64(m.CurrentEpochInterval))
	}
	if m.FirstBlockHeight != 0 {
		n += 1 + sovEpoching(uint64(m.FirstBlockHeight))
	}
	return n
}

func (m *QueuedMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxId)
	if l > 0 {
		n += 1 + l + sovEpoching(uint64(l))
	}
	l = len(m.MsgId)
	if l > 0 {
		n += 1 + l + sovEpoching(uint64(l))
	}
	if m.Msg != nil {
		n += m.Msg.Size()
	}
	return n
}

func (m *QueuedMessage_MsgCreateValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MsgCreateValidator != nil {
		l = m.MsgCreateValidator.Size()
		n += 1 + l + sovEpoching(uint64(l))
	}
	return n
}
func (m *QueuedMessage_MsgDelegate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MsgDelegate != nil {
		l = m.MsgDelegate.Size()
		n += 1 + l + sovEpoching(uint64(l))
	}
	return n
}
func (m *QueuedMessage_MsgUndelegate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MsgUndelegate != nil {
		l = m.MsgUndelegate.Size()
		n += 1 + l + sovEpoching(uint64(l))
	}
	return n
}
func (m *QueuedMessage_MsgBeginRedelegate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MsgBeginRedelegate != nil {
		l = m.MsgBeginRedelegate.Size()
		n += 1 + l + sovEpoching(uint64(l))
	}
	return n
}

func sovEpoching(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEpoching(x uint64) (n int) {
	return sovEpoching(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Epoch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpoching
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
			return fmt.Errorf("proto: Epoch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Epoch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoching
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochInterval", wireType)
			}
			m.CurrentEpochInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoching
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpochInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstBlockHeight", wireType)
			}
			m.FirstBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoching
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FirstBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpoching(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpoching
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
func (m *QueuedMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpoching
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
			return fmt.Errorf("proto: QueuedMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueuedMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoching
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
				return ErrInvalidLengthEpoching
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEpoching
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxId = append(m.TxId[:0], dAtA[iNdEx:postIndex]...)
			if m.TxId == nil {
				m.TxId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoching
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
				return ErrInvalidLengthEpoching
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEpoching
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgId = append(m.MsgId[:0], dAtA[iNdEx:postIndex]...)
			if m.MsgId == nil {
				m.MsgId = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgCreateValidator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoching
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
				return ErrInvalidLengthEpoching
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoching
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.MsgCreateValidator{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Msg = &QueuedMessage_MsgCreateValidator{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgDelegate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoching
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
				return ErrInvalidLengthEpoching
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoching
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.MsgDelegate{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Msg = &QueuedMessage_MsgDelegate{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgUndelegate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoching
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
				return ErrInvalidLengthEpoching
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoching
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.MsgUndelegate{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Msg = &QueuedMessage_MsgUndelegate{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgBeginRedelegate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoching
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
				return ErrInvalidLengthEpoching
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoching
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.MsgBeginRedelegate{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Msg = &QueuedMessage_MsgBeginRedelegate{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEpoching(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpoching
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
func skipEpoching(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEpoching
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
					return 0, ErrIntOverflowEpoching
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
					return 0, ErrIntOverflowEpoching
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
				return 0, ErrInvalidLengthEpoching
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEpoching
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEpoching
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEpoching        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEpoching          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEpoching = fmt.Errorf("proto: unexpected end of group")
)
