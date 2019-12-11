package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	ID         bson.ObjectId `bson:"_id" json:"_id,omitempty"`
	Name       string
	Priority   int
	Timing     int64
	Subtask    []Task
	User       User
	IsComplete bool
	IsFail     bool
}