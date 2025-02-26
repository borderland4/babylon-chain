// Code generated by MockGen. DO NOT EDIT.
// Source: x/finality/types/expected_keepers.go

// Package types is a generated GoMock package.
package types

import (
	context "context"
	reflect "reflect"

	types "github.com/babylonchain/babylon/types"
	types0 "github.com/babylonchain/babylon/x/btcstaking/types"
	gomock "github.com/golang/mock/gomock"
)

// MockBTCStakingKeeper is a mock of BTCStakingKeeper interface.
type MockBTCStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBTCStakingKeeperMockRecorder
}

// MockBTCStakingKeeperMockRecorder is the mock recorder for MockBTCStakingKeeper.
type MockBTCStakingKeeperMockRecorder struct {
	mock *MockBTCStakingKeeper
}

// NewMockBTCStakingKeeper creates a new mock instance.
func NewMockBTCStakingKeeper(ctrl *gomock.Controller) *MockBTCStakingKeeper {
	mock := &MockBTCStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockBTCStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBTCStakingKeeper) EXPECT() *MockBTCStakingKeeperMockRecorder {
	return m.recorder
}

// GetBTCStakingActivatedHeight mocks base method.
func (m *MockBTCStakingKeeper) GetBTCStakingActivatedHeight(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBTCStakingActivatedHeight", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBTCStakingActivatedHeight indicates an expected call of GetBTCStakingActivatedHeight.
func (mr *MockBTCStakingKeeperMockRecorder) GetBTCStakingActivatedHeight(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBTCStakingActivatedHeight", reflect.TypeOf((*MockBTCStakingKeeper)(nil).GetBTCStakingActivatedHeight), ctx)
}

// GetFinalityProvider mocks base method.
func (m *MockBTCStakingKeeper) GetFinalityProvider(ctx context.Context, fpBTCPK []byte) (*types0.FinalityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalityProvider", ctx, fpBTCPK)
	ret0, _ := ret[0].(*types0.FinalityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFinalityProvider indicates an expected call of GetFinalityProvider.
func (mr *MockBTCStakingKeeperMockRecorder) GetFinalityProvider(ctx, fpBTCPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalityProvider", reflect.TypeOf((*MockBTCStakingKeeper)(nil).GetFinalityProvider), ctx, fpBTCPK)
}

// GetLastFinalizedEpoch mocks base method.
func (m *MockBTCStakingKeeper) GetLastFinalizedEpoch(ctx context.Context) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastFinalizedEpoch", ctx)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetLastFinalizedEpoch indicates an expected call of GetLastFinalizedEpoch.
func (mr *MockBTCStakingKeeperMockRecorder) GetLastFinalizedEpoch(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastFinalizedEpoch", reflect.TypeOf((*MockBTCStakingKeeper)(nil).GetLastFinalizedEpoch), ctx)
}

// GetParams mocks base method.
func (m *MockBTCStakingKeeper) GetParams(ctx context.Context) types0.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types0.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockBTCStakingKeeperMockRecorder) GetParams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockBTCStakingKeeper)(nil).GetParams), ctx)
}

// GetVotingPower mocks base method.
func (m *MockBTCStakingKeeper) GetVotingPower(ctx context.Context, fpBTCPK []byte, height uint64) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVotingPower", ctx, fpBTCPK, height)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetVotingPower indicates an expected call of GetVotingPower.
func (mr *MockBTCStakingKeeperMockRecorder) GetVotingPower(ctx, fpBTCPK, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVotingPower", reflect.TypeOf((*MockBTCStakingKeeper)(nil).GetVotingPower), ctx, fpBTCPK, height)
}

// GetVotingPowerDistCache mocks base method.
func (m *MockBTCStakingKeeper) GetVotingPowerDistCache(ctx context.Context, height uint64) (*types0.VotingPowerDistCache, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVotingPowerDistCache", ctx, height)
	ret0, _ := ret[0].(*types0.VotingPowerDistCache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVotingPowerDistCache indicates an expected call of GetVotingPowerDistCache.
func (mr *MockBTCStakingKeeperMockRecorder) GetVotingPowerDistCache(ctx, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVotingPowerDistCache", reflect.TypeOf((*MockBTCStakingKeeper)(nil).GetVotingPowerDistCache), ctx, height)
}

// GetVotingPowerTable mocks base method.
func (m *MockBTCStakingKeeper) GetVotingPowerTable(ctx context.Context, height uint64) map[string]uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVotingPowerTable", ctx, height)
	ret0, _ := ret[0].(map[string]uint64)
	return ret0
}

// GetVotingPowerTable indicates an expected call of GetVotingPowerTable.
func (mr *MockBTCStakingKeeperMockRecorder) GetVotingPowerTable(ctx, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVotingPowerTable", reflect.TypeOf((*MockBTCStakingKeeper)(nil).GetVotingPowerTable), ctx, height)
}

// HasFinalityProvider mocks base method.
func (m *MockBTCStakingKeeper) HasFinalityProvider(ctx context.Context, fpBTCPK []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasFinalityProvider", ctx, fpBTCPK)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasFinalityProvider indicates an expected call of HasFinalityProvider.
func (mr *MockBTCStakingKeeperMockRecorder) HasFinalityProvider(ctx, fpBTCPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasFinalityProvider", reflect.TypeOf((*MockBTCStakingKeeper)(nil).HasFinalityProvider), ctx, fpBTCPK)
}

// RemoveVotingPowerDistCache mocks base method.
func (m *MockBTCStakingKeeper) RemoveVotingPowerDistCache(ctx context.Context, height uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveVotingPowerDistCache", ctx, height)
}

// RemoveVotingPowerDistCache indicates an expected call of RemoveVotingPowerDistCache.
func (mr *MockBTCStakingKeeperMockRecorder) RemoveVotingPowerDistCache(ctx, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVotingPowerDistCache", reflect.TypeOf((*MockBTCStakingKeeper)(nil).RemoveVotingPowerDistCache), ctx, height)
}

// RevertSluggishFinalityProvider mocks base method.
func (m *MockBTCStakingKeeper) RevertSluggishFinalityProvider(ctx context.Context, fpBTCPK []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevertSluggishFinalityProvider", ctx, fpBTCPK)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevertSluggishFinalityProvider indicates an expected call of RevertSluggishFinalityProvider.
func (mr *MockBTCStakingKeeperMockRecorder) RevertSluggishFinalityProvider(ctx, fpBTCPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevertSluggishFinalityProvider", reflect.TypeOf((*MockBTCStakingKeeper)(nil).RevertSluggishFinalityProvider), ctx, fpBTCPK)
}

// SlashFinalityProvider mocks base method.
func (m *MockBTCStakingKeeper) SlashFinalityProvider(ctx context.Context, fpBTCPK []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlashFinalityProvider", ctx, fpBTCPK)
	ret0, _ := ret[0].(error)
	return ret0
}

// SlashFinalityProvider indicates an expected call of SlashFinalityProvider.
func (mr *MockBTCStakingKeeperMockRecorder) SlashFinalityProvider(ctx, fpBTCPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlashFinalityProvider", reflect.TypeOf((*MockBTCStakingKeeper)(nil).SlashFinalityProvider), ctx, fpBTCPK)
}

// MockIncentiveKeeper is a mock of IncentiveKeeper interface.
type MockIncentiveKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockIncentiveKeeperMockRecorder
}

// MockIncentiveKeeperMockRecorder is the mock recorder for MockIncentiveKeeper.
type MockIncentiveKeeperMockRecorder struct {
	mock *MockIncentiveKeeper
}

// NewMockIncentiveKeeper creates a new mock instance.
func NewMockIncentiveKeeper(ctrl *gomock.Controller) *MockIncentiveKeeper {
	mock := &MockIncentiveKeeper{ctrl: ctrl}
	mock.recorder = &MockIncentiveKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIncentiveKeeper) EXPECT() *MockIncentiveKeeperMockRecorder {
	return m.recorder
}

// RewardBTCStaking mocks base method.
func (m *MockIncentiveKeeper) RewardBTCStaking(ctx context.Context, height uint64, filteredDc *types0.VotingPowerDistCache) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RewardBTCStaking", ctx, height, filteredDc)
}

// RewardBTCStaking indicates an expected call of RewardBTCStaking.
func (mr *MockIncentiveKeeperMockRecorder) RewardBTCStaking(ctx, height, filteredDc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RewardBTCStaking", reflect.TypeOf((*MockIncentiveKeeper)(nil).RewardBTCStaking), ctx, height, filteredDc)
}

// MockBtcStakingHooks is a mock of BtcStakingHooks interface.
type MockBtcStakingHooks struct {
	ctrl     *gomock.Controller
	recorder *MockBtcStakingHooksMockRecorder
}

// MockBtcStakingHooksMockRecorder is the mock recorder for MockBtcStakingHooks.
type MockBtcStakingHooksMockRecorder struct {
	mock *MockBtcStakingHooks
}

// NewMockBtcStakingHooks creates a new mock instance.
func NewMockBtcStakingHooks(ctrl *gomock.Controller) *MockBtcStakingHooks {
	mock := &MockBtcStakingHooks{ctrl: ctrl}
	mock.recorder = &MockBtcStakingHooksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBtcStakingHooks) EXPECT() *MockBtcStakingHooksMockRecorder {
	return m.recorder
}

// AfterFinalityProviderActivated mocks base method.
func (m *MockBtcStakingHooks) AfterFinalityProviderActivated(ctx context.Context, btcPk *types.BIP340PubKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterFinalityProviderActivated", ctx, btcPk)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterFinalityProviderActivated indicates an expected call of AfterFinalityProviderActivated.
func (mr *MockBtcStakingHooksMockRecorder) AfterFinalityProviderActivated(ctx, btcPk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterFinalityProviderActivated", reflect.TypeOf((*MockBtcStakingHooks)(nil).AfterFinalityProviderActivated), ctx, btcPk)
}

// MockFinalityHooks is a mock of FinalityHooks interface.
type MockFinalityHooks struct {
	ctrl     *gomock.Controller
	recorder *MockFinalityHooksMockRecorder
}

// MockFinalityHooksMockRecorder is the mock recorder for MockFinalityHooks.
type MockFinalityHooksMockRecorder struct {
	mock *MockFinalityHooks
}

// NewMockFinalityHooks creates a new mock instance.
func NewMockFinalityHooks(ctrl *gomock.Controller) *MockFinalityHooks {
	mock := &MockFinalityHooks{ctrl: ctrl}
	mock.recorder = &MockFinalityHooksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFinalityHooks) EXPECT() *MockFinalityHooksMockRecorder {
	return m.recorder
}

// AfterSluggishFinalityProviderDetected mocks base method.
func (m *MockFinalityHooks) AfterSluggishFinalityProviderDetected(ctx context.Context, btcPk *types.BIP340PubKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterSluggishFinalityProviderDetected", ctx, btcPk)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterSluggishFinalityProviderDetected indicates an expected call of AfterSluggishFinalityProviderDetected.
func (mr *MockFinalityHooksMockRecorder) AfterSluggishFinalityProviderDetected(ctx, btcPk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterSluggishFinalityProviderDetected", reflect.TypeOf((*MockFinalityHooks)(nil).AfterSluggishFinalityProviderDetected), ctx, btcPk)
}
