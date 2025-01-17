// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/architectv/estate-task/pkg/repository (interfaces: Booking)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	model "github.com/architectv/estate-task/pkg/model"
	gomock "github.com/golang/mock/gomock"
)

// MockBooking is a mock of Booking interface.
type MockBooking struct {
	ctrl     *gomock.Controller
	recorder *MockBookingMockRecorder
}

// MockBookingMockRecorder is the mock recorder for MockBooking.
type MockBookingMockRecorder struct {
	mock *MockBooking
}

// NewMockBooking creates a new mock instance.
func NewMockBooking(ctrl *gomock.Controller) *MockBooking {
	mock := &MockBooking{ctrl: ctrl}
	mock.recorder = &MockBookingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBooking) EXPECT() *MockBookingMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBooking) Create(arg0 *model.Booking) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBookingMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBooking)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockBooking) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBookingMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBooking)(nil).Delete), arg0)
}

// GetById mocks base method.
func (m *MockBooking) GetById(arg0 int) (*model.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(*model.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockBookingMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockBooking)(nil).GetById), arg0)
}

// GetByRoomId mocks base method.
func (m *MockBooking) GetByRoomId(arg0 int) ([]*model.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByRoomId", arg0)
	ret0, _ := ret[0].([]*model.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByRoomId indicates an expected call of GetByRoomId.
func (mr *MockBookingMockRecorder) GetByRoomId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByRoomId", reflect.TypeOf((*MockBooking)(nil).GetByRoomId), arg0)
}
