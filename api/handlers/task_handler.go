package handlers

import (
	"TM/dbcore"
	"TM/models"
	"encoding/json"
	"time"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)

type TaskHandler interface {
	HandleOptionsRequest(w http.ResponseWriter, r *http.Request)
	CreateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	ChangeTask(w http.ResponseWriter, r *http.Request)
	CompleteTask(w http.ResponseWriter, r *http.Request)
	GetInboxTasks(w http.ResponseWriter, r *http.Request)
	GetCompleteTasks(w http.ResponseWriter, r *http.Request)
	GetFailTasks(w http.ResponseWriter, r *http.Request)
	GetTaskByID(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct {
	DB dbcore.TaskDBCore
}

func CreateTaskHandler(dbCore dbcore.TaskDBCore) TaskHandler {
	return &taskHandler{DB: dbCore}
}

func checkUser() models.User {
	user := models.User{Username: "Anton"}
	return user
}

func (th *taskHandler) HandleOptionsRequest(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w)
}

func (th *taskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	user := checkUser()
	body, _ := ioutil.ReadAll(r.Body)
	time := time.Now().Unix()
	task := models.Task{ID: bson.NewObjectId(), Timing: time, User: user}
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
func (th *taskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
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
func (th *taskHandler) ChangeTask(w http.ResponseWriter, r *http.Request) {
	// user := checkUser()
	body, _ := ioutil.ReadAll(r.Body)
	task := models.Task{}
	err := json.Unmarshal(body, &task)
	fmt.Println(task)
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
func (th *taskHandler) CompleteTask(w http.ResponseWriter, r *http.Request) {
	// user := checkUser()
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	res, err := th.DB.Complete(id)
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
func (th *taskHandler) GetInboxTasks(w http.ResponseWriter, r *http.Request) {
	user := checkUser()
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
func (th *taskHandler) GetCompleteTasks(w http.ResponseWriter, r *http.Request) {
	user := checkUser()
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
func (th *taskHandler) GetFailTasks(w http.ResponseWriter, r *http.Request) {
	user := checkUser()
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
func (th *taskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// user := checkUser()
	vars := mux.Vars(r)
	id := vars["id"]
	taskJSON, err := th.DB.GetByID(id)
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