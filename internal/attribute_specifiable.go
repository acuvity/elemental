// Code generated by MockGen. DO NOT EDIT.
// Source: go.acuvity.ai/elemental (interfaces: AttributeSpecifiable)

// Package internal is a generated GoMock package.
package internal

import (
	gomock "github.com/golang/mock/gomock"
	elemental "go.acuvity.ai/elemental"
	reflect "reflect"
)

// MockAttributeSpecifiable is a mock of AttributeSpecifiable interface
type MockAttributeSpecifiable struct {
	ctrl     *gomock.Controller
	recorder *MockAttributeSpecifiableMockRecorder
}

// MockAttributeSpecifiableMockRecorder is the mock recorder for MockAttributeSpecifiable
type MockAttributeSpecifiableMockRecorder struct {
	mock *MockAttributeSpecifiable
}

// NewMockAttributeSpecifiable creates a new mock instance
func NewMockAttributeSpecifiable(ctrl *gomock.Controller) *MockAttributeSpecifiable {
	mock := &MockAttributeSpecifiable{ctrl: ctrl}
	mock.recorder = &MockAttributeSpecifiableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAttributeSpecifiable) EXPECT() *MockAttributeSpecifiableMockRecorder {
	return m.recorder
}

// AttributeSpecifications mocks base method
func (m *MockAttributeSpecifiable) AttributeSpecifications() map[string]elemental.AttributeSpecification {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttributeSpecifications")
	ret0, _ := ret[0].(map[string]elemental.AttributeSpecification)
	return ret0
}

// AttributeSpecifications indicates an expected call of AttributeSpecifications
func (mr *MockAttributeSpecifiableMockRecorder) AttributeSpecifications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttributeSpecifications", reflect.TypeOf((*MockAttributeSpecifiable)(nil).AttributeSpecifications))
}

// SpecificationForAttribute mocks base method
func (m *MockAttributeSpecifiable) SpecificationForAttribute(arg0 string) elemental.AttributeSpecification {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpecificationForAttribute", arg0)
	ret0, _ := ret[0].(elemental.AttributeSpecification)
	return ret0
}

// SpecificationForAttribute indicates an expected call of SpecificationForAttribute
func (mr *MockAttributeSpecifiableMockRecorder) SpecificationForAttribute(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpecificationForAttribute", reflect.TypeOf((*MockAttributeSpecifiable)(nil).SpecificationForAttribute), arg0)
}

// ValueForAttribute mocks base method
func (m *MockAttributeSpecifiable) ValueForAttribute(arg0 string) any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValueForAttribute", arg0)
	ret0, _ := ret[0].(any)
	return ret0
}

// ValueForAttribute indicates an expected call of ValueForAttribute
func (mr *MockAttributeSpecifiableMockRecorder) ValueForAttribute(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValueForAttribute", reflect.TypeOf((*MockAttributeSpecifiable)(nil).ValueForAttribute), arg0)
}
