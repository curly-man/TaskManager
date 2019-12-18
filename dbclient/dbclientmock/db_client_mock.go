package dbclientmock

import (
        models "TM/models"
        gomock "github.com/golang/mock/gomock"
        bson "go.mongodb.org/mongo-driver/bson"
        mongo "go.mongodb.org/mongo-driver/mongo"
        reflect "reflect"
)

// MockDBClient is a mock of DBClient interface
type MockDBClient struct {
        ctrl     *gomock.Controller
        recorder *MockDBClientMockRecorder
}

// MockDBClientMockRecorder is the mock recorder for MockDBClient
type MockDBClientMockRecorder struct {
        mock *MockDBClient
}

// NewMockDBClient creates a new mock instance
func NewMockDBClient(ctrl *gomock.Controller) *MockDBClient {
        mock := &MockDBClient{ctrl: ctrl}
        mock.recorder = &MockDBClientMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDBClient) EXPECT() *MockDBClientMockRecorder {
        return m.recorder
}

// NewClient mocks base method
func (m *MockDBClient) NewClient() error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "NewClient")
        ret0, _ := ret[0].(error)
        return ret0
}

// NewClient indicates an expected call of NewClient
func (mr *MockDBClientMockRecorder) NewClient() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClient", reflect.TypeOf((*MockDBClient)(nil).NewClient))
}

// Connect mocks base method
func (m *MockDBClient) Connect() error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Connect")
        ret0, _ := ret[0].(error)
        return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockDBClientMockRecorder) Connect() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockDBClient)(nil).Connect))
}

// Disconnect mocks base method
func (m *MockDBClient) Disconnect() error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Disconnect")
        ret0, _ := ret[0].(error)
        return ret0
}

// Disconnect indicates an expected call of Disconnect
func (mr *MockDBClientMockRecorder) Disconnect() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockDBClient)(nil).Disconnect))
}

// Ping mocks base method
func (m *MockDBClient) Ping() error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Ping")
        ret0, _ := ret[0].(error)
        return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockDBClientMockRecorder) Ping() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDBClient)(nil).Ping))
}

// InsertOne mocks base method
func (m *MockDBClient) InsertOne(task models.Task) (*mongo.InsertOneResult, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "InsertOne", task)
        ret0, _ := ret[0].(*mongo.InsertOneResult)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne
func (mr *MockDBClientMockRecorder) InsertOne(task interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockDBClient)(nil).InsertOne), task)
}

// Find mocks base method
func (m *MockDBClient) Find(filter bson.D) (*mongo.Cursor, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Find", filter)
        ret0, _ := ret[0].(*mongo.Cursor)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockDBClientMockRecorder) Find(filter bson.D) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockDBClient)(nil).Find), filter)
}

// FindOne mocks base method
func (m *MockDBClient) FindOne(filter bson.D) *mongo.SingleResult {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "FindOne", filter)
        ret0, _ := ret[0].(*mongo.SingleResult)
        return ret0
}

// FindOne indicates an expected call of FindOne
func (mr *MockDBClientMockRecorder) FindOne(filter interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockDBClient)(nil).FindOne), filter)
}

// DeleteOne mocks base method
func (m *MockDBClient) DeleteOne(filter bson.D) (*mongo.DeleteResult, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteOne", filter)
        ret0, _ := ret[0].(*mongo.DeleteResult)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// DeleteOne indicates an expected call of DeleteOne
func (mr *MockDBClientMockRecorder) DeleteOne(filter bson.D) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockDBClient)(nil).DeleteOne), filter)
}

// UpdateOne mocks base method
func (m *MockDBClient) UpdateOne(filter, update bson.D) (*mongo.UpdateResult, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "UpdateOne", filter, update)
        ret0, _ := ret[0].(*mongo.UpdateResult)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// UpdateOne indicates an expected call of UpdateOne
func (mr *MockDBClientMockRecorder) UpdateOne(filter, update bson.D) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockDBClient)(nil).UpdateOne), filter, update)
}