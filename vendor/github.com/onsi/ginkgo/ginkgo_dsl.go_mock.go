// Code generated by MockGen. DO NOT EDIT.
// Source: vendor/github.com/onsi/ginkgo/ginkgo_dsl.go

// Package mock_ginkgo is a generated GoMock package.
package mock_ginkgo

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockGinkgoTestingT is a mock of GinkgoTestingT interface.
type MockGinkgoTestingT struct {
	ctrl     *gomock.Controller
	recorder *MockGinkgoTestingTMockRecorder
}

// MockGinkgoTestingTMockRecorder is the mock recorder for MockGinkgoTestingT.
type MockGinkgoTestingTMockRecorder struct {
	mock *MockGinkgoTestingT
}

// NewMockGinkgoTestingT creates a new mock instance.
func NewMockGinkgoTestingT(ctrl *gomock.Controller) *MockGinkgoTestingT {
	mock := &MockGinkgoTestingT{ctrl: ctrl}
	mock.recorder = &MockGinkgoTestingTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGinkgoTestingT) EXPECT() *MockGinkgoTestingTMockRecorder {
	return m.recorder
}

// Fail mocks base method.
func (m *MockGinkgoTestingT) Fail() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Fail")
}

// Fail indicates an expected call of Fail.
func (mr *MockGinkgoTestingTMockRecorder) Fail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fail", reflect.TypeOf((*MockGinkgoTestingT)(nil).Fail))
}

// MockGinkgoTInterface is a mock of GinkgoTInterface interface.
type MockGinkgoTInterface struct {
	ctrl     *gomock.Controller
	recorder *MockGinkgoTInterfaceMockRecorder
}

// MockGinkgoTInterfaceMockRecorder is the mock recorder for MockGinkgoTInterface.
type MockGinkgoTInterfaceMockRecorder struct {
	mock *MockGinkgoTInterface
}

// NewMockGinkgoTInterface creates a new mock instance.
func NewMockGinkgoTInterface(ctrl *gomock.Controller) *MockGinkgoTInterface {
	mock := &MockGinkgoTInterface{ctrl: ctrl}
	mock.recorder = &MockGinkgoTInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGinkgoTInterface) EXPECT() *MockGinkgoTInterfaceMockRecorder {
	return m.recorder
}

// Cleanup mocks base method.
func (m *MockGinkgoTInterface) Cleanup(arg0 func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Cleanup", arg0)
}

// Cleanup indicates an expected call of Cleanup.
func (mr *MockGinkgoTInterfaceMockRecorder) Cleanup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cleanup", reflect.TypeOf((*MockGinkgoTInterface)(nil).Cleanup), arg0)
}

// Error mocks base method.
func (m *MockGinkgoTInterface) Error(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockGinkgoTInterfaceMockRecorder) Error(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockGinkgoTInterface)(nil).Error), args...)
}

// Errorf mocks base method.
func (m *MockGinkgoTInterface) Errorf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockGinkgoTInterfaceMockRecorder) Errorf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockGinkgoTInterface)(nil).Errorf), varargs...)
}

// Fail mocks base method.
func (m *MockGinkgoTInterface) Fail() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Fail")
}

// Fail indicates an expected call of Fail.
func (mr *MockGinkgoTInterfaceMockRecorder) Fail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fail", reflect.TypeOf((*MockGinkgoTInterface)(nil).Fail))
}

// FailNow mocks base method.
func (m *MockGinkgoTInterface) FailNow() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FailNow")
}

// FailNow indicates an expected call of FailNow.
func (mr *MockGinkgoTInterfaceMockRecorder) FailNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailNow", reflect.TypeOf((*MockGinkgoTInterface)(nil).FailNow))
}

// Failed mocks base method.
func (m *MockGinkgoTInterface) Failed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Failed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Failed indicates an expected call of Failed.
func (mr *MockGinkgoTInterfaceMockRecorder) Failed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Failed", reflect.TypeOf((*MockGinkgoTInterface)(nil).Failed))
}

// Fatal mocks base method.
func (m *MockGinkgoTInterface) Fatal(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockGinkgoTInterfaceMockRecorder) Fatal(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockGinkgoTInterface)(nil).Fatal), args...)
}

// Fatalf mocks base method.
func (m *MockGinkgoTInterface) Fatalf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MockGinkgoTInterfaceMockRecorder) Fatalf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockGinkgoTInterface)(nil).Fatalf), varargs...)
}

// Helper mocks base method.
func (m *MockGinkgoTInterface) Helper() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Helper")
}

// Helper indicates an expected call of Helper.
func (mr *MockGinkgoTInterfaceMockRecorder) Helper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Helper", reflect.TypeOf((*MockGinkgoTInterface)(nil).Helper))
}

// Log mocks base method.
func (m *MockGinkgoTInterface) Log(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Log", varargs...)
}

// Log indicates an expected call of Log.
func (mr *MockGinkgoTInterfaceMockRecorder) Log(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockGinkgoTInterface)(nil).Log), args...)
}

// Logf mocks base method.
func (m *MockGinkgoTInterface) Logf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Logf", varargs...)
}

// Logf indicates an expected call of Logf.
func (mr *MockGinkgoTInterfaceMockRecorder) Logf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logf", reflect.TypeOf((*MockGinkgoTInterface)(nil).Logf), varargs...)
}

// Name mocks base method.
func (m *MockGinkgoTInterface) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockGinkgoTInterfaceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockGinkgoTInterface)(nil).Name))
}

// Parallel mocks base method.
func (m *MockGinkgoTInterface) Parallel() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Parallel")
}

// Parallel indicates an expected call of Parallel.
func (mr *MockGinkgoTInterfaceMockRecorder) Parallel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parallel", reflect.TypeOf((*MockGinkgoTInterface)(nil).Parallel))
}

// Setenv mocks base method.
func (m *MockGinkgoTInterface) Setenv(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Setenv", key, value)
}

// Setenv indicates an expected call of Setenv.
func (mr *MockGinkgoTInterfaceMockRecorder) Setenv(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setenv", reflect.TypeOf((*MockGinkgoTInterface)(nil).Setenv), key, value)
}

// Skip mocks base method.
func (m *MockGinkgoTInterface) Skip(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Skip", varargs...)
}

// Skip indicates an expected call of Skip.
func (mr *MockGinkgoTInterfaceMockRecorder) Skip(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skip", reflect.TypeOf((*MockGinkgoTInterface)(nil).Skip), args...)
}

// SkipNow mocks base method.
func (m *MockGinkgoTInterface) SkipNow() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SkipNow")
}

// SkipNow indicates an expected call of SkipNow.
func (mr *MockGinkgoTInterfaceMockRecorder) SkipNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SkipNow", reflect.TypeOf((*MockGinkgoTInterface)(nil).SkipNow))
}

// Skipf mocks base method.
func (m *MockGinkgoTInterface) Skipf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Skipf", varargs...)
}

// Skipf indicates an expected call of Skipf.
func (mr *MockGinkgoTInterfaceMockRecorder) Skipf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skipf", reflect.TypeOf((*MockGinkgoTInterface)(nil).Skipf), varargs...)
}

// Skipped mocks base method.
func (m *MockGinkgoTInterface) Skipped() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Skipped")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Skipped indicates an expected call of Skipped.
func (mr *MockGinkgoTInterfaceMockRecorder) Skipped() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skipped", reflect.TypeOf((*MockGinkgoTInterface)(nil).Skipped))
}

// TempDir mocks base method.
func (m *MockGinkgoTInterface) TempDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TempDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// TempDir indicates an expected call of TempDir.
func (mr *MockGinkgoTInterfaceMockRecorder) TempDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TempDir", reflect.TypeOf((*MockGinkgoTInterface)(nil).TempDir))
}

// MockBenchmarker is a mock of Benchmarker interface.
type MockBenchmarker struct {
	ctrl     *gomock.Controller
	recorder *MockBenchmarkerMockRecorder
}

// MockBenchmarkerMockRecorder is the mock recorder for MockBenchmarker.
type MockBenchmarkerMockRecorder struct {
	mock *MockBenchmarker
}

// NewMockBenchmarker creates a new mock instance.
func NewMockBenchmarker(ctrl *gomock.Controller) *MockBenchmarker {
	mock := &MockBenchmarker{ctrl: ctrl}
	mock.recorder = &MockBenchmarkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBenchmarker) EXPECT() *MockBenchmarkerMockRecorder {
	return m.recorder
}

// RecordValue mocks base method.
func (m *MockBenchmarker) RecordValue(name string, value float64, info ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, value}
	for _, a := range info {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "RecordValue", varargs...)
}

// RecordValue indicates an expected call of RecordValue.
func (mr *MockBenchmarkerMockRecorder) RecordValue(name, value interface{}, info ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, value}, info...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordValue", reflect.TypeOf((*MockBenchmarker)(nil).RecordValue), varargs...)
}

// RecordValueWithPrecision mocks base method.
func (m *MockBenchmarker) RecordValueWithPrecision(name string, value float64, units string, precision int, info ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, value, units, precision}
	for _, a := range info {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "RecordValueWithPrecision", varargs...)
}

// RecordValueWithPrecision indicates an expected call of RecordValueWithPrecision.
func (mr *MockBenchmarkerMockRecorder) RecordValueWithPrecision(name, value, units, precision interface{}, info ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, value, units, precision}, info...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordValueWithPrecision", reflect.TypeOf((*MockBenchmarker)(nil).RecordValueWithPrecision), varargs...)
}

// Time mocks base method.
func (m *MockBenchmarker) Time(name string, body func(), info ...interface{}) time.Duration {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, body}
	for _, a := range info {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Time", varargs...)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Time indicates an expected call of Time.
func (mr *MockBenchmarkerMockRecorder) Time(name, body interface{}, info ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, body}, info...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time", reflect.TypeOf((*MockBenchmarker)(nil).Time), varargs...)
}
