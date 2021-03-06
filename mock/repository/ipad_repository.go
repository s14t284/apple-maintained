// Code generated by MockGen. DO NOT EDIT.
// Source: ipad_repository.go

// Package repository is a generated GoMock package.
package repository

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/s14t284/apple-maitained-bot/domain/model"
	reflect "reflect"
	time "time"
)

// MockIPadRepository is a mock of IPadRepository interface
type MockIPadRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPadRepositoryMockRecorder
}

// MockIPadRepositoryMockRecorder is the mock recorder for MockIPadRepository
type MockIPadRepositoryMockRecorder struct {
	mock *MockIPadRepository
}

// NewMockIPadRepository creates a new mock instance
func NewMockIPadRepository(ctrl *gomock.Controller) *MockIPadRepository {
	mock := &MockIPadRepository{ctrl: ctrl}
	mock.recorder = &MockIPadRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPadRepository) EXPECT() *MockIPadRepositoryMockRecorder {
	return m.recorder
}

// FindIPad mocks base method
func (m *MockIPadRepository) FindIPad(param *model.IPadRequestParam) (model.IPads, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIPad", param)
	ret0, _ := ret[0].(model.IPads)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIPad indicates an expected call of FindIPad
func (mr *MockIPadRepositoryMockRecorder) FindIPad(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIPad", reflect.TypeOf((*MockIPadRepository)(nil).FindIPad), param)
}

// FindIPadAll mocks base method
func (m *MockIPadRepository) FindIPadAll() (model.IPads, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIPadAll")
	ret0, _ := ret[0].(model.IPads)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIPadAll indicates an expected call of FindIPadAll
func (mr *MockIPadRepositoryMockRecorder) FindIPadAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIPadAll", reflect.TypeOf((*MockIPadRepository)(nil).FindIPadAll))
}

// FindByURL mocks base method
func (m *MockIPadRepository) FindByURL(url string) (*model.IPad, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByURL", url)
	ret0, _ := ret[0].(*model.IPad)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByURL indicates an expected call of FindByURL
func (mr *MockIPadRepositoryMockRecorder) FindByURL(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByURL", reflect.TypeOf((*MockIPadRepository)(nil).FindByURL), url)
}

// IsExist mocks base method
func (m *MockIPadRepository) IsExist(ipad *model.IPad) (bool, uint, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExist", ipad)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(uint)
	ret2, _ := ret[2].(time.Time)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// IsExist indicates an expected call of IsExist
func (mr *MockIPadRepositoryMockRecorder) IsExist(ipad interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExist", reflect.TypeOf((*MockIPadRepository)(nil).IsExist), ipad)
}

// AddIPad mocks base method
func (m *MockIPadRepository) AddIPad(ipad *model.IPad) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddIPad", ipad)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddIPad indicates an expected call of AddIPad
func (mr *MockIPadRepositoryMockRecorder) AddIPad(ipad interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIPad", reflect.TypeOf((*MockIPadRepository)(nil).AddIPad), ipad)
}

// UpdateIPad mocks base method
func (m *MockIPadRepository) UpdateIPad(ipad *model.IPad) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIPad", ipad)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIPad indicates an expected call of UpdateIPad
func (mr *MockIPadRepositoryMockRecorder) UpdateIPad(ipad interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIPad", reflect.TypeOf((*MockIPadRepository)(nil).UpdateIPad), ipad)
}

// UpdateAllSoldTemporary mocks base method
func (m *MockIPadRepository) UpdateAllSoldTemporary() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSoldTemporary")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAllSoldTemporary indicates an expected call of UpdateAllSoldTemporary
func (mr *MockIPadRepositoryMockRecorder) UpdateAllSoldTemporary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSoldTemporary", reflect.TypeOf((*MockIPadRepository)(nil).UpdateAllSoldTemporary))
}

// RemoveIPad mocks base method
func (m *MockIPadRepository) RemoveIPad(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveIPad", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveIPad indicates an expected call of RemoveIPad
func (mr *MockIPadRepositoryMockRecorder) RemoveIPad(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIPad", reflect.TypeOf((*MockIPadRepository)(nil).RemoveIPad), id)
}
