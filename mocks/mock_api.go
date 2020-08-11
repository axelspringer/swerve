// Code generated by MockGen. DO NOT EDIT.
// Source: api/interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockModelAdapter is a mock of ModelAdapter interface.
type MockModelAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockModelAdapterMockRecorder
}

// MockModelAdapterMockRecorder is the mock recorder for MockModelAdapter.
type MockModelAdapterMockRecorder struct {
	mock *MockModelAdapter
}

// NewMockModelAdapter creates a new mock instance.
func NewMockModelAdapter(ctrl *gomock.Controller) *MockModelAdapter {
	mock := &MockModelAdapter{ctrl: ctrl}
	mock.recorder = &MockModelAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelAdapter) EXPECT() *MockModelAdapterMockRecorder {
	return m.recorder
}

// CreateRedirectFromJSON mocks base method.
func (m *MockModelAdapter) CreateRedirectFromJSON(jsonStr []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRedirectFromJSON", jsonStr)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRedirectFromJSON indicates an expected call of CreateRedirectFromJSON.
func (mr *MockModelAdapterMockRecorder) CreateRedirectFromJSON(jsonStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRedirectFromJSON", reflect.TypeOf((*MockModelAdapter)(nil).CreateRedirectFromJSON), jsonStr)
}

// UpdateRedirectByDomainWithJSON mocks base method.
func (m *MockModelAdapter) UpdateRedirectByDomainWithJSON(domain string, jsonStr []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRedirectByDomainWithJSON", domain, jsonStr)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRedirectByDomainWithJSON indicates an expected call of UpdateRedirectByDomainWithJSON.
func (mr *MockModelAdapterMockRecorder) UpdateRedirectByDomainWithJSON(domain, jsonStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRedirectByDomainWithJSON", reflect.TypeOf((*MockModelAdapter)(nil).UpdateRedirectByDomainWithJSON), domain, jsonStr)
}

// DeleteRedirectByDomain mocks base method.
func (m *MockModelAdapter) DeleteRedirectByDomain(domain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRedirectByDomain", domain)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRedirectByDomain indicates an expected call of DeleteRedirectByDomain.
func (mr *MockModelAdapterMockRecorder) DeleteRedirectByDomain(domain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRedirectByDomain", reflect.TypeOf((*MockModelAdapter)(nil).DeleteRedirectByDomain), domain)
}

// GetRedirectByDomainAsJSON mocks base method.
func (m *MockModelAdapter) GetRedirectByDomainAsJSON(domain string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRedirectByDomainAsJSON", domain)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRedirectByDomainAsJSON indicates an expected call of GetRedirectByDomainAsJSON.
func (mr *MockModelAdapterMockRecorder) GetRedirectByDomainAsJSON(domain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRedirectByDomainAsJSON", reflect.TypeOf((*MockModelAdapter)(nil).GetRedirectByDomainAsJSON), domain)
}

// ImportRedirectsFromJSON mocks base method.
func (m *MockModelAdapter) ImportRedirectsFromJSON(jsonStr []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportRedirectsFromJSON", jsonStr)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportRedirectsFromJSON indicates an expected call of ImportRedirectsFromJSON.
func (mr *MockModelAdapterMockRecorder) ImportRedirectsFromJSON(jsonStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportRedirectsFromJSON", reflect.TypeOf((*MockModelAdapter)(nil).ImportRedirectsFromJSON), jsonStr)
}

// ExportRedirectsAsJSON mocks base method.
func (m *MockModelAdapter) ExportRedirectsAsJSON() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportRedirectsAsJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportRedirectsAsJSON indicates an expected call of ExportRedirectsAsJSON.
func (mr *MockModelAdapterMockRecorder) ExportRedirectsAsJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportRedirectsAsJSON", reflect.TypeOf((*MockModelAdapter)(nil).ExportRedirectsAsJSON))
}

// GetRedirectsPaginatedAsJSON mocks base method.
func (m *MockModelAdapter) GetRedirectsPaginatedAsJSON(cursor string) ([]byte, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRedirectsPaginatedAsJSON", cursor)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRedirectsPaginatedAsJSON indicates an expected call of GetRedirectsPaginatedAsJSON.
func (mr *MockModelAdapterMockRecorder) GetRedirectsPaginatedAsJSON(cursor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRedirectsPaginatedAsJSON", reflect.TypeOf((*MockModelAdapter)(nil).GetRedirectsPaginatedAsJSON), cursor)
}

// CheckPasswordFromJSON mocks base method.
func (m *MockModelAdapter) CheckPasswordFromJSON(jsonStr []byte, secret string) (string, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPasswordFromJSON", jsonStr, secret)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CheckPasswordFromJSON indicates an expected call of CheckPasswordFromJSON.
func (mr *MockModelAdapterMockRecorder) CheckPasswordFromJSON(jsonStr, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPasswordFromJSON", reflect.TypeOf((*MockModelAdapter)(nil).CheckPasswordFromJSON), jsonStr, secret)
}

// CheckToken mocks base method.
func (m *MockModelAdapter) CheckToken(token, secret string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckToken", token, secret)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckToken indicates an expected call of CheckToken.
func (mr *MockModelAdapterMockRecorder) CheckToken(token, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckToken", reflect.TypeOf((*MockModelAdapter)(nil).CheckToken), token, secret)
}
