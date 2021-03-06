// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anurag120799/patient-management/internal/service (interfaces: Patient)

// Package service is a generated GoMock package.
package service

import (
	gofr "developer.zopsmart.com/go/gofr/pkg/gofr"
	gomock "github.com/golang/mock/gomock"
	models "github.com/anurag120799/patient-management/internal/models"
	reflect "reflect"
)

// MockPatient is a mock of Patient interface
type MockPatient struct {
	ctrl     *gomock.Controller
	recorder *MockPatientMockRecorder
}

// MockPatientMockRecorder is the mock recorder for MockPatient
type MockPatientMockRecorder struct {
	mock *MockPatient
}

// NewMockPatient creates a new mock instance
func NewMockPatient(ctrl *gomock.Controller) *MockPatient {
	mock := &MockPatient{ctrl: ctrl}
	mock.recorder = &MockPatientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPatient) EXPECT() *MockPatientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockPatient) Create(arg0 *gofr.Context, arg1 *models.Patient) (*models.Patient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*models.Patient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockPatientMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPatient)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockPatient) Delete(arg0 *gofr.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPatientMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPatient)(nil).Delete), arg0, arg1)
}

// Get mocks base method
func (m *MockPatient) Get(arg0 *gofr.Context) ([]*models.Patient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]*models.Patient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPatientMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPatient)(nil).Get), arg0)
}

// GetByID mocks base method
func (m *MockPatient) GetByID(arg0 *gofr.Context, arg1 string) (*models.Patient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*models.Patient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockPatientMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockPatient)(nil).GetByID), arg0, arg1)
}

// Update mocks base method
func (m *MockPatient) Update(arg0 *gofr.Context, arg1 string, arg2 *models.Patient) (*models.Patient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Patient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockPatientMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPatient)(nil).Update), arg0, arg1, arg2)
}
