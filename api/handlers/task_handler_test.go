package handlers

import (
	"TM/dbcore"
	"TM/dbcore/db_core_mocks"
	"TM/models"
	"TM/dbclient"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"errors"
	"github.com/gorilla/mux"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/x/mongo/driver"
	"gopkg.in/mgo.v2/bson"
)

func TestCreateTaskHandler(t *testing.T) {
	t.Parallel()
	dbClient := dbclient.NewDBClient()
	dbCore := dbcore.CreateTaskDBCore(dbClient)
	taskH := CreateTaskHandler(dbCore)

	var th = &taskHandler{DB: dbCore}

	assert.Equal(t, taskH, th)
}

func TestCreateTaskStatusOk(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	task := models.Task{ID: bson.NewObjectId(), Name: "First Task"}
	dbmock.EXPECT().Add(task).Return(task.ID.Hex(), nil)

	th := CreateTaskHandler(dbmock)
	body, _ := json.Marshal(task)

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.CreateTask(w, r)

	assert.Equal(t, w.Code, http.StatusCreated)
}

func TestCreateTaskBadRequest(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)

	th := CreateTaskHandler(dbmock)
	body := "TEST"

	r := httptest.NewRequest("POST", "/", strings.NewReader(body))
	w := httptest.NewRecorder()
	th.CreateTask(w, r)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestCreateTaskBadGateway(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	task := models.Task{ID: bson.NewObjectId(), Name: "First Task"}
	dbmock.EXPECT().Add(task).Return("", driver.ErrFilterType)

	th := CreateTaskHandler(dbmock)
	body, _ := json.Marshal(task)

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.CreateTask(w, r)

	assert.Equal(t, w.Code, http.StatusBadGateway)
}

func TestDeleteTaskStatusOk(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	id := bson.ObjectIdHex("5dee57c62166b12990c90102")
	dbmock.EXPECT().Delete(id).Return(int64(1), nil)

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("DELETE", "/5dee57c62166b12990c90102", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
        "id": "5dee57c62166b12990c90102",
    }

    r = mux.SetURLVars(r, vars)
	th.DeleteTask(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestDeleteTaskStatusBadGateway(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	id := bson.ObjectIdHex("5dee57c62166b12990c90102")
	dbmock.EXPECT().Delete(id).Return(int64(0), driver.ErrFilterType)

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("DELETE", "/5dee57c62166b12990c90102", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
        "id": "5dee57c62166b12990c90102",
    }

    r = mux.SetURLVars(r, vars)
	th.DeleteTask(w, r)

	assert.Equal(t, w.Code, http.StatusBadGateway)
}

func TestChangeTaskStatusOk(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	task := models.Task{ID: bson.NewObjectId(), Name: "First Task"}
	dbmock.EXPECT().Change(task).Return(int64(1), nil)

	th := CreateTaskHandler(dbmock)
	body, _ := json.Marshal(task)

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.ChangeTask(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestChangeTaskStatusBadRequest(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)

	th := CreateTaskHandler(dbmock)
	body := "TEST"

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.ChangeTask(w, r)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestChangeTaskStatusBadGateway(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	task := models.Task{ID: bson.NewObjectId(), Name: "First Task"}
	dbmock.EXPECT().Change(task).Return(int64(0), driver.ErrFilterType)

	th := CreateTaskHandler(dbmock)
	body, _ := json.Marshal(task)

	r := httptest.NewRequest("POST", "/", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	th.ChangeTask(w, r)

	assert.Equal(t, w.Code, http.StatusBadGateway)
}

func TestCompleteTaskStatusOk(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	id := bson.ObjectIdHex("5dee57c62166b12990c90102")
	dbmock.EXPECT().Complete(id).Return(int64(1), nil)

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("PUT", "/", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
        "id": "5dee57c62166b12990c90102",
    }

    r = mux.SetURLVars(r, vars)
	th.CompleteTask(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestCompleteTaskStatusBadGateway(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	id := bson.ObjectIdHex("5dee57c62166b12990c90102")
	dbmock.EXPECT().Complete(id).Return(int64(0), driver.ErrFilterType)

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
        "id": "5dee57c62166b12990c90102",
    }

    r = mux.SetURLVars(r, vars)
	th.CompleteTask(w, r)

	assert.Equal(t, w.Code, http.StatusBadGateway)
}

func TestGetInboxTaskStatusOk(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	user := models.User{Username: "Anton"}
	dbmock.EXPECT().Get(user, false, false).Return("null", nil)

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	th.GetInboxTasks(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestGetInboxTaskStatusNotFound(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	user := models.User{Username: "Anton"}
	dbmock.EXPECT().Get(user, false, false).Return("null", errors.New("some err"))

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	th.GetInboxTasks(w, r)

	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestGetCompleteTaskStatusOk(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	user := models.User{Username: "Anton"}
	dbmock.EXPECT().Get(user, true, false).Return("null", nil)

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	th.GetCompleteTasks(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestGetCompleteTaskStatusNotFound(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	user := models.User{Username: "Anton"}
	dbmock.EXPECT().Get(user, true, false).Return("null", errors.New("some err"))

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	th.GetCompleteTasks(w, r)

	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestGetFailTaskStatusOk(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	user := models.User{Username: "Anton"}
	dbmock.EXPECT().Get(user, false, true).Return("null", nil)

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	th.GetFailTasks(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestGetFailTaskStatusNotFound(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	user := models.User{Username: "Anton"}
	dbmock.EXPECT().Get(user, false, true).Return("null", errors.New("some err"))

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	th.GetFailTasks(w, r)

	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestGetTaskByIDStatusOk(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	id := "5dee57c62166b12990c90102"
	dbmock.EXPECT().GetByID(id).Return("null", nil)

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/5de56b8e2166b11f9ca7b333", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
        "id": "5dee57c62166b12990c90102",
    }

    r = mux.SetURLVars(r, vars)
	th.GetTaskByID(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestGetTaskByIDStatusNotFound(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := db_core_mocks.NewMockTaskDBCore(ctrl)
	id := "5dee57c62166b12990c90102"
	dbmock.EXPECT().GetByID(id).Return("", driver.ErrFilterType)

	th := CreateTaskHandler(dbmock)

	r := httptest.NewRequest("GET", "/5de56b8e2166b11f9ca7b333", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
        "id": "5dee57c62166b12990c90102",
    }

    r = mux.SetURLVars(r, vars)
	th.GetTaskByID(w, r)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
