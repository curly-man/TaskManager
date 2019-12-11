package main

import (
	"TM/api"
	"TM/api/handlers"
	"TM/config"
	"TM/dbcore"
)

func main(){
	dbCore := dbcore.CreateTaskDBCore(config.MongoURL, config.MongoDB)
	taskHandler := handlers.CreateTaskHandler(dbCore)
	server := api.CreateServer(":8080", taskHandler)
	server.Start()
}