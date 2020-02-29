// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fx2y/learn-web-service-with-golang/internal/crawler (interfaces: URLGetter,PrivateNetworkDetector,Graph,Indexer)

// Package mocks is a generated GoMock package.
package mocks

import (
	graph "github.com/fx2y/learn-web-service-with-golang/internal/linkgraph/graph"
	index "github.com/fx2y/learn-web-service-with-golang/internal/textindexer/index"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	http "net/http"
	reflect "reflect"
	time "time"
)

// MockURLGetter is a mock of URLGetter interface
type MockURLGetter struct {
	ctrl     *gomock.Controller
	recorder *MockURLGetterMockRecorder
}

// MockURLGetterMockRecorder is the mock recorder for MockURLGetter
type MockURLGetterMockRecorder struct {
	mock *MockURLGetter
}

// NewMockURLGetter creates a new mock instance
func NewMockURLGetter(ctrl *gomock.Controller) *MockURLGetter {
	mock := &MockURLGetter{ctrl: ctrl}
	mock.recorder = &MockURLGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockURLGetter) EXPECT() *MockURLGetterMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockURLGetter) Get(arg0 string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockURLGetterMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockURLGetter)(nil).Get), arg0)
}

// MockPrivateNetworkDetector is a mock of PrivateNetworkDetector interface
type MockPrivateNetworkDetector struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateNetworkDetectorMockRecorder
}

// MockPrivateNetworkDetectorMockRecorder is the mock recorder for MockPrivateNetworkDetector
type MockPrivateNetworkDetectorMockRecorder struct {
	mock *MockPrivateNetworkDetector
}

// NewMockPrivateNetworkDetector creates a new mock instance
func NewMockPrivateNetworkDetector(ctrl *gomock.Controller) *MockPrivateNetworkDetector {
	mock := &MockPrivateNetworkDetector{ctrl: ctrl}
	mock.recorder = &MockPrivateNetworkDetectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrivateNetworkDetector) EXPECT() *MockPrivateNetworkDetectorMockRecorder {
	return m.recorder
}

// IsPrivate mocks base method
func (m *MockPrivateNetworkDetector) IsPrivate(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPrivate", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsPrivate indicates an expected call of IsPrivate
func (mr *MockPrivateNetworkDetectorMockRecorder) IsPrivate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPrivate", reflect.TypeOf((*MockPrivateNetworkDetector)(nil).IsPrivate), arg0)
}

// MockGraph is a mock of Graph interface
type MockGraph struct {
	ctrl     *gomock.Controller
	recorder *MockGraphMockRecorder
}

// MockGraphMockRecorder is the mock recorder for MockGraph
type MockGraphMockRecorder struct {
	mock *MockGraph
}

// NewMockGraph creates a new mock instance
func NewMockGraph(ctrl *gomock.Controller) *MockGraph {
	mock := &MockGraph{ctrl: ctrl}
	mock.recorder = &MockGraphMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGraph) EXPECT() *MockGraphMockRecorder {
	return m.recorder
}

// RemoveStaleEdges mocks base method
func (m *MockGraph) RemoveStaleEdges(arg0 uuid.UUID, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveStaleEdges", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveStaleEdges indicates an expected call of RemoveStaleEdges
func (mr *MockGraphMockRecorder) RemoveStaleEdges(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveStaleEdges", reflect.TypeOf((*MockGraph)(nil).RemoveStaleEdges), arg0, arg1)
}

// UpsertEdge mocks base method
func (m *MockGraph) UpsertEdge(arg0 *graph.Edge) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertEdge", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertEdge indicates an expected call of UpsertEdge
func (mr *MockGraphMockRecorder) UpsertEdge(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertEdge", reflect.TypeOf((*MockGraph)(nil).UpsertEdge), arg0)
}

// UpsertLink mocks base method
func (m *MockGraph) UpsertLink(arg0 *graph.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertLink", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertLink indicates an expected call of UpsertLink
func (mr *MockGraphMockRecorder) UpsertLink(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertLink", reflect.TypeOf((*MockGraph)(nil).UpsertLink), arg0)
}

// MockIndexer is a mock of Indexer interface
type MockIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockIndexerMockRecorder
}

// MockIndexerMockRecorder is the mock recorder for MockIndexer
type MockIndexerMockRecorder struct {
	mock *MockIndexer
}

// NewMockIndexer creates a new mock instance
func NewMockIndexer(ctrl *gomock.Controller) *MockIndexer {
	mock := &MockIndexer{ctrl: ctrl}
	mock.recorder = &MockIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexer) EXPECT() *MockIndexerMockRecorder {
	return m.recorder
}

// Index mocks base method
func (m *MockIndexer) Index(arg0 *index.Document) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Index indicates an expected call of Index
func (mr *MockIndexerMockRecorder) Index(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockIndexer)(nil).Index), arg0)
}
