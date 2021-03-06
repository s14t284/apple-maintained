// Code generated by MockGen. DO NOT EDIT.
// Source: page_parser.go

// Package repository is a generated GoMock package.
package repository

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/s14t284/apple-maitained-bot/domain"
	reflect "reflect"
)

// MockPageParser is a mock of PageParser interface
type MockPageParser struct {
	ctrl     *gomock.Controller
	recorder *MockPageParserMockRecorder
}

// MockPageParserMockRecorder is the mock recorder for MockPageParser
type MockPageParserMockRecorder struct {
	mock *MockPageParser
}

// NewMockPageParser creates a new mock instance
func NewMockPageParser(ctrl *gomock.Controller) *MockPageParser {
	mock := &MockPageParser{ctrl: ctrl}
	mock.recorder = &MockPageParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPageParser) EXPECT() *MockPageParserMockRecorder {
	return m.recorder
}

// ParsePage mocks base method
func (m *MockPageParser) ParsePage(target string, page domain.Page) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParsePage", target, page)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParsePage indicates an expected call of ParsePage
func (mr *MockPageParserMockRecorder) ParsePage(target, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParsePage", reflect.TypeOf((*MockPageParser)(nil).ParsePage), target, page)
}
