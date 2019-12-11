package mock_handlers

import (
        gomock "github.com/golang/mock/gomock"
        http "net/http"
        reflect "reflect"
)

// MockTask_Handler is a mock of Task_Handler interface
type MockTask_Handler struct {
        ctrl     *gomock.Controller
        recorder *MockTask_HandlerMockRecorder
}

// MockTask_HandlerMockRecorder is the mock recorder for MockTask_Handler
type MockTask_HandlerMockRecorder struct {
        mock *MockTask_Handler
}

// NewMockTask_Handler creates a new mock instance
func NewMockTask_Handler(ctrl *gomock.Controller) *MockTask_Handler {
        mock := &MockTask_Handler{ctrl: ctrl}
        mock.recorder = &MockTask_HandlerMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTask_Handler) EXPECT() *MockTask_HandlerMockRecorder {
        return m.recorder
}

// Create_Task mocks base method
func (m *MockTask_Handler) Create_Task(w http.ResponseWriter, r *http.Request) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "Create_Task", w, r)
}

// Create_Task indicates an expected call of Create_Task
func (mr *MockTask_HandlerMockRecorder) Create_Task(w, r interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create_Task", reflect.TypeOf((*MockTask_Handler)(nil).Create_Task), w, r)
}

// Delete_Task mocks base method
func (m *MockTask_Handler) Delete_Task(w http.ResponseWriter, r *http.Request) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "Delete_Task", w, r)
}

// Delete_Task indicates an expected call of Delete_Task
func (mr *MockTask_HandlerMockRecorder) Delete_Task(w, r interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete_Task", reflect.TypeOf((*MockTask_Handler)(nil).Delete_Task), w, r)
}

// Change_Task mocks base method
func (m *MockTask_Handler) Change_Task(w http.ResponseWriter, r *http.Request) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "Change_Task", w, r)
}

// Change_Task indicates an expected call of Change_Task
func (mr *MockTask_HandlerMockRecorder) Change_Task(w, r interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Change_Task", reflect.TypeOf((*MockTask_Handler)(nil).Change_Task), w, r)
}

// Complete_Task mocks base method
func (m *MockTask_Handler) Complete_Task(w http.ResponseWriter, r *http.Request) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "Complete_Task", w, r)
}

// Complete_Task indicates an expected call of Complete_Task
func (mr *MockTask_HandlerMockRecorder) Complete_Task(w, r interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Complete_Task", reflect.TypeOf((*MockTask_Handler)(nil).Complete_Task), w, r)
}

// Get_Inbox_Tasks mocks base method
func (m *MockTask_Handler) Get_Inbox_Tasks(w http.ResponseWriter, r *http.Request) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "Get_Inbox_Tasks", w, r)
}

// Get_Inbox_Tasks indicates an expected call of Get_Inbox_Tasks
func (mr *MockTask_HandlerMockRecorder) Get_Inbox_Tasks(w, r interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get_Inbox_Tasks", reflect.TypeOf((*MockTask_Handler)(nil).Get_Inbox_Tasks), w, r)
}

// Get_Complete_Tasks mocks base method
func (m *MockTask_Handler) Get_Complete_Tasks(w http.ResponseWriter, r *http.Request) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "Get_Complete_Tasks", w, r)
}

// Get_Complete_Tasks indicates an expected call of Get_Complete_Tasks
func (mr *MockTask_HandlerMockRecorder) Get_Complete_Tasks(w, r interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get_Complete_Tasks", reflect.TypeOf((*MockTask_Handler)(nil).Get_Complete_Tasks), w, r)
}

// Get_Fail_Tasks mocks base method
func (m *MockTask_Handler) Get_Fail_Tasks(w http.ResponseWriter, r *http.Request) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "Get_Fail_Tasks", w, r)
}

// Get_Fail_Tasks indicates an expected call of Get_Fail_Tasks
func (mr *MockTask_HandlerMockRecorder) Get_Fail_Tasks(w, r interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get_Fail_Tasks", reflect.TypeOf((*MockTask_Handler)(nil).Get_Fail_Tasks), w, r)
}

// Get_Task_By_ID mocks base method
func (m *MockTask_Handler) Get_Task_By_ID(w http.ResponseWriter, r *http.Request) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "Get_Task_By_ID", w, r)
}

// Get_Task_By_ID indicates an expected call of Get_Task_By_ID
func (mr *MockTask_HandlerMockRecorder) Get_Task_By_ID(w, r interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get_Task_By_ID", reflect.TypeOf((*MockTask_Handler)(nil).Get_Task_By_ID), w, r)
}