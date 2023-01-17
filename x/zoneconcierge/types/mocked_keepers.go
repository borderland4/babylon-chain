// Code generated by MockGen. DO NOT EDIT.
// Source: x/zoneconcierge/types/expected_keepers.go

// Package types is a generated GoMock package.
package types

import (
	context "context"
	reflect "reflect"

	types "github.com/babylonchain/babylon/x/btccheckpoint/types"
	types0 "github.com/babylonchain/babylon/x/checkpointing/types"
	types1 "github.com/babylonchain/babylon/x/epoching/types"
	types2 "github.com/cosmos/cosmos-sdk/types"
	types3 "github.com/cosmos/cosmos-sdk/x/auth/types"
	types4 "github.com/cosmos/cosmos-sdk/x/capability/types"
	types5 "github.com/cosmos/ibc-go/v5/modules/core/03-connection/types"
	types6 "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	exported "github.com/cosmos/ibc-go/v5/modules/core/exported"
	gomock "github.com/golang/mock/gomock"
	crypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetModuleAccount mocks base method.
func (m *MockAccountKeeper) GetModuleAccount(ctx types2.Context, name string) types3.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", ctx, name)
	ret0, _ := ret[0].(types3.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccount(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccount), ctx, name)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(name string) types2.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", name)
	ret0, _ := ret[0].(types2.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), name)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// BlockedAddr mocks base method.
func (m *MockBankKeeper) BlockedAddr(addr types2.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockedAddr", addr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// BlockedAddr indicates an expected call of BlockedAddr.
func (mr *MockBankKeeperMockRecorder) BlockedAddr(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockedAddr", reflect.TypeOf((*MockBankKeeper)(nil).BlockedAddr), addr)
}

// BurnCoins mocks base method.
func (m *MockBankKeeper) BurnCoins(ctx types2.Context, moduleName string, amt types2.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BurnCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// BurnCoins indicates an expected call of BurnCoins.
func (mr *MockBankKeeperMockRecorder) BurnCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BurnCoins", reflect.TypeOf((*MockBankKeeper)(nil).BurnCoins), ctx, moduleName, amt)
}

// MintCoins mocks base method.
func (m *MockBankKeeper) MintCoins(ctx types2.Context, moduleName string, amt types2.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// MintCoins indicates an expected call of MintCoins.
func (mr *MockBankKeeperMockRecorder) MintCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCoins", reflect.TypeOf((*MockBankKeeper)(nil).MintCoins), ctx, moduleName, amt)
}

// SendCoins mocks base method.
func (m *MockBankKeeper) SendCoins(ctx types2.Context, fromAddr, toAddr types2.AccAddress, amt types2.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoins", ctx, fromAddr, toAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoins indicates an expected call of SendCoins.
func (mr *MockBankKeeperMockRecorder) SendCoins(ctx, fromAddr, toAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoins", reflect.TypeOf((*MockBankKeeper)(nil).SendCoins), ctx, fromAddr, toAddr, amt)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx types2.Context, senderAddr types2.AccAddress, recipientModule string, amt types2.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx types2.Context, senderModule string, recipientAddr types2.AccAddress, amt types2.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// MockICS4Wrapper is a mock of ICS4Wrapper interface.
type MockICS4Wrapper struct {
	ctrl     *gomock.Controller
	recorder *MockICS4WrapperMockRecorder
}

// MockICS4WrapperMockRecorder is the mock recorder for MockICS4Wrapper.
type MockICS4WrapperMockRecorder struct {
	mock *MockICS4Wrapper
}

// NewMockICS4Wrapper creates a new mock instance.
func NewMockICS4Wrapper(ctrl *gomock.Controller) *MockICS4Wrapper {
	mock := &MockICS4Wrapper{ctrl: ctrl}
	mock.recorder = &MockICS4WrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICS4Wrapper) EXPECT() *MockICS4WrapperMockRecorder {
	return m.recorder
}

// SendPacket mocks base method.
func (m *MockICS4Wrapper) SendPacket(ctx types2.Context, channelCap *types4.Capability, packet exported.PacketI) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendPacket", ctx, channelCap, packet)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendPacket indicates an expected call of SendPacket.
func (mr *MockICS4WrapperMockRecorder) SendPacket(ctx, channelCap, packet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPacket", reflect.TypeOf((*MockICS4Wrapper)(nil).SendPacket), ctx, channelCap, packet)
}

// MockChannelKeeper is a mock of ChannelKeeper interface.
type MockChannelKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockChannelKeeperMockRecorder
}

// MockChannelKeeperMockRecorder is the mock recorder for MockChannelKeeper.
type MockChannelKeeperMockRecorder struct {
	mock *MockChannelKeeper
}

// NewMockChannelKeeper creates a new mock instance.
func NewMockChannelKeeper(ctrl *gomock.Controller) *MockChannelKeeper {
	mock := &MockChannelKeeper{ctrl: ctrl}
	mock.recorder = &MockChannelKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelKeeper) EXPECT() *MockChannelKeeperMockRecorder {
	return m.recorder
}

// GetAllChannels mocks base method.
func (m *MockChannelKeeper) GetAllChannels(ctx types2.Context) []types6.IdentifiedChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllChannels", ctx)
	ret0, _ := ret[0].([]types6.IdentifiedChannel)
	return ret0
}

// GetAllChannels indicates an expected call of GetAllChannels.
func (mr *MockChannelKeeperMockRecorder) GetAllChannels(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllChannels", reflect.TypeOf((*MockChannelKeeper)(nil).GetAllChannels), ctx)
}

// GetChannel mocks base method.
func (m *MockChannelKeeper) GetChannel(ctx types2.Context, srcPort, srcChan string) (types6.Channel, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannel", ctx, srcPort, srcChan)
	ret0, _ := ret[0].(types6.Channel)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetChannel indicates an expected call of GetChannel.
func (mr *MockChannelKeeperMockRecorder) GetChannel(ctx, srcPort, srcChan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannel", reflect.TypeOf((*MockChannelKeeper)(nil).GetChannel), ctx, srcPort, srcChan)
}

// GetNextSequenceSend mocks base method.
func (m *MockChannelKeeper) GetNextSequenceSend(ctx types2.Context, portID, channelID string) (uint64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextSequenceSend", ctx, portID, channelID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetNextSequenceSend indicates an expected call of GetNextSequenceSend.
func (mr *MockChannelKeeperMockRecorder) GetNextSequenceSend(ctx, portID, channelID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextSequenceSend", reflect.TypeOf((*MockChannelKeeper)(nil).GetNextSequenceSend), ctx, portID, channelID)
}

// MockClientKeeper is a mock of ClientKeeper interface.
type MockClientKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockClientKeeperMockRecorder
}

// MockClientKeeperMockRecorder is the mock recorder for MockClientKeeper.
type MockClientKeeperMockRecorder struct {
	mock *MockClientKeeper
}

// NewMockClientKeeper creates a new mock instance.
func NewMockClientKeeper(ctrl *gomock.Controller) *MockClientKeeper {
	mock := &MockClientKeeper{ctrl: ctrl}
	mock.recorder = &MockClientKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientKeeper) EXPECT() *MockClientKeeperMockRecorder {
	return m.recorder
}

// GetClientConsensusState mocks base method.
func (m *MockClientKeeper) GetClientConsensusState(ctx types2.Context, clientID string) (exported.ConsensusState, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientConsensusState", ctx, clientID)
	ret0, _ := ret[0].(exported.ConsensusState)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetClientConsensusState indicates an expected call of GetClientConsensusState.
func (mr *MockClientKeeperMockRecorder) GetClientConsensusState(ctx, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientConsensusState", reflect.TypeOf((*MockClientKeeper)(nil).GetClientConsensusState), ctx, clientID)
}

// MockConnectionKeeper is a mock of ConnectionKeeper interface.
type MockConnectionKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionKeeperMockRecorder
}

// MockConnectionKeeperMockRecorder is the mock recorder for MockConnectionKeeper.
type MockConnectionKeeperMockRecorder struct {
	mock *MockConnectionKeeper
}

// NewMockConnectionKeeper creates a new mock instance.
func NewMockConnectionKeeper(ctrl *gomock.Controller) *MockConnectionKeeper {
	mock := &MockConnectionKeeper{ctrl: ctrl}
	mock.recorder = &MockConnectionKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionKeeper) EXPECT() *MockConnectionKeeperMockRecorder {
	return m.recorder
}

// GetConnection mocks base method.
func (m *MockConnectionKeeper) GetConnection(ctx types2.Context, connectionID string) (types5.ConnectionEnd, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnection", ctx, connectionID)
	ret0, _ := ret[0].(types5.ConnectionEnd)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetConnection indicates an expected call of GetConnection.
func (mr *MockConnectionKeeperMockRecorder) GetConnection(ctx, connectionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnection", reflect.TypeOf((*MockConnectionKeeper)(nil).GetConnection), ctx, connectionID)
}

// MockPortKeeper is a mock of PortKeeper interface.
type MockPortKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockPortKeeperMockRecorder
}

// MockPortKeeperMockRecorder is the mock recorder for MockPortKeeper.
type MockPortKeeperMockRecorder struct {
	mock *MockPortKeeper
}

// NewMockPortKeeper creates a new mock instance.
func NewMockPortKeeper(ctrl *gomock.Controller) *MockPortKeeper {
	mock := &MockPortKeeper{ctrl: ctrl}
	mock.recorder = &MockPortKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPortKeeper) EXPECT() *MockPortKeeperMockRecorder {
	return m.recorder
}

// BindPort mocks base method.
func (m *MockPortKeeper) BindPort(ctx types2.Context, portID string) *types4.Capability {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindPort", ctx, portID)
	ret0, _ := ret[0].(*types4.Capability)
	return ret0
}

// BindPort indicates an expected call of BindPort.
func (mr *MockPortKeeperMockRecorder) BindPort(ctx, portID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindPort", reflect.TypeOf((*MockPortKeeper)(nil).BindPort), ctx, portID)
}

// MockScopedKeeper is a mock of ScopedKeeper interface.
type MockScopedKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockScopedKeeperMockRecorder
}

// MockScopedKeeperMockRecorder is the mock recorder for MockScopedKeeper.
type MockScopedKeeperMockRecorder struct {
	mock *MockScopedKeeper
}

// NewMockScopedKeeper creates a new mock instance.
func NewMockScopedKeeper(ctrl *gomock.Controller) *MockScopedKeeper {
	mock := &MockScopedKeeper{ctrl: ctrl}
	mock.recorder = &MockScopedKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScopedKeeper) EXPECT() *MockScopedKeeperMockRecorder {
	return m.recorder
}

// AuthenticateCapability mocks base method.
func (m *MockScopedKeeper) AuthenticateCapability(ctx types2.Context, cap *types4.Capability, name string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateCapability", ctx, cap, name)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthenticateCapability indicates an expected call of AuthenticateCapability.
func (mr *MockScopedKeeperMockRecorder) AuthenticateCapability(ctx, cap, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateCapability", reflect.TypeOf((*MockScopedKeeper)(nil).AuthenticateCapability), ctx, cap, name)
}

// ClaimCapability mocks base method.
func (m *MockScopedKeeper) ClaimCapability(ctx types2.Context, cap *types4.Capability, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClaimCapability", ctx, cap, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClaimCapability indicates an expected call of ClaimCapability.
func (mr *MockScopedKeeperMockRecorder) ClaimCapability(ctx, cap, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClaimCapability", reflect.TypeOf((*MockScopedKeeper)(nil).ClaimCapability), ctx, cap, name)
}

// GetCapability mocks base method.
func (m *MockScopedKeeper) GetCapability(ctx types2.Context, name string) (*types4.Capability, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCapability", ctx, name)
	ret0, _ := ret[0].(*types4.Capability)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetCapability indicates an expected call of GetCapability.
func (mr *MockScopedKeeperMockRecorder) GetCapability(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCapability", reflect.TypeOf((*MockScopedKeeper)(nil).GetCapability), ctx, name)
}

// LookupModules mocks base method.
func (m *MockScopedKeeper) LookupModules(ctx types2.Context, name string) ([]string, *types4.Capability, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupModules", ctx, name)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(*types4.Capability)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LookupModules indicates an expected call of LookupModules.
func (mr *MockScopedKeeperMockRecorder) LookupModules(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupModules", reflect.TypeOf((*MockScopedKeeper)(nil).LookupModules), ctx, name)
}

// MockBtcCheckpointKeeper is a mock of BtcCheckpointKeeper interface.
type MockBtcCheckpointKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBtcCheckpointKeeperMockRecorder
}

// MockBtcCheckpointKeeperMockRecorder is the mock recorder for MockBtcCheckpointKeeper.
type MockBtcCheckpointKeeperMockRecorder struct {
	mock *MockBtcCheckpointKeeper
}

// NewMockBtcCheckpointKeeper creates a new mock instance.
func NewMockBtcCheckpointKeeper(ctrl *gomock.Controller) *MockBtcCheckpointKeeper {
	mock := &MockBtcCheckpointKeeper{ctrl: ctrl}
	mock.recorder = &MockBtcCheckpointKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBtcCheckpointKeeper) EXPECT() *MockBtcCheckpointKeeperMockRecorder {
	return m.recorder
}

// GetFinalizedEpochDataWithBestSubmission mocks base method.
func (m *MockBtcCheckpointKeeper) GetFinalizedEpochDataWithBestSubmission(ctx types2.Context, e uint64) (types.BtcStatus, *types.SubmissionKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalizedEpochDataWithBestSubmission", ctx, e)
	ret0, _ := ret[0].(types.BtcStatus)
	ret1, _ := ret[1].(*types.SubmissionKey)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetFinalizedEpochDataWithBestSubmission indicates an expected call of GetFinalizedEpochDataWithBestSubmission.
func (mr *MockBtcCheckpointKeeperMockRecorder) GetFinalizedEpochDataWithBestSubmission(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalizedEpochDataWithBestSubmission", reflect.TypeOf((*MockBtcCheckpointKeeper)(nil).GetFinalizedEpochDataWithBestSubmission), ctx, e)
}

// GetSubmissionData mocks base method.
func (m *MockBtcCheckpointKeeper) GetSubmissionData(ctx types2.Context, sk types.SubmissionKey) *types.SubmissionData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubmissionData", ctx, sk)
	ret0, _ := ret[0].(*types.SubmissionData)
	return ret0
}

// GetSubmissionData indicates an expected call of GetSubmissionData.
func (mr *MockBtcCheckpointKeeperMockRecorder) GetSubmissionData(ctx, sk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubmissionData", reflect.TypeOf((*MockBtcCheckpointKeeper)(nil).GetSubmissionData), ctx, sk)
}

// MockCheckpointingKeeper is a mock of CheckpointingKeeper interface.
type MockCheckpointingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockCheckpointingKeeperMockRecorder
}

// MockCheckpointingKeeperMockRecorder is the mock recorder for MockCheckpointingKeeper.
type MockCheckpointingKeeperMockRecorder struct {
	mock *MockCheckpointingKeeper
}

// NewMockCheckpointingKeeper creates a new mock instance.
func NewMockCheckpointingKeeper(ctrl *gomock.Controller) *MockCheckpointingKeeper {
	mock := &MockCheckpointingKeeper{ctrl: ctrl}
	mock.recorder = &MockCheckpointingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckpointingKeeper) EXPECT() *MockCheckpointingKeeperMockRecorder {
	return m.recorder
}

// GetBLSPubKeySet mocks base method.
func (m *MockCheckpointingKeeper) GetBLSPubKeySet(ctx types2.Context, epochNumber uint64) ([]*types0.ValidatorWithBlsKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBLSPubKeySet", ctx, epochNumber)
	ret0, _ := ret[0].([]*types0.ValidatorWithBlsKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBLSPubKeySet indicates an expected call of GetBLSPubKeySet.
func (mr *MockCheckpointingKeeperMockRecorder) GetBLSPubKeySet(ctx, epochNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBLSPubKeySet", reflect.TypeOf((*MockCheckpointingKeeper)(nil).GetBLSPubKeySet), ctx, epochNumber)
}

// GetRawCheckpoint mocks base method.
func (m *MockCheckpointingKeeper) GetRawCheckpoint(ctx types2.Context, epochNumber uint64) (*types0.RawCheckpointWithMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawCheckpoint", ctx, epochNumber)
	ret0, _ := ret[0].(*types0.RawCheckpointWithMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawCheckpoint indicates an expected call of GetRawCheckpoint.
func (mr *MockCheckpointingKeeperMockRecorder) GetRawCheckpoint(ctx, epochNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawCheckpoint", reflect.TypeOf((*MockCheckpointingKeeper)(nil).GetRawCheckpoint), ctx, epochNumber)
}

// MockEpochingKeeper is a mock of EpochingKeeper interface.
type MockEpochingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockEpochingKeeperMockRecorder
}

// MockEpochingKeeperMockRecorder is the mock recorder for MockEpochingKeeper.
type MockEpochingKeeperMockRecorder struct {
	mock *MockEpochingKeeper
}

// NewMockEpochingKeeper creates a new mock instance.
func NewMockEpochingKeeper(ctrl *gomock.Controller) *MockEpochingKeeper {
	mock := &MockEpochingKeeper{ctrl: ctrl}
	mock.recorder = &MockEpochingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEpochingKeeper) EXPECT() *MockEpochingKeeperMockRecorder {
	return m.recorder
}

// GetEpoch mocks base method.
func (m *MockEpochingKeeper) GetEpoch(ctx types2.Context) *types1.Epoch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEpoch", ctx)
	ret0, _ := ret[0].(*types1.Epoch)
	return ret0
}

// GetEpoch indicates an expected call of GetEpoch.
func (mr *MockEpochingKeeperMockRecorder) GetEpoch(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpoch", reflect.TypeOf((*MockEpochingKeeper)(nil).GetEpoch), ctx)
}

// GetHistoricalEpoch mocks base method.
func (m *MockEpochingKeeper) GetHistoricalEpoch(ctx types2.Context, epochNumber uint64) (*types1.Epoch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoricalEpoch", ctx, epochNumber)
	ret0, _ := ret[0].(*types1.Epoch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistoricalEpoch indicates an expected call of GetHistoricalEpoch.
func (mr *MockEpochingKeeperMockRecorder) GetHistoricalEpoch(ctx, epochNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoricalEpoch", reflect.TypeOf((*MockEpochingKeeper)(nil).GetHistoricalEpoch), ctx, epochNumber)
}

// ProveAppHashInEpoch mocks base method.
func (m *MockEpochingKeeper) ProveAppHashInEpoch(ctx types2.Context, height, epochNumber uint64) (*crypto.Proof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProveAppHashInEpoch", ctx, height, epochNumber)
	ret0, _ := ret[0].(*crypto.Proof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProveAppHashInEpoch indicates an expected call of ProveAppHashInEpoch.
func (mr *MockEpochingKeeperMockRecorder) ProveAppHashInEpoch(ctx, height, epochNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProveAppHashInEpoch", reflect.TypeOf((*MockEpochingKeeper)(nil).ProveAppHashInEpoch), ctx, height, epochNumber)
}

// MockTMClient is a mock of TMClient interface.
type MockTMClient struct {
	ctrl     *gomock.Controller
	recorder *MockTMClientMockRecorder
}

// MockTMClientMockRecorder is the mock recorder for MockTMClient.
type MockTMClientMockRecorder struct {
	mock *MockTMClient
}

// NewMockTMClient creates a new mock instance.
func NewMockTMClient(ctrl *gomock.Controller) *MockTMClient {
	mock := &MockTMClient{ctrl: ctrl}
	mock.recorder = &MockTMClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTMClient) EXPECT() *MockTMClientMockRecorder {
	return m.recorder
}

// Tx mocks base method.
func (m *MockTMClient) Tx(ctx context.Context, hash []byte, prove bool) (*coretypes.ResultTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx", ctx, hash, prove)
	ret0, _ := ret[0].(*coretypes.ResultTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tx indicates an expected call of Tx.
func (mr *MockTMClientMockRecorder) Tx(ctx, hash, prove interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockTMClient)(nil).Tx), ctx, hash, prove)
}
