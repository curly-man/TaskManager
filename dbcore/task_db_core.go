package dbcore

import (
	"TM/dbclient"
	"TM/models"
	"fmt"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	id "gopkg.in/mgo.v2/bson"
	"time"
)

type TaskDBCore interface {
	Add(task models.Task) (string, error)
	Delete(taskID id.ObjectId) (int64, error)
	Change(task models.Task) (int64, error)
	Complete(ID id.ObjectId) (int64, error)
	Get(user models.User, iscomplete bool, isfail bool) (string, error)
	GetByID(ID string) (string, error)
}

type taskDBCore struct {
	dbClient dbclient.DBClient
}

func CreateTaskDBCore(dbClient dbclient.DBClient) TaskDBCore {
	return &taskDBCore{dbClient: dbClient}
}
func (db *taskDBCore) Add(task models.Task) (string, error) {
	err := db.dbClient.NewClient()
	if err != nil {
		return "", err
	}
	err = db.dbClient.Connect()
	if err != nil {
		return "", err
	}
	defer db.dbClient.Disconnect()
	err = db.dbClient.Ping()
	if err != nil {
		return "", err
	}
	result, err := db.dbClient.InsertOne(task)
	print(result.InsertedID.(string))
	return id.ObjectId(result.InsertedID.(string)).Hex(), err
}
func (db *taskDBCore) Delete(taskID id.ObjectId) (int64, error) {
	err := db.dbClient.NewClient()
	if err != nil {
		return 0, err
	}
	err = db.dbClient.Connect()
	if err != nil {
		return 0, err
	}
	defer db.dbClient.Disconnect()
	err = db.dbClient.Ping()
	if err != nil {
		return 0, err
	}
	filter := bson.D{{"_id", taskID}}
	res, err := db.dbClient.DeleteOne(filter)
	return res.DeletedCount, err
}
func (db *taskDBCore) Change(task models.Task) (int64, error) {
	err := db.dbClient.NewClient()
	if err != nil {
		return 0, err
	}
	err = db.dbClient.Connect()
	if err != nil {
		return 0, err
	}
	defer db.dbClient.Disconnect()
	err = db.dbClient.Ping()
	if err != nil {
		return 0, err
	}
	filter := bson.D{{"_id", task.ID}}
	update := bson.D{
		{
			"$set",
			bson.D{
				{"name", task.Name},
				{"priority", task.Priority},
				{"deadline", task.Deadline},
				{"iscomplete", task.IsComplete},
				{"timing", task.Timing},
				{"isfail", task.IsFail},
			},
		},
	}
	res, err := db.dbClient.UpdateOne(filter, update)
	return res.ModifiedCount, err
}
func (db *taskDBCore) Complete(ID id.ObjectId) (int64, error) {
	err := db.dbClient.NewClient()
	if err != nil {
		return 0, err
	}
	err = db.dbClient.Connect()
	if err != nil {
		return 0, err
	}
	defer db.dbClient.Disconnect()
	err = db.dbClient.Ping()
	if err != nil {
		return 0, err
	}
	filter := bson.D{{"_id", ID}}
	update := bson.D{
		{
			"$set",
			bson.D{
				{"iscomplete", true},
				{"isfail", false},
			},
		},
	}
	res, err := db.dbClient.UpdateOne(filter, update)
	return res.ModifiedCount, err
}
func (db *taskDBCore) Get(user models.User, iscomplete, isfail bool) (string, error) {
	err := db.dbClient.NewClient()
	if err != nil {
		return "", err
	}
	err = db.dbClient.Connect()
	if err != nil {
		return "", err
	}
	defer db.dbClient.Disconnect()
	err = db.dbClient.Ping()
	if err != nil {
		return "", err
	}
	filter := bson.D{
		{"user", user},
		{"iscomplete", iscomplete},
		{"isfail", isfail},
	}
	cursor, err := db.dbClient.Find(filter)
	if err != nil {
		return "", err
	}
	var tasks []models.Task
	for cursor.Next(context.TODO()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return "", err
		}
		if task.Deadline && !task.IsComplete && !task.IsFail {
			fmt.Println("task is deadline and not completed")
			fmt.Println(task.Timing, time.Now().Unix())
			if task.Timing < time.Now().Unix() {
				task.IsFail = true
				fmt.Println("task is fail")
				db.Change(task)
				continue
			}
		}
		tasks = append(tasks, task)
	}
	if err != nil {
		return "", err
	}
	bytes, err := json.Marshal(tasks)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
func (db *taskDBCore) GetByID(ID string) (string, error) {
	err := db.dbClient.NewClient()
	if err != nil {
		return "", err
	}
	err = db.dbClient.Connect()
	if err != nil {
		return "", err
	}
	defer db.dbClient.Disconnect()
	err = db.dbClient.Ping()
	if err != nil {
		return "", err
	}
	filter := bson.D{{"_id", id.ObjectIdHex(ID)}}
	result := db.dbClient.FindOne(filter)
	var task models.Task
	err = result.Decode(&task)
	if err != nil {
		return "null", nil
	}
	bytes, err := json.Marshal(task)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
