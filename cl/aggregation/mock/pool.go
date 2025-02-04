// Code generated by MockGen. DO NOT EDIT.
// Source: ./cl/aggregation/pool.go
//
// Generated by this command:
//
//	mockgen -source=./cl/aggregation/pool.go -destination=./cl/aggregation/mock/pool.go
//

// Package mock_aggregation is a generated GoMock package.
package mock_aggregation

import (
	reflect "reflect"

	common "github.com/ledgerwatch/erigon-lib/common"
	solid "github.com/ledgerwatch/erigon/cl/cltypes/solid"
	gomock "go.uber.org/mock/gomock"
)

// MockAggregationPool is a mock of AggregationPool interface.
type MockAggregationPool struct {
	ctrl     *gomock.Controller
	recorder *MockAggregationPoolMockRecorder
}

// MockAggregationPoolMockRecorder is the mock recorder for MockAggregationPool.
type MockAggregationPoolMockRecorder struct {
	mock *MockAggregationPool
}

// NewMockAggregationPool creates a new mock instance.
func NewMockAggregationPool(ctrl *gomock.Controller) *MockAggregationPool {
	mock := &MockAggregationPool{ctrl: ctrl}
	mock.recorder = &MockAggregationPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAggregationPool) EXPECT() *MockAggregationPoolMockRecorder {
	return m.recorder
}

// AddAttestation mocks base method.
func (m *MockAggregationPool) AddAttestation(att *solid.Attestation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAttestation", att)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAttestation indicates an expected call of AddAttestation.
func (mr *MockAggregationPoolMockRecorder) AddAttestation(att any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAttestation", reflect.TypeOf((*MockAggregationPool)(nil).AddAttestation), att)
}

// GetAggregatationByRoot mocks base method.
func (m *MockAggregationPool) GetAggregatationByRoot(root common.Hash) *solid.Attestation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAggregatationByRoot", root)
	ret0, _ := ret[0].(*solid.Attestation)
	return ret0
}

// GetAggregatationByRoot indicates an expected call of GetAggregatationByRoot.
func (mr *MockAggregationPoolMockRecorder) GetAggregatationByRoot(root any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAggregatationByRoot", reflect.TypeOf((*MockAggregationPool)(nil).GetAggregatationByRoot), root)
}
