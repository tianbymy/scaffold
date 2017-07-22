// Code generated by MockGen. DO NOT EDIT.
// Source: fs/fs.go

package fs

import (
	gomock "github.com/golang/mock/gomock"
)

// MockFS is a mock of FS interface
type MockFS struct {
	ctrl     *gomock.Controller
	recorder *MockFSMockRecorder
}

// MockFSMockRecorder is the mock recorder for MockFS
type MockFSMockRecorder struct {
	mock *MockFS
}

// NewMockFS creates a new mock instance
func NewMockFS(ctrl *gomock.Controller) *MockFS {
	mock := &MockFS{ctrl: ctrl}
	mock.recorder = &MockFSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockFS) EXPECT() *MockFSMockRecorder {
	return _m.recorder
}

// GetDirs mocks base method
func (_m *MockFS) GetDirs(path string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetDirs", path)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDirs indicates an expected call of GetDirs
func (_mr *MockFSMockRecorder) GetDirs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDirs", arg0)
}

// ReadFile mocks base method
func (_m *MockFS) ReadFile(path string) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "ReadFile", path)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile
func (_mr *MockFSMockRecorder) ReadFile(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadFile", arg0)
}

// Walk mocks base method
func (_m *MockFS) Walk(path string, cb func(string, bool, error) error) error {
	ret := _m.ctrl.Call(_m, "Walk", path, cb)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk
func (_mr *MockFSMockRecorder) Walk(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Walk", arg0, arg1)
}

// CreateDir mocks base method
func (_m *MockFS) CreateDir(path string) error {
	ret := _m.ctrl.Call(_m, "CreateDir", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDir indicates an expected call of CreateDir
func (_mr *MockFSMockRecorder) CreateDir(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateDir", arg0)
}

// CreateFile mocks base method
func (_m *MockFS) CreateFile(path string, content string) error {
	ret := _m.ctrl.Call(_m, "CreateFile", path, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFile indicates an expected call of CreateFile
func (_mr *MockFSMockRecorder) CreateFile(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateFile", arg0, arg1)
}

// Exists mocks base method
func (_m *MockFS) Exists(path string) (bool, error) {
	ret := _m.ctrl.Call(_m, "Exists", path)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (_mr *MockFSMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Exists", arg0)
}

// DirExists mocks base method
func (_m *MockFS) DirExists(path string) (bool, error) {
	ret := _m.ctrl.Call(_m, "DirExists", path)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DirExists indicates an expected call of DirExists
func (_mr *MockFSMockRecorder) DirExists(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DirExists", arg0)
}
