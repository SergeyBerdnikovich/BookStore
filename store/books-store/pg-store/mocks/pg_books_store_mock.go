// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gtforge/BookStore/store/books-store/pg-store (interfaces: PgBooksStoreInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/gtforge/BookStore/models"
)

// MockPgBooksStoreInterface is a mock of PgBooksStoreInterface interface
type MockPgBooksStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPgBooksStoreInterfaceMockRecorder
}

// MockPgBooksStoreInterfaceMockRecorder is the mock recorder for MockPgBooksStoreInterface
type MockPgBooksStoreInterfaceMockRecorder struct {
	mock *MockPgBooksStoreInterface
}

// NewMockPgBooksStoreInterface creates a new mock instance
func NewMockPgBooksStoreInterface(ctrl *gomock.Controller) *MockPgBooksStoreInterface {
	mock := &MockPgBooksStoreInterface{ctrl: ctrl}
	mock.recorder = &MockPgBooksStoreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPgBooksStoreInterface) EXPECT() *MockPgBooksStoreInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockPgBooksStoreInterface) Create(arg0 string, arg1 int) error {
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockPgBooksStoreInterfaceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPgBooksStoreInterface)(nil).Create), arg0, arg1)
}

// DecreaseQuantity mocks base method
func (m *MockPgBooksStoreInterface) DecreaseQuantity(arg0 uint) error {
	ret := m.ctrl.Call(m, "DecreaseQuantity", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecreaseQuantity indicates an expected call of DecreaseQuantity
func (mr *MockPgBooksStoreInterfaceMockRecorder) DecreaseQuantity(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseQuantity", reflect.TypeOf((*MockPgBooksStoreInterface)(nil).DecreaseQuantity), arg0)
}

// Delete mocks base method
func (m *MockPgBooksStoreInterface) Delete(arg0 uint) error {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPgBooksStoreInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPgBooksStoreInterface)(nil).Delete), arg0)
}

// GetAll mocks base method
func (m *MockPgBooksStoreInterface) GetAll() ([]models.Book, error) {
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockPgBooksStoreInterfaceMockRecorder) GetAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPgBooksStoreInterface)(nil).GetAll))
}

// GetById mocks base method
func (m *MockPgBooksStoreInterface) GetById(arg0 uint) (models.Book, error) {
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById
func (mr *MockPgBooksStoreInterfaceMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockPgBooksStoreInterface)(nil).GetById), arg0)
}

// GetPerPage mocks base method
func (m *MockPgBooksStoreInterface) GetPerPage(arg0 int) ([]models.Book, error) {
	ret := m.ctrl.Call(m, "GetPerPage", arg0)
	ret0, _ := ret[0].([]models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPerPage indicates an expected call of GetPerPage
func (mr *MockPgBooksStoreInterfaceMockRecorder) GetPerPage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPerPage", reflect.TypeOf((*MockPgBooksStoreInterface)(nil).GetPerPage), arg0)
}

// Update mocks base method
func (m *MockPgBooksStoreInterface) Update(arg0 uint, arg1 string, arg2 int) error {
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockPgBooksStoreInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPgBooksStoreInterface)(nil).Update), arg0, arg1, arg2)
}
