package api

import (
	"TM/api/handlers"
	"net/http"
	"github.com/gorilla/mux"
)

type Server struct{
	port string
	taskHandler handlers.Task_Handler
}

func CreateServer(port string, taskHandler handlers.Task_Handler) *Server{
	return &Server{
		port: port,
		taskHandler: taskHandler,
	}
}

func (s *Server) Start(){
	router := mux.NewRouter()
	router.HandleFunc("/api/task/create/", s.taskHandler.Create_Task).Methods("POST")
	router.HandleFunc("/api/task/delete/", s.taskHandler.Delete_Task).Methods("DELETE")
	router.HandleFunc("/api/task/change/", s.taskHandler.Change_Task).Methods("POST")
	router.HandleFunc("/api/task/complete/", s.taskHandler.Complete_Task).Methods("PUT")
	router.HandleFunc("/api/tasks/inbox/", s.taskHandler.Get_Inbox_Tasks).Methods("GET")
	router.HandleFunc("/api/tasks/complete/", s.taskHandler.Get_Complete_Tasks).Methods("GET")
	router.HandleFunc("/api/tasks/fail/", s.taskHandler.Get_Fail_Tasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", s.taskHandler.Get_Task_By_ID).Methods("GET")

	http.ListenAndServe(s.port, router)
}