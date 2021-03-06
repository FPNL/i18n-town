// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/core/service/word.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	entity "github.com/FPNL/i18n-town/src/core/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockIWordService is a mock of IWordService interface.
type MockIWordService struct {
	ctrl     *gomock.Controller
	recorder *MockIWordServiceMockRecorder
}

// MockIWordServiceMockRecorder is the mock recorder for MockIWordService.
type MockIWordServiceMockRecorder struct {
	mock *MockIWordService
}

// NewMockIWordService creates a new mock instance.
func NewMockIWordService(ctrl *gomock.Controller) *MockIWordService {
	mock := &MockIWordService{ctrl: ctrl}
	mock.recorder = &MockIWordServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWordService) EXPECT() *MockIWordServiceMockRecorder {
	return m.recorder
}

// AddManyWords mocks base method.
func (m *MockIWordService) AddManyWords(arg0 []entity.Word) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddManyWords", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddManyWords indicates an expected call of AddManyWords.
func (mr *MockIWordServiceMockRecorder) AddManyWords(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddManyWords", reflect.TypeOf((*MockIWordService)(nil).AddManyWords), arg0)
}

// AddOneWord mocks base method.
func (m *MockIWordService) AddOneWord(arg0 *entity.Word) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOneWord", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOneWord indicates an expected call of AddOneWord.
func (mr *MockIWordServiceMockRecorder) AddOneWord(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOneWord", reflect.TypeOf((*MockIWordService)(nil).AddOneWord), arg0)
}

// DeleteAll mocks base method.
func (m *MockIWordService) DeleteAll() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockIWordServiceMockRecorder) DeleteAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockIWordService)(nil).DeleteAll))
}

// DeleteManyWord mocks base method.
func (m *MockIWordService) DeleteManyWord(arg0 []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteManyWord", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteManyWord indicates an expected call of DeleteManyWord.
func (mr *MockIWordServiceMockRecorder) DeleteManyWord(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteManyWord", reflect.TypeOf((*MockIWordService)(nil).DeleteManyWord), arg0)
}

// DeleteOneWord mocks base method.
func (m *MockIWordService) DeleteOneWord(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOneWord", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOneWord indicates an expected call of DeleteOneWord.
func (mr *MockIWordServiceMockRecorder) DeleteOneWord(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOneWord", reflect.TypeOf((*MockIWordService)(nil).DeleteOneWord), arg0)
}

// FetchAllWords mocks base method.
func (m *MockIWordService) FetchAllWords() ([]entity.Word, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllWords")
	ret0, _ := ret[0].([]entity.Word)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAllWords indicates an expected call of FetchAllWords.
func (mr *MockIWordServiceMockRecorder) FetchAllWords() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllWords", reflect.TypeOf((*MockIWordService)(nil).FetchAllWords))
}

// UpdateManyWords mocks base method.
func (m *MockIWordService) UpdateManyWords(arg0 map[int]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateManyWords", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateManyWords indicates an expected call of UpdateManyWords.
func (mr *MockIWordServiceMockRecorder) UpdateManyWords(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateManyWords", reflect.TypeOf((*MockIWordService)(nil).UpdateManyWords), arg0)
}

// UpdateOneWord mocks base method.
func (m *MockIWordService) UpdateOneWord(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOneWord", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOneWord indicates an expected call of UpdateOneWord.
func (mr *MockIWordServiceMockRecorder) UpdateOneWord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOneWord", reflect.TypeOf((*MockIWordService)(nil).UpdateOneWord), arg0, arg1)
}
