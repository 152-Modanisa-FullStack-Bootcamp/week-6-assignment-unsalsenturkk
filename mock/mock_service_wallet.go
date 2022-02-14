// Code generated by MockGen. DO NOT EDIT.
// Source: service/service.go

// Package mock_service is a generated GoMock package.
package mock

import (
	reflect "reflect"
	models "week-6-assignment/models"

	gomock "github.com/golang/mock/gomock"
)

// MockIWallet is a mock of IWallet interface.
type MockIWallet struct {
	ctrl     *gomock.Controller
	recorder *MockIWalletMockRecorder
}

// MockIWalletMockRecorder is the mock recorder for MockIWallet.
type MockIWalletMockRecorder struct {
	mock *MockIWallet
}

// NewMockIWallet creates a new mock instance.
func NewMockIWallet(ctrl *gomock.Controller) *MockIWallet {
	mock := &MockIWallet{ctrl: ctrl}
	mock.recorder = &MockIWalletMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWallet) EXPECT() *MockIWalletMockRecorder {
	return m.recorder
}

// GetWallet mocks base method.
func (m *MockIWallet) GetWallet() (*models.ServiceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWallet")
	ret0, _ := ret[0].(*models.ServiceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWallet indicates an expected call of GetWallet.
func (mr *MockIWalletMockRecorder) GetWallet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWallet", reflect.TypeOf((*MockIWallet)(nil).GetWallet))
}

// GetWalletByUsername mocks base method.
func (m *MockIWallet) GetWalletByUsername(username string) (*models.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletByUsername", username)
	ret0, _ := ret[0].(*models.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletByUsername indicates an expected call of GetWalletByUsername.
func (mr *MockIWalletMockRecorder) GetWalletByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletByUsername", reflect.TypeOf((*MockIWallet)(nil).GetWalletByUsername), username)
}

// PostWalletByUsername mocks base method.
func (m *MockIWallet) PostWalletByUsername(username string, balance float64) (*models.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostWalletByUsername", username, balance)
	ret0, _ := ret[0].(*models.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostWalletByUsername indicates an expected call of PostWalletByUsername.
func (mr *MockIWalletMockRecorder) PostWalletByUsername(username, balance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostWalletByUsername", reflect.TypeOf((*MockIWallet)(nil).PostWalletByUsername), username, balance)
}

// PutWalletByUsername mocks base method.
func (m *MockIWallet) PutWalletByUsername(username string) (*models.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutWalletByUsername", username)
	ret0, _ := ret[0].(*models.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutWalletByUsername indicates an expected call of PutWalletByUsername.
func (mr *MockIWalletMockRecorder) PutWalletByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutWalletByUsername", reflect.TypeOf((*MockIWallet)(nil).PutWalletByUsername), username)
}
