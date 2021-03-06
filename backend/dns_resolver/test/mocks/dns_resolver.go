// Code generated by MockGen. DO NOT EDIT.
// Source: dns_client_wrapper.go

// Package mock_dns_resolver is a generated GoMock package.
package mock_dns_resolver

import (
reflect "reflect"
time "time"

gomock "github.com/golang/mock/gomock"
dns "github.com/miekg/dns"
)

// MockDnsClientWrapper is a mock of DnsClientWrapper interface.
type MockDnsClientWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockDnsClientWrapperMockRecorder
}

// MockDnsClientWrapperMockRecorder is the mock recorder for MockDnsClientWrapper.
type MockDnsClientWrapperMockRecorder struct {
	mock *MockDnsClientWrapper
}

// NewMockDnsClientWrapper creates a new mock instance.
func NewMockDnsClientWrapper(ctrl *gomock.Controller) *MockDnsClientWrapper {
	mock := &MockDnsClientWrapper{ctrl: ctrl}
	mock.recorder = &MockDnsClientWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDnsClientWrapper) EXPECT() *MockDnsClientWrapperMockRecorder {
	return m.recorder
}

// Lookup mocks base method.
func (m *MockDnsClientWrapper) Lookup(queryMsg *dns.Msg, nServer string) (time.Duration, string, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", queryMsg, nServer)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].([]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Lookup indicates an expected call of Lookup.
func (mr *MockDnsClientWrapperMockRecorder) Lookup(queryMsg, nServer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockDnsClientWrapper)(nil).Lookup), queryMsg, nServer)
}
