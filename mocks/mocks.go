// Code generated by MockGen. DO NOT EDIT.
// Source: proto/sampleManager_grpc.pb.go

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	context "context"
	reflect "reflect"
	proto "sampleManager/proto"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockSampleManagerClient is a mock of SampleManagerClient interface.
type MockSampleManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockSampleManagerClientMockRecorder
}

// MockSampleManagerClientMockRecorder is the mock recorder for MockSampleManagerClient.
type MockSampleManagerClientMockRecorder struct {
	mock *MockSampleManagerClient
}

// NewMockSampleManagerClient creates a new mock instance.
func NewMockSampleManagerClient(ctrl *gomock.Controller) *MockSampleManagerClient {
	mock := &MockSampleManagerClient{ctrl: ctrl}
	mock.recorder = &MockSampleManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSampleManagerClient) EXPECT() *MockSampleManagerClientMockRecorder {
	return m.recorder
}

// CreateMapping mocks base method.
func (m *MockSampleManagerClient) CreateMapping(ctx context.Context, in *proto.CreateRequest, opts ...grpc.CallOption) (*proto.CreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMapping", varargs...)
	ret0, _ := ret[0].(*proto.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMapping indicates an expected call of CreateMapping.
func (mr *MockSampleManagerClientMockRecorder) CreateMapping(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMapping", reflect.TypeOf((*MockSampleManagerClient)(nil).CreateMapping), varargs...)
}

// GetSampleId mocks base method.
func (m *MockSampleManagerClient) GetSampleId(ctx context.Context, in *proto.GetRequest, opts ...grpc.CallOption) (*proto.GetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSampleId", varargs...)
	ret0, _ := ret[0].(*proto.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSampleId indicates an expected call of GetSampleId.
func (mr *MockSampleManagerClientMockRecorder) GetSampleId(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSampleId", reflect.TypeOf((*MockSampleManagerClient)(nil).GetSampleId), varargs...)
}

// MockSampleManagerServer is a mock of SampleManagerServer interface.
type MockSampleManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockSampleManagerServerMockRecorder
}

// MockSampleManagerServerMockRecorder is the mock recorder for MockSampleManagerServer.
type MockSampleManagerServerMockRecorder struct {
	mock *MockSampleManagerServer
}

// NewMockSampleManagerServer creates a new mock instance.
func NewMockSampleManagerServer(ctrl *gomock.Controller) *MockSampleManagerServer {
	mock := &MockSampleManagerServer{ctrl: ctrl}
	mock.recorder = &MockSampleManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSampleManagerServer) EXPECT() *MockSampleManagerServerMockRecorder {
	return m.recorder
}

// CreateMapping mocks base method.
func (m *MockSampleManagerServer) CreateMapping(arg0 context.Context, arg1 *proto.CreateRequest) (*proto.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMapping", arg0, arg1)
	ret0, _ := ret[0].(*proto.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMapping indicates an expected call of CreateMapping.
func (mr *MockSampleManagerServerMockRecorder) CreateMapping(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMapping", reflect.TypeOf((*MockSampleManagerServer)(nil).CreateMapping), arg0, arg1)
}

// GetSampleId mocks base method.
func (m *MockSampleManagerServer) GetSampleId(arg0 context.Context, arg1 *proto.GetRequest) (*proto.GetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSampleId", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSampleId indicates an expected call of GetSampleId.
func (mr *MockSampleManagerServerMockRecorder) GetSampleId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSampleId", reflect.TypeOf((*MockSampleManagerServer)(nil).GetSampleId), arg0, arg1)
}

// mustEmbedUnimplementedSampleManagerServer mocks base method.
func (m *MockSampleManagerServer) mustEmbedUnimplementedSampleManagerServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSampleManagerServer")
}

// mustEmbedUnimplementedSampleManagerServer indicates an expected call of mustEmbedUnimplementedSampleManagerServer.
func (mr *MockSampleManagerServerMockRecorder) mustEmbedUnimplementedSampleManagerServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSampleManagerServer", reflect.TypeOf((*MockSampleManagerServer)(nil).mustEmbedUnimplementedSampleManagerServer))
}

// MockUnsafeSampleManagerServer is a mock of UnsafeSampleManagerServer interface.
type MockUnsafeSampleManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeSampleManagerServerMockRecorder
}

// MockUnsafeSampleManagerServerMockRecorder is the mock recorder for MockUnsafeSampleManagerServer.
type MockUnsafeSampleManagerServerMockRecorder struct {
	mock *MockUnsafeSampleManagerServer
}

// NewMockUnsafeSampleManagerServer creates a new mock instance.
func NewMockUnsafeSampleManagerServer(ctrl *gomock.Controller) *MockUnsafeSampleManagerServer {
	mock := &MockUnsafeSampleManagerServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeSampleManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeSampleManagerServer) EXPECT() *MockUnsafeSampleManagerServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedSampleManagerServer mocks base method.
func (m *MockUnsafeSampleManagerServer) mustEmbedUnimplementedSampleManagerServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSampleManagerServer")
}

// mustEmbedUnimplementedSampleManagerServer indicates an expected call of mustEmbedUnimplementedSampleManagerServer.
func (mr *MockUnsafeSampleManagerServerMockRecorder) mustEmbedUnimplementedSampleManagerServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSampleManagerServer", reflect.TypeOf((*MockUnsafeSampleManagerServer)(nil).mustEmbedUnimplementedSampleManagerServer))
}
