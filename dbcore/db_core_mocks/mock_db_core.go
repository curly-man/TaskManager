package db_core_mocks

import (
        models "TM/models"
        gomock "github.com/golang/mock/gomock"
        bson "gopkg.in/mgo.v2/bson"
        reflect "reflect"
)

// MockTaskDBCore is a mock of TaskDBCore interface
type MockTaskDBCore struct {
        ctrl     *gomock.Controller
        recorder *MockTaskDBCoreMockRecorder
}

// MockTaskDBCoreMockRecorder is the mock recorder for MockTaskDBCore
type MockTaskDBCoreMockRecorder struct {
        mock *MockTaskDBCore
}

// NewMockTaskDBCore creates a new mock instance
func NewMockTaskDBCore(ctrl *gomock.Controller) *MockTaskDBCore {
        mock := &MockTaskDBCore{ctrl: ctrl}
        mock.recorder = &MockTaskDBCoreMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskDBCore) EXPECT() *MockTaskDBCoreMockRecorder {
        return m.recorder
}

// Add mocks base method
func (m *MockTaskDBCore) Add(task models.Task) (interface{}, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Add", task)
        ret0, _ := ret[0].(interface{})
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockTaskDBCoreMockRecorder) Add(task interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTaskDBCore)(nil).Add), task)
}

// Delete mocks base method
func (m *MockTaskDBCore) Delete(taskID bson.ObjectId) (interface{}, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Delete", taskID)
        ret0, _ := ret[0].(interface{})
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockTaskDBCoreMockRecorder) Delete(taskID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTaskDBCore)(nil).Delete), taskID)
}

// Change mocks base method
func (m *MockTaskDBCore) Change(task models.Task) (interface{}, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Change", task)
        ret0, _ := ret[0].(interface{})
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Change indicates an expected call of Change
func (mr *MockTaskDBCoreMockRecorder) Change(task interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Change", reflect.TypeOf((*MockTaskDBCore)(nil).Change), task)
}

// Get mocks base method
func (m *MockTaskDBCore) Get(user models.User, iscomplete, isfail bool) (string, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Get", user, iscomplete, isfail)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTaskDBCoreMockRecorder) Get(user, iscomplete, isfail interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTaskDBCore)(nil).Get), user, iscomplete, isfail)
}

// Get_By_ID mocks base method
func (m *MockTaskDBCore) Get_By_ID(ID string) (string, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Get_By_ID", ID)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Get_By_ID indicates an expected call of Get_By_ID
func (mr *MockTaskDBCoreMockRecorder) Get_By_ID(ID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get_By_ID", reflect.TypeOf((*MockTaskDBCore)(nil).Get_By_ID), ID)
}