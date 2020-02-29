// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fx2y/learn-web-service-with-golang/internal/linkgraph/graph (interfaces: LinkIterator,EdgeIterator)

// Package mocks is a generated GoMock package.
package mocks

import (
	graph "github.com/fx2y/learn-web-service-with-golang/internal/linkgraph/graph"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLinkIterator is a mock of LinkIterator interface
type MockLinkIterator struct {
	ctrl     *gomock.Controller
	recorder *MockLinkIteratorMockRecorder
}

// MockLinkIteratorMockRecorder is the mock recorder for MockLinkIterator
type MockLinkIteratorMockRecorder struct {
	mock *MockLinkIterator
}

// NewMockLinkIterator creates a new mock instance
func NewMockLinkIterator(ctrl *gomock.Controller) *MockLinkIterator {
	mock := &MockLinkIterator{ctrl: ctrl}
	mock.recorder = &MockLinkIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLinkIterator) EXPECT() *MockLinkIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockLinkIterator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockLinkIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLinkIterator)(nil).Close))
}

// Error mocks base method
func (m *MockLinkIterator) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error
func (mr *MockLinkIteratorMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLinkIterator)(nil).Error))
}

// Link mocks base method
func (m *MockLinkIterator) Link() *graph.Link {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Link")
	ret0, _ := ret[0].(*graph.Link)
	return ret0
}

// Link indicates an expected call of Link
func (mr *MockLinkIteratorMockRecorder) Link() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Link", reflect.TypeOf((*MockLinkIterator)(nil).Link))
}

// Next mocks base method
func (m *MockLinkIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockLinkIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockLinkIterator)(nil).Next))
}

// MockEdgeIterator is a mock of EdgeIterator interface
type MockEdgeIterator struct {
	ctrl     *gomock.Controller
	recorder *MockEdgeIteratorMockRecorder
}

// MockEdgeIteratorMockRecorder is the mock recorder for MockEdgeIterator
type MockEdgeIteratorMockRecorder struct {
	mock *MockEdgeIterator
}

// NewMockEdgeIterator creates a new mock instance
func NewMockEdgeIterator(ctrl *gomock.Controller) *MockEdgeIterator {
	mock := &MockEdgeIterator{ctrl: ctrl}
	mock.recorder = &MockEdgeIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEdgeIterator) EXPECT() *MockEdgeIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockEdgeIterator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockEdgeIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEdgeIterator)(nil).Close))
}

// Edge mocks base method
func (m *MockEdgeIterator) Edge() *graph.Edge {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edge")
	ret0, _ := ret[0].(*graph.Edge)
	return ret0
}

// Edge indicates an expected call of Edge
func (mr *MockEdgeIteratorMockRecorder) Edge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edge", reflect.TypeOf((*MockEdgeIterator)(nil).Edge))
}

// Error mocks base method
func (m *MockEdgeIterator) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error
func (mr *MockEdgeIteratorMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockEdgeIterator)(nil).Error))
}

// Next mocks base method
func (m *MockEdgeIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockEdgeIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockEdgeIterator)(nil).Next))
}
