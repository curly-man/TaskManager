package api

import (
	"TM/api/handlers"
	"net/http"
	"github.com/gorilla/mux"
)

type Server struct{
	port string
	taskHandler handlers.TaskHandler
}

func CreateServer(port string, taskHandler handlers.TaskHandler) *Server{
	return &Server{
		port: port,
		taskHandler: taskHandler,
	}
}

func (s *Server) Start(){
	router := mux.NewRouter()
	router.HandleFunc("/api/task/", s.taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/", s.taskHandler.ChangeTask).Methods("PUT")
	router.HandleFunc("/api/task/", s.taskHandler.HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/api/task/{id}/", s.taskHandler.HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/api/task/{id}/", s.taskHandler.DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/task/{id}/", s.taskHandler.CompleteTask).Methods("POST")
	router.HandleFunc("/api/task/{id}/", s.taskHandler.GetTaskByID).Methods("GET")
	router.HandleFunc("/api/tasks/inbox/", s.taskHandler.GetInboxTasks).Methods("GET")
	router.HandleFunc("/api/tasks/complete/", s.taskHandler.GetCompleteTasks).Methods("GET")
	router.HandleFunc("/api/tasks/fail/", s.taskHandler.GetFailTasks).Methods("GET")

	http.ListenAndServe(s.port, router)
}