// Code generated by MockGen. DO NOT EDIT.
// Source: server.go

// Package quic is a generated GoMock package.
package quic

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/lucas-clemente/quic-go/internal/protocol"
)

// MockQuicConn is a mock of QuicConn interface.
type MockQuicConn struct {
	ctrl     *gomock.Controller
	recorder *MockQuicConnMockRecorder
}

// MockQuicConnMockRecorder is the mock recorder for MockQuicConn.
type MockQuicConnMockRecorder struct {
	mock *MockQuicConn
}

// NewMockQuicConn creates a new mock instance.
func NewMockQuicConn(ctrl *gomock.Controller) *MockQuicConn {
	mock := &MockQuicConn{ctrl: ctrl}
	mock.recorder = &MockQuicConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuicConn) EXPECT() *MockQuicConnMockRecorder {
	return m.recorder
}

// AcceptStream mocks base method.
func (m *MockQuicConn) AcceptStream(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptStream", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptStream indicates an expected call of AcceptStream.
func (mr *MockQuicConnMockRecorder) AcceptStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptStream", reflect.TypeOf((*MockQuicConn)(nil).AcceptStream), arg0)
}

// AcceptUniStream mocks base method.
func (m *MockQuicConn) AcceptUniStream(arg0 context.Context) (ReceiveStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptUniStream", arg0)
	ret0, _ := ret[0].(ReceiveStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptUniStream indicates an expected call of AcceptUniStream.
func (mr *MockQuicConnMockRecorder) AcceptUniStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptUniStream", reflect.TypeOf((*MockQuicConn)(nil).AcceptUniStream), arg0)
}

// BandwidthEstimate mocks base method.
func (m *MockQuicConn) BandwidthEstimate() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BandwidthEstimate")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// BandwidthEstimate indicates an expected call of BandwidthEstimate.
func (mr *MockQuicConnMockRecorder) BandwidthEstimate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BandwidthEstimate", reflect.TypeOf((*MockQuicConn)(nil).BandwidthEstimate))
}

// CloseWithError mocks base method.
func (m *MockQuicConn) CloseWithError(arg0 ApplicationErrorCode, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseWithError", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseWithError indicates an expected call of CloseWithError.
func (mr *MockQuicConnMockRecorder) CloseWithError(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseWithError", reflect.TypeOf((*MockQuicConn)(nil).CloseWithError), arg0, arg1)
}

// ConnectionState mocks base method.
func (m *MockQuicConn) ConnectionState() ConnectionState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectionState")
	ret0, _ := ret[0].(ConnectionState)
	return ret0
}

// ConnectionState indicates an expected call of ConnectionState.
func (mr *MockQuicConnMockRecorder) ConnectionState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionState", reflect.TypeOf((*MockQuicConn)(nil).ConnectionState))
}

// Context mocks base method.
func (m *MockQuicConn) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockQuicConnMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockQuicConn)(nil).Context))
}

// GetVersion mocks base method.
func (m *MockQuicConn) GetVersion() protocol.VersionNumber {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(protocol.VersionNumber)
	return ret0
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockQuicConnMockRecorder) GetVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockQuicConn)(nil).GetVersion))
}

// HandshakeComplete mocks base method.
func (m *MockQuicConn) HandshakeComplete() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandshakeComplete")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// HandshakeComplete indicates an expected call of HandshakeComplete.
func (mr *MockQuicConnMockRecorder) HandshakeComplete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandshakeComplete", reflect.TypeOf((*MockQuicConn)(nil).HandshakeComplete))
}

// LocalAddr mocks base method.
func (m *MockQuicConn) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockQuicConnMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockQuicConn)(nil).LocalAddr))
}

// NextConnection mocks base method.
func (m *MockQuicConn) NextConnection() Connection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextConnection")
	ret0, _ := ret[0].(Connection)
	return ret0
}

// NextConnection indicates an expected call of NextConnection.
func (mr *MockQuicConnMockRecorder) NextConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextConnection", reflect.TypeOf((*MockQuicConn)(nil).NextConnection))
}

// OpenStream mocks base method.
func (m *MockQuicConn) OpenStream() (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStream")
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStream indicates an expected call of OpenStream.
func (mr *MockQuicConnMockRecorder) OpenStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStream", reflect.TypeOf((*MockQuicConn)(nil).OpenStream))
}

// OpenStreamSync mocks base method.
func (m *MockQuicConn) OpenStreamSync(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStreamSync", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStreamSync indicates an expected call of OpenStreamSync.
func (mr *MockQuicConnMockRecorder) OpenStreamSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStreamSync", reflect.TypeOf((*MockQuicConn)(nil).OpenStreamSync), arg0)
}

// OpenUniStream mocks base method.
func (m *MockQuicConn) OpenUniStream() (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStream")
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStream indicates an expected call of OpenUniStream.
func (mr *MockQuicConnMockRecorder) OpenUniStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStream", reflect.TypeOf((*MockQuicConn)(nil).OpenUniStream))
}

// OpenUniStreamSync mocks base method.
func (m *MockQuicConn) OpenUniStreamSync(arg0 context.Context) (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStreamSync", arg0)
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStreamSync indicates an expected call of OpenUniStreamSync.
func (mr *MockQuicConnMockRecorder) OpenUniStreamSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStreamSync", reflect.TypeOf((*MockQuicConn)(nil).OpenUniStreamSync), arg0)
}

// ReceiveMessage mocks base method.
func (m *MockQuicConn) ReceiveMessage() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveMessage")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage.
func (mr *MockQuicConnMockRecorder) ReceiveMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MockQuicConn)(nil).ReceiveMessage))
}

// RemoteAddr mocks base method.
func (m *MockQuicConn) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockQuicConnMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockQuicConn)(nil).RemoteAddr))
}

// SendMessage mocks base method.
func (m *MockQuicConn) SendMessage(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockQuicConnMockRecorder) SendMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockQuicConn)(nil).SendMessage), arg0)
}

// destroy mocks base method.
func (m *MockQuicConn) destroy(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "destroy", arg0)
}

// destroy indicates an expected call of destroy.
func (mr *MockQuicConnMockRecorder) destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "destroy", reflect.TypeOf((*MockQuicConn)(nil).destroy), arg0)
}

// earlyConnReady mocks base method.
func (m *MockQuicConn) earlyConnReady() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "earlyConnReady")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// earlyConnReady indicates an expected call of earlyConnReady.
func (mr *MockQuicConnMockRecorder) earlyConnReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "earlyConnReady", reflect.TypeOf((*MockQuicConn)(nil).earlyConnReady))
}

// getPerspective mocks base method.
func (m *MockQuicConn) getPerspective() protocol.Perspective {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getPerspective")
	ret0, _ := ret[0].(protocol.Perspective)
	return ret0
}

// getPerspective indicates an expected call of getPerspective.
func (mr *MockQuicConnMockRecorder) getPerspective() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getPerspective", reflect.TypeOf((*MockQuicConn)(nil).getPerspective))
}

// handlePacket mocks base method.
func (m *MockQuicConn) handlePacket(arg0 *receivedPacket) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handlePacket", arg0)
}

// handlePacket indicates an expected call of handlePacket.
func (mr *MockQuicConnMockRecorder) handlePacket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handlePacket", reflect.TypeOf((*MockQuicConn)(nil).handlePacket), arg0)
}

// run mocks base method.
func (m *MockQuicConn) run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "run")
	ret0, _ := ret[0].(error)
	return ret0
}

// run indicates an expected call of run.
func (mr *MockQuicConnMockRecorder) run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "run", reflect.TypeOf((*MockQuicConn)(nil).run))
}

// shutdown mocks base method.
func (m *MockQuicConn) shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "shutdown")
}

// shutdown indicates an expected call of shutdown.
func (mr *MockQuicConnMockRecorder) shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "shutdown", reflect.TypeOf((*MockQuicConn)(nil).shutdown))
}
