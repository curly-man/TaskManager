package main

import (
	"TM/api"
	"TM/api/handlers"
	"TM/dbcore"
	"TM/dbclient"
)

func main(){
	dbClient := dbclient.NewDBClient()
	dbCore := dbcore.CreateTaskDBCore(dbClient)
	taskHandler := handlers.CreateTaskHandler(dbCore)
	server := api.CreateServer(":8080", taskHandler)
	server.Start()
}