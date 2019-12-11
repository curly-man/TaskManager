package handlers

import (
	"TM/config"
	"TM/dbcore"
	"TM/dbcore/db_core_mocks"
	"TM/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/x/mongo/driver"
	"gopkg.in/mgo.v2/bson"
)

func TestCreateTaskHandler(t *testing.T) {
	dbCore := dbcore.CreateTaskDBCore(config.MongoURL, config.MongoDB)
	taskH := CreateTaskHandler(dbCore)

	var th = &taskHandler{DB: dbCore}

	assert.Equal(t, taskH, th)
}

func TestCreateTaskStatusOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	task := models.Task{ID: bson.NewObjectId(), Name: "First Task"}
	dbmock.EXPECT().Add(task).Return(task.ID, nil)

	th := CreateTaskHandler(dbmock)
	body, _ := json.Marshal(task)

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.Create_Task(w, r)

	assert.Equal(t, w.Code, http.StatusCreated)
}

func TestCreateTaskBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)

	th := CreateTaskHandler(dbmock)
	body := "TEST"

	r := httptest.NewRequest("POST", "/", strings.NewReader(body))
	w := httptest.NewRecorder()
	th.Create_Task(w, r)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestCreateTaskBadGateway(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	task := models.Task{ID: bson.NewObjectId(), Name: "First Task"}
	dbmock.EXPECT().Add(task).Return("", driver.ErrFilterType)

	th := CreateTaskHandler(dbmock)
	body, _ := json.Marshal(task)

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.Create_Task(w, r)

	assert.Equal(t, w.Code, http.StatusBadGateway)
}

func TestDeleteTaskStatusOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	id := bson.NewObjectId()
	dbmock.EXPECT().Delete(id).Return(id, nil)

	th := CreateTaskHandler(dbmock)
	body, _ := json.Marshal(id)

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.Delete_Task(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestDeleteTaskStatusBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)

	th := CreateTaskHandler(dbmock)
	body := "TEST"

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.Delete_Task(w, r)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestDeleteTaskStatusBadGateway(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	id := bson.NewObjectId()
	dbmock.EXPECT().Delete(id).Return("", driver.ErrFilterType)

	th := CreateTaskHandler(dbmock)
	body, _ := json.Marshal(id)

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.Delete_Task(w, r)

	assert.Equal(t, w.Code, http.StatusBadGateway)
}

func TestChangeTaskStatusOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	task := models.Task{ID: bson.NewObjectId(), Name: "First Task"}
	dbmock.EXPECT().Change(task).Return(task.ID, nil)

	th := CreateTaskHandler(dbmock)
	body, _ := json.Marshal(task)

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.Change_Task(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestChangeTaskStatusBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)

	th := CreateTaskHandler(dbmock)
	body := "TEST"

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.Change_Task(w, r)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestChangeTaskStatusBadGateway(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	task := models.Task{ID: bson.NewObjectId(), Name: "First Task"}
	dbmock.EXPECT().Change(task).Return("", driver.ErrFilterType)

	th := CreateTaskHandler(dbmock)
	body, _ := json.Marshal(task)

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.Change_Task(w, r)

	assert.Equal(t, w.Code, http.StatusBadGateway)
}

func TestCompleteTaskStatusOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	task := models.Task{ID: bson.NewObjectId(), Name: "First Task", IsComplete: true}
	dbmock.EXPECT().Change(task).Return(task.ID, nil)

	th := CreateTaskHandler(dbmock)
	body, _ := json.Marshal(task)

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.Complete_Task(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestCompleteTaskStatusBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)

	th := CreateTaskHandler(dbmock)
	body := "TEST"

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.Complete_Task(w, r)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestCompleteTaskStatusBadGateway(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	task := models.Task{ID: bson.NewObjectId(), Name: "First Task", IsComplete: true}
	dbmock.EXPECT().Change(task).Return("", driver.ErrFilterType)

	th := CreateTaskHandler(dbmock)
	body, _ := json.Marshal(task)

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.Complete_Task(w, r)

	assert.Equal(t, w.Code, http.StatusBadGateway)
}

func TestGetInboxTaskStatusOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	user := models.User{Username: "Anton"}
	dbmock.EXPECT().Get(user, false, false).Return("null", nil)

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	th.Get_Inbox_Tasks(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestGetInboxTaskStatusNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	user := models.User{Username: "Anton"}
	dbmock.EXPECT().Get(user, false, false).Return("null", errors.New("some err"))

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	th.Get_Inbox_Tasks(w, r)

	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestGetCompleteTaskStatusOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	user := models.User{Username: "Anton"}
	dbmock.EXPECT().Get(user, true, false).Return("null", nil)

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	th.Get_Complete_Tasks(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestGetCompleteTaskStatusNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	user := models.User{Username: "Anton"}
	dbmock.EXPECT().Get(user, true, false).Return("null", errors.New("some err"))

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	th.Get_Complete_Tasks(w, r)

	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestGetFailTaskStatusOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	user := models.User{Username: "Anton"}
	dbmock.EXPECT().Get(user, false, true).Return("null", nil)

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	th.Get_Fail_Tasks(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestGetFailTaskStatusNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	user := models.User{Username: "Anton"}
	dbmock.EXPECT().Get(user, false, true).Return("null", errors.New("some err"))

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	th.Get_Fail_Tasks(w, r)

	assert.Equal(t, w.Code, http.StatusNotFound)
}

// func TestGetTaskByIDStatusOk(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
// 	ID := "5de56b8e2166b11f9ca7b333"
// 	dbmock.EXPECT().Get_By_ID(ID).Return("null", nil)

// 	th := CreateTaskHandler(dbmock)

// 	r := httptest.NewRequest("GET", "http://127.0.0.1:80/api/tasks/5de56b8e2166b11f9ca7b333", strings.NewReader(""))
// 	w := httptest.NewRecorder()
// 	th.Get_Task_By_ID(w, r)

// 	assert.Equal(t, w.Code, http.StatusOK)
// }
