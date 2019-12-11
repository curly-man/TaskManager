package dbcore

import (
	"github.com/stretchr/testify/assert"
	"TM/config"
	"TM/models"
	"testing"
	"gopkg.in/mgo.v2/bson"
)

// func TestAddMethod(t *testing.T) {
// 	db := CreateTaskDBCore(config.MongoURL, config.MongoTest)
// 	id := bson.NewObjectId()
// 	task := models.Task{ID: id, Name: "First Task"}
// 	res, _ := db.Add(task)

// 	assert.Equal(t, res.(string), id.Hex())
// }

func TestDeleteMethod(t *testing.T) {
	db := CreateTaskDBCore(config.MongoURL, config.MongoTest)
	id := bson.NewObjectId()
	task := models.Task{ID: id, Name: "First Task"}
	res, _ := db.Add(task)
	res, _ = db.Delete(id)

	assert.Equal(t, res.(int64), int64(1))
}

func TestGetMethod(t *testing.T) {
	db := CreateTaskDBCore(config.MongoURL, config.MongoTest)
	user := models.User{Username: "Anton"}
	taskJSON, _ := db.Get(user, false, false)
	result := "[{\"_id\":\"5dee57bc2166b12990c90101\",\"Name\":\"Good Task\",\"Priority\":0,\"Timing\":0,\"Subtask\":null,\"User\":{\"Username\":\"Anton\",\"Login\":\"\",\"Password\":null},\"IsComplete\":false,\"IsFail\":false},{\"_id\":\"5dee57c62166b12990c90102\",\"Name\":\"Bad Task\",\"Priority\":0,\"Timing\":0,\"Subtask\":null,\"User\":{\"Username\":\"Anton\",\"Login\":\"\",\"Password\":null},\"IsComplete\":false,\"IsFail\":false},{\"_id\":\"5dee57d62166b12990c90103\",\"Name\":\"Task\",\"Priority\":0,\"Timing\":0,\"Subtask\":null,\"User\":{\"Username\":\"Anton\",\"Login\":\"\",\"Password\":null},\"IsComplete\":false,\"IsFail\":false}]"
	assert.Equal(t, taskJSON, result)
}

func TestGetByIDMethod(t *testing.T) {
	db := CreateTaskDBCore(config.MongoURL, config.MongoTest)
	taskJSON, _ := db.Get_By_ID("5dee57c62166b12990c90102")
	result := "{\"_id\":\"5dee57c62166b12990c90102\",\"Name\":\"Bad Task\",\"Priority\":0,\"Timing\":0,\"Subtask\":null,\"User\":{\"Username\":\"Anton\",\"Login\":\"\",\"Password\":null},\"IsComplete\":false,\"IsFail\":false}"
	assert.Equal(t, taskJSON, result)
}