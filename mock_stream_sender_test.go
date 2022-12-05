// Code generated by MockGen. DO NOT EDIT.
// Source: stream.go

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/kixelated/quic-go/internal/protocol"
	wire "github.com/kixelated/quic-go/internal/wire"
)

// MockStreamSender is a mock of StreamSender interface.
type MockStreamSender struct {
	ctrl     *gomock.Controller
	recorder *MockStreamSenderMockRecorder
}

// MockStreamSenderMockRecorder is the mock recorder for MockStreamSender.
type MockStreamSenderMockRecorder struct {
	mock *MockStreamSender
}

// NewMockStreamSender creates a new mock instance.
func NewMockStreamSender(ctrl *gomock.Controller) *MockStreamSender {
	mock := &MockStreamSender{ctrl: ctrl}
	mock.recorder = &MockStreamSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamSender) EXPECT() *MockStreamSenderMockRecorder {
	return m.recorder
}

// onHasStreamData mocks base method.
func (m *MockStreamSender) onHasStreamData(arg0 protocol.StreamID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "onHasStreamData", arg0)
}

// onHasStreamData indicates an expected call of onHasStreamData.
func (mr *MockStreamSenderMockRecorder) onHasStreamData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "onHasStreamData", reflect.TypeOf((*MockStreamSender)(nil).onHasStreamData), arg0)
}

// onStreamCompleted mocks base method.
func (m *MockStreamSender) onStreamCompleted(arg0 protocol.StreamID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "onStreamCompleted", arg0)
}

// onStreamCompleted indicates an expected call of onStreamCompleted.
func (mr *MockStreamSenderMockRecorder) onStreamCompleted(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "onStreamCompleted", reflect.TypeOf((*MockStreamSender)(nil).onStreamCompleted), arg0)
}

// queueControlFrame mocks base method.
func (m *MockStreamSender) queueControlFrame(arg0 wire.Frame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "queueControlFrame", arg0)
}

// queueControlFrame indicates an expected call of queueControlFrame.
func (mr *MockStreamSenderMockRecorder) queueControlFrame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "queueControlFrame", reflect.TypeOf((*MockStreamSender)(nil).queueControlFrame), arg0)
}
