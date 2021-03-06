// Code generated by MockGen. DO NOT EDIT.
// Source: db/querier.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/zweimach/xendit-test/db"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreateCommentByOrganizationLogin mocks base method.
func (m *MockQuerier) CreateCommentByOrganizationLogin(ctx context.Context, arg db.CreateCommentByOrganizationLoginParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommentByOrganizationLogin", ctx, arg)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCommentByOrganizationLogin indicates an expected call of CreateCommentByOrganizationLogin.
func (mr *MockQuerierMockRecorder) CreateCommentByOrganizationLogin(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommentByOrganizationLogin", reflect.TypeOf((*MockQuerier)(nil).CreateCommentByOrganizationLogin), ctx, arg)
}

// DeleteCommentsByOrganizationLogin mocks base method.
func (m *MockQuerier) DeleteCommentsByOrganizationLogin(ctx context.Context, login string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCommentsByOrganizationLogin", ctx, login)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCommentsByOrganizationLogin indicates an expected call of DeleteCommentsByOrganizationLogin.
func (mr *MockQuerierMockRecorder) DeleteCommentsByOrganizationLogin(ctx, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCommentsByOrganizationLogin", reflect.TypeOf((*MockQuerier)(nil).DeleteCommentsByOrganizationLogin), ctx, login)
}

// ListCommentsByOrganizationLogin mocks base method.
func (m *MockQuerier) ListCommentsByOrganizationLogin(ctx context.Context, arg db.ListCommentsByOrganizationLoginParams) ([]db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommentsByOrganizationLogin", ctx, arg)
	ret0, _ := ret[0].([]db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommentsByOrganizationLogin indicates an expected call of ListCommentsByOrganizationLogin.
func (mr *MockQuerierMockRecorder) ListCommentsByOrganizationLogin(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommentsByOrganizationLogin", reflect.TypeOf((*MockQuerier)(nil).ListCommentsByOrganizationLogin), ctx, arg)
}

// ListUsersByOrganizationLogin mocks base method.
func (m *MockQuerier) ListUsersByOrganizationLogin(ctx context.Context, arg db.ListUsersByOrganizationLoginParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsersByOrganizationLogin", ctx, arg)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsersByOrganizationLogin indicates an expected call of ListUsersByOrganizationLogin.
func (mr *MockQuerierMockRecorder) ListUsersByOrganizationLogin(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsersByOrganizationLogin", reflect.TypeOf((*MockQuerier)(nil).ListUsersByOrganizationLogin), ctx, arg)
}
