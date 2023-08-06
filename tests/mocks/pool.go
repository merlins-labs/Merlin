// Code generated by MockGen. DO NOT EDIT.
// Source: x/poolmanager/types/pool.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
	types0 "github.com/merlins-labs/merlin/v16/x/poolmanager/types"
)

// MockPoolI is a mock of PoolI interface.
type MockPoolI struct {
	ctrl     *gomock.Controller
	recorder *MockPoolIMockRecorder
}

// MockPoolIMockRecorder is the mock recorder for MockPoolI.
type MockPoolIMockRecorder struct {
	mock *MockPoolI
}

// NewMockPoolI creates a new mock instance.
func NewMockPoolI(ctrl *gomock.Controller) *MockPoolI {
	mock := &MockPoolI{ctrl: ctrl}
	mock.recorder = &MockPoolIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPoolI) EXPECT() *MockPoolIMockRecorder {
	return m.recorder
}

// AsSerializablePool mocks base method.
func (m *MockPoolI) AsSerializablePool() types0.PoolI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsSerializablePool")
	ret0, _ := ret[0].(types0.PoolI)
	return ret0
}

// AsSerializablePool indicates an expected call of AsSerializablePool.
func (mr *MockPoolIMockRecorder) AsSerializablePool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsSerializablePool", reflect.TypeOf((*MockPoolI)(nil).AsSerializablePool))
}

// GetAddress mocks base method.
func (m *MockPoolI) GetAddress() types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddress")
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetAddress indicates an expected call of GetAddress.
func (mr *MockPoolIMockRecorder) GetAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddress", reflect.TypeOf((*MockPoolI)(nil).GetAddress))
}

// GetId mocks base method.
func (m *MockPoolI) GetId() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetId")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetId indicates an expected call of GetId.
func (mr *MockPoolIMockRecorder) GetId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetId", reflect.TypeOf((*MockPoolI)(nil).GetId))
}

// GetSpreadFactor mocks base method.
func (m *MockPoolI) GetSpreadFactor(ctx types.Context) types.Dec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpreadFactor", ctx)
	ret0, _ := ret[0].(types.Dec)
	return ret0
}

// GetSpreadFactor indicates an expected call of GetSpreadFactor.
func (mr *MockPoolIMockRecorder) GetSpreadFactor(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpreadFactor", reflect.TypeOf((*MockPoolI)(nil).GetSpreadFactor), ctx)
}

// GetType mocks base method.
func (m *MockPoolI) GetType() types0.PoolType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(types0.PoolType)
	return ret0
}

// GetType indicates an expected call of GetType.
func (mr *MockPoolIMockRecorder) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockPoolI)(nil).GetType))
}

// IsActive mocks base method.
func (m *MockPoolI) IsActive(ctx types.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsActive", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsActive indicates an expected call of IsActive.
func (mr *MockPoolIMockRecorder) IsActive(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActive", reflect.TypeOf((*MockPoolI)(nil).IsActive), ctx)
}

// ProtoMessage mocks base method.
func (m *MockPoolI) ProtoMessage() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProtoMessage")
}

// ProtoMessage indicates an expected call of ProtoMessage.
func (mr *MockPoolIMockRecorder) ProtoMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtoMessage", reflect.TypeOf((*MockPoolI)(nil).ProtoMessage))
}

// Reset mocks base method.
func (m *MockPoolI) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockPoolIMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockPoolI)(nil).Reset))
}

// SpotPrice mocks base method.
func (m *MockPoolI) SpotPrice(ctx types.Context, quoteAssetDenom, baseAssetDenom string) (types.Dec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpotPrice", ctx, quoteAssetDenom, baseAssetDenom)
	ret0, _ := ret[0].(types.Dec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpotPrice indicates an expected call of SpotPrice.
func (mr *MockPoolIMockRecorder) SpotPrice(ctx, quoteAssetDenom, baseAssetDenom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpotPrice", reflect.TypeOf((*MockPoolI)(nil).SpotPrice), ctx, quoteAssetDenom, baseAssetDenom)
}

// String mocks base method.
func (m *MockPoolI) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockPoolIMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockPoolI)(nil).String))
}
