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
func (m *MockTaskDBCore) Add(task models.Task) (string, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Add", task)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockTaskDBCoreMockRecorder) Add(task interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTaskDBCore)(nil).Add), task)
}

// Delete mocks base method
func (m *MockTaskDBCore) Delete(taskID bson.ObjectId) (int64, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Delete", taskID)
        ret0, _ := ret[0].(int64)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockTaskDBCoreMockRecorder) Delete(taskID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTaskDBCore)(nil).Delete), taskID)
}

// Change mocks base method
func (m *MockTaskDBCore) Change(task models.Task) (int64, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Change", task)
        ret0, _ := ret[0].(int64)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Change indicates an expected call of Change
func (mr *MockTaskDBCoreMockRecorder) Change(task interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Change", reflect.TypeOf((*MockTaskDBCore)(nil).Change), task)
}

// Complete mocks base method
func (m *MockTaskDBCore) Complete(ID bson.ObjectId) (int64, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Complete", ID)
        ret0, _ := ret[0].(int64)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Complete indicates an expected call of Complete
func (mr *MockTaskDBCoreMockRecorder) Complete(ID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Complete", reflect.TypeOf((*MockTaskDBCore)(nil).Complete), ID)
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

// GetByID mocks base method
func (m *MockTaskDBCore) GetByID(ID string) (string, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetByID", ID)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockTaskDBCoreMockRecorder) GetByID(ID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTaskDBCore)(nil).GetByID), ID)
}