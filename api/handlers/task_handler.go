package handlers

import (
	"TM/dbcore"
	"TM/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)

type Task_Handler interface {
	Create_Task(w http.ResponseWriter, r *http.Request)
	Delete_Task(w http.ResponseWriter, r *http.Request)
	Change_Task(w http.ResponseWriter, r *http.Request)
	Complete_Task(w http.ResponseWriter, r *http.Request)
	Get_Inbox_Tasks(w http.ResponseWriter, r *http.Request)
	Get_Complete_Tasks(w http.ResponseWriter, r *http.Request)
	Get_Fail_Tasks(w http.ResponseWriter, r *http.Request)
	Get_Task_By_ID(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct {
	DB dbcore.TaskDBCore
}

func CreateTaskHandler(dbCore dbcore.TaskDBCore) Task_Handler {
	return &taskHandler{DB: dbCore}
}

func check_User() models.User {
	user := models.User{Username: "Anton"}
	return user
}

func (th *taskHandler) Create_Task(w http.ResponseWriter, r *http.Request) {
	user := check_User()
	body, _ := ioutil.ReadAll(r.Body)
	task := models.Task{ID: bson.NewObjectId(), User: user}
	err := json.Unmarshal(body, &task)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w)
		return
	}
	fmt.Println(task)
	res, err := th.DB.Add(task)
	if err != nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, res)
}
func (th *taskHandler) Delete_Task(w http.ResponseWriter, r *http.Request) {
	// user := check_User()
	body, _ := ioutil.ReadAll(r.Body)
	var id bson.ObjectId
	err := json.Unmarshal(body, &id)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w)
		return
	}
	res, err := th.DB.Delete(id)
	if err != nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, res)
}
func (th *taskHandler) Change_Task(w http.ResponseWriter, r *http.Request) {
	// user := check_User()
	body, _ := ioutil.ReadAll(r.Body)
	task := models.Task{}
	err := json.Unmarshal(body, &task)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w)
		return
	}
	res, err := th.DB.Change(task)
	if err != nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, res)
}
func (th *taskHandler) Complete_Task(w http.ResponseWriter, r *http.Request) {
	// user := check_User()
	body, _ := ioutil.ReadAll(r.Body)
	task := models.Task{}
	err := json.Unmarshal(body, &task)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w)
		return
	}
	task.IsComplete = true
	res, err := th.DB.Change(task)
	if err != nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, res)
}
func (th *taskHandler) Get_Inbox_Tasks(w http.ResponseWriter, r *http.Request) {
	user := check_User()
	tasksJSON, err := th.DB.Get(user, false, false)
	if err != nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tasksJSON)
}
func (th *taskHandler) Get_Complete_Tasks(w http.ResponseWriter, r *http.Request) {
	user := check_User()
	tasksJSON, err := th.DB.Get(user, true, false)
	if err != nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tasksJSON)
}
func (th *taskHandler) Get_Fail_Tasks(w http.ResponseWriter, r *http.Request) {
	user := check_User()
	tasksJSON, err := th.DB.Get(user, false, true)
	if err != nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tasksJSON)
}
func (th *taskHandler) Get_Task_By_ID(w http.ResponseWriter, r *http.Request) {
	// user := check_User()
	vars := mux.Vars(r)
	id := vars["id"]
	taskJSON, err := th.DB.Get_By_ID(id)
	if err != nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, taskJSON)
}