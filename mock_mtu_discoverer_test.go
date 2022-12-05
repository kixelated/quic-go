// Code generated by MockGen. DO NOT EDIT.
// Source: mtu_discoverer.go

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	ackhandler "github.com/kixelated/quic-go/internal/ackhandler"
	protocol "github.com/kixelated/quic-go/internal/protocol"
)

// MockMtuDiscoverer is a mock of MtuDiscoverer interface.
type MockMtuDiscoverer struct {
	ctrl     *gomock.Controller
	recorder *MockMtuDiscovererMockRecorder
}

// MockMtuDiscovererMockRecorder is the mock recorder for MockMtuDiscoverer.
type MockMtuDiscovererMockRecorder struct {
	mock *MockMtuDiscoverer
}

// NewMockMtuDiscoverer creates a new mock instance.
func NewMockMtuDiscoverer(ctrl *gomock.Controller) *MockMtuDiscoverer {
	mock := &MockMtuDiscoverer{ctrl: ctrl}
	mock.recorder = &MockMtuDiscovererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMtuDiscoverer) EXPECT() *MockMtuDiscovererMockRecorder {
	return m.recorder
}

// GetPing mocks base method.
func (m *MockMtuDiscoverer) GetPing() (ackhandler.Frame, protocol.ByteCount) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPing")
	ret0, _ := ret[0].(ackhandler.Frame)
	ret1, _ := ret[1].(protocol.ByteCount)
	return ret0, ret1
}

// GetPing indicates an expected call of GetPing.
func (mr *MockMtuDiscovererMockRecorder) GetPing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPing", reflect.TypeOf((*MockMtuDiscoverer)(nil).GetPing))
}

// ShouldSendProbe mocks base method.
func (m *MockMtuDiscoverer) ShouldSendProbe(now time.Time) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldSendProbe", now)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldSendProbe indicates an expected call of ShouldSendProbe.
func (mr *MockMtuDiscovererMockRecorder) ShouldSendProbe(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldSendProbe", reflect.TypeOf((*MockMtuDiscoverer)(nil).ShouldSendProbe), now)
}
