// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/architectv/estate-task/pkg/repository (interfaces: Room)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	model "github.com/architectv/estate-task/pkg/model"
	gomock "github.com/golang/mock/gomock"
)

// MockRoom is a mock of Room interface.
type MockRoom struct {
	ctrl     *gomock.Controller
	recorder *MockRoomMockRecorder
}

// MockRoomMockRecorder is the mock recorder for MockRoom.
type MockRoomMockRecorder struct {
	mock *MockRoom
}

// NewMockRoom creates a new mock instance.
func NewMockRoom(ctrl *gomock.Controller) *MockRoom {
	mock := &MockRoom{ctrl: ctrl}
	mock.recorder = &MockRoomMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoom) EXPECT() *MockRoomMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRoom) Create(arg0 *model.Room) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRoomMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoom)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockRoom) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRoomMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoom)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockRoom) GetAll(arg0 string, arg1 bool) ([]*model.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1)
	ret0, _ := ret[0].([]*model.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRoomMockRecorder) GetAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRoom)(nil).GetAll), arg0, arg1)
}

// GetById mocks base method.
func (m *MockRoom) GetById(arg0 int) (*model.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(*model.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockRoomMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockRoom)(nil).GetById), arg0)
}
