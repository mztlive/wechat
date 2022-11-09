// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/wechat/contracts.go

// Package wechat is a generated GoMock package.
package wechat

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	oa "github.com/mztlive/wechat/oa"
)

// MockOfficialAccountQR is a mock of OfficialAccountQR interface.
type MockOfficialAccountQR struct {
	ctrl     *gomock.Controller
	recorder *MockOfficialAccountQRMockRecorder
}

// MockOfficialAccountQRMockRecorder is the mock recorder for MockOfficialAccountQR.
type MockOfficialAccountQRMockRecorder struct {
	mock *MockOfficialAccountQR
}

// NewMockOfficialAccountQR creates a new mock instance.
func NewMockOfficialAccountQR(ctrl *gomock.Controller) *MockOfficialAccountQR {
	mock := &MockOfficialAccountQR{ctrl: ctrl}
	mock.recorder = &MockOfficialAccountQRMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOfficialAccountQR) EXPECT() *MockOfficialAccountQRMockRecorder {
	return m.recorder
}

// GenTemporary mocks base method.
func (m *MockOfficialAccountQR) GenTemporary(ctx context.Context, officialAccounts []oa.OfficialAccountKey, sceneStr string) (oa.QrCodes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenTemporary", ctx, officialAccounts, sceneStr)
	ret0, _ := ret[0].(oa.QrCodes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenTemporary indicates an expected call of GenTemporary.
func (mr *MockOfficialAccountQRMockRecorder) GenTemporary(ctx, officialAccounts, sceneStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenTemporary", reflect.TypeOf((*MockOfficialAccountQR)(nil).GenTemporary), ctx, officialAccounts, sceneStr)
}