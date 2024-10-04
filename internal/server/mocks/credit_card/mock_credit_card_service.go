// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/msmkdenis/yap-gophkeeper/internal/server/credit_card/api/v1/grpchandlers (interfaces: CreditCardService)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination=internal/server/mocks/credit_card/mock_credit_card_service.go -package=mocks github.com/msmkdenis/yap-gophkeeper/internal/server/credit_card/api/v1/grpchandlers CreditCardService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	specification "github.com/msmkdenis/yap-gophkeeper/internal/server/credit_card/specification"
	model "github.com/msmkdenis/yap-gophkeeper/internal/server/model"
	gomock "go.uber.org/mock/gomock"
)

// MockCreditCardService is a mock of CreditCardService interface.
type MockCreditCardService struct {
	ctrl     *gomock.Controller
	recorder *MockCreditCardServiceMockRecorder
}

// MockCreditCardServiceMockRecorder is the mock recorder for MockCreditCardService.
type MockCreditCardServiceMockRecorder struct {
	mock *MockCreditCardService
}

// NewMockCreditCardService creates a new mock instance.
func NewMockCreditCardService(ctrl *gomock.Controller) *MockCreditCardService {
	mock := &MockCreditCardService{ctrl: ctrl}
	mock.recorder = &MockCreditCardServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreditCardService) EXPECT() *MockCreditCardServiceMockRecorder {
	return m.recorder
}

// LoadAllCreditCard mocks base method.
func (m *MockCreditCardService) LoadAllCreditCard(arg0 context.Context, arg1 specification.CreditCardSpecification) ([]model.CreditCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadAllCreditCard", arg0, arg1)
	ret0, _ := ret[0].([]model.CreditCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadAllCreditCard indicates an expected call of LoadAllCreditCard.
func (mr *MockCreditCardServiceMockRecorder) LoadAllCreditCard(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadAllCreditCard", reflect.TypeOf((*MockCreditCardService)(nil).LoadAllCreditCard), arg0, arg1)
}

// SaveCreditCard mocks base method.
func (m *MockCreditCardService) SaveCreditCard(arg0 context.Context, arg1 model.CreditCardPostRequest) (model.CreditCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCreditCard", arg0, arg1)
	ret0, _ := ret[0].(model.CreditCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveCreditCard indicates an expected call of SaveCreditCard.
func (mr *MockCreditCardServiceMockRecorder) SaveCreditCard(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCreditCard", reflect.TypeOf((*MockCreditCardService)(nil).SaveCreditCard), arg0, arg1)
}
