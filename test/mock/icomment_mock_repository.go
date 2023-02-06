// Code generated by MockGen. DO NOT EDIT.
// Source: q1/repository/interface (interfaces: ICommentRep)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	repository "q1/models/repository"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICommentRep is a mock of ICommentRep interface.
type MockICommentRep struct {
	ctrl     *gomock.Controller
	recorder *MockICommentRepMockRecorder
}

// MockICommentRepMockRecorder is the mock recorder for MockICommentRep.
type MockICommentRepMockRecorder struct {
	mock *MockICommentRep
}

// NewMockICommentRep creates a new mock instance.
func NewMockICommentRep(ctrl *gomock.Controller) *MockICommentRep {
	mock := &MockICommentRep{ctrl: ctrl}
	mock.recorder = &MockICommentRepMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICommentRep) EXPECT() *MockICommentRepMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockICommentRep) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockICommentRepMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockICommentRep)(nil).Close))
}

// Delete mocks base method.
func (m *MockICommentRep) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockICommentRepMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockICommentRep)(nil).Delete), arg0, arg1)
}

// Find mocks base method.
func (m *MockICommentRep) Find(arg0 context.Context, arg1 string) (*repository.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*repository.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockICommentRepMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockICommentRep)(nil).Find), arg0, arg1)
}

// Insert mocks base method.
func (m *MockICommentRep) Insert(arg0 context.Context, arg1 *repository.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockICommentRepMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockICommentRep)(nil).Insert), arg0, arg1)
}

// Updates mocks base method.
func (m *MockICommentRep) Updates(arg0 context.Context, arg1 *repository.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Updates indicates an expected call of Updates.
func (mr *MockICommentRepMockRecorder) Updates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*MockICommentRep)(nil).Updates), arg0, arg1)
}
