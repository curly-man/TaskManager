package dbcore

import (
	"github.com/stretchr/testify/assert"
	// "TM/config"
	"TM/dbclient/dbclientmock"
	"TM/models"
	"testing"

	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	// "TM/dbclient"
)

func TestAddMethod(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	dbclient := dbclientmock.NewMockDBClient(ctrl)
	db := CreateTaskDBCore(dbclient)
	task := models.Task{ID: bson.NewObjectId()}
	dbclient.EXPECT().NewClient().Return(nil)
	dbclient.EXPECT().Connect().Return(nil)
	dbclient.EXPECT().Disconnect().Return(nil)
	dbclient.EXPECT().Ping().Return(nil)
	dbclient.EXPECT().InsertOne(task).Return(&mongo.InsertOneResult{InsertedID: task.ID.String()}, nil)
	res, _ := db.Add(task)
	assert.Equal(t, res, bson.ObjectId(task.ID.String()).Hex())
}

func TestDeleteMethod(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	dbclient := dbclientmock.NewMockDBClient(ctrl)
	db := CreateTaskDBCore(dbclient)
	id := bson.NewObjectId()
	dbclient.EXPECT().NewClient().Return(nil)
	dbclient.EXPECT().Connect().Return(nil)
	dbclient.EXPECT().Disconnect().Return(nil)
	dbclient.EXPECT().Ping().Return(nil)
	filter := primitive.D{{"_id", id}}
	dbclient.EXPECT().DeleteOne(filter).Return(&mongo.DeleteResult{DeletedCount: int64(1)}, nil)

	res, _ := db.Delete(id)
	print(12)
	assert.Equal(t, res, int64(1))
}

func TestChangeMethod(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	dbclient := dbclientmock.NewMockDBClient(ctrl)
	db := CreateTaskDBCore(dbclient)
	task := models.Task{ID: bson.NewObjectId()}
	dbclient.EXPECT().NewClient().Return(nil)
	dbclient.EXPECT().Connect().Return(nil)
	dbclient.EXPECT().Disconnect().Return(nil)
	dbclient.EXPECT().Ping().Return(nil)
	filter := primitive.D{{"_id", task.ID}}
	update := primitive.D{
		{
			"$set",
			primitive.D{
				{"name", task.Name},
				{"priority", task.Priority},
				{"iscomplete", task.IsComplete},
				{"timing", task.Timing},
				{"isfail", task.IsFail},
			},
		},
	}
	dbclient.EXPECT().UpdateOne(filter, update).Return(&mongo.UpdateResult{ModifiedCount: int64(1)}, nil)

	res, _ := db.Change(task)
	assert.Equal(t, res, int64(1))
}

func TestCompleteMethod(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	dbclient := dbclientmock.NewMockDBClient(ctrl)
	db := CreateTaskDBCore(dbclient)
	id := bson.NewObjectId()
	dbclient.EXPECT().NewClient().Return(nil)
	dbclient.EXPECT().Connect().Return(nil)
	dbclient.EXPECT().Disconnect().Return(nil)
	dbclient.EXPECT().Ping().Return(nil)
	filter := primitive.D{{"_id", id}}
	update := primitive.D{
		{
			"$set",
			primitive.D{
				{"iscomplete", true},
				{"isfail", false},
			},
		},
	}
	dbclient.EXPECT().UpdateOne(filter, update).Return(&mongo.UpdateResult{ModifiedCount: int64(1)}, nil)

	res, _ := db.Complete(id)
	assert.Equal(t, res, int64(1))
}

func TestGetByIDMethod(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	dbclient := dbclientmock.NewMockDBClient(ctrl)
	db := CreateTaskDBCore(dbclient)
	id := bson.NewObjectId()
	dbclient.EXPECT().NewClient().Return(nil)
	dbclient.EXPECT().Connect().Return(nil)
	dbclient.EXPECT().Disconnect().Return(nil)
	dbclient.EXPECT().Ping().Return(nil)
	filter := primitive.D{{"_id", id}}
	dbclient.EXPECT().FindOne(filter).Return(&mongo.SingleResult{})

	res, _ := db.GetByID(id.Hex())
	assert.Equal(t, res, "null")
}

// func TestGetMethod(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	dbclient := dbclientmock.NewMockDBClient(ctrl)
// 	db := CreateTaskDBCore(dbclient)
// 	dbclient.EXPECT().NewClient().Return(nil)
// 	dbclient.EXPECT().Connect().Return(nil)
// 	dbclient.EXPECT().Disconnect().Return(nil)
// 	dbclient.EXPECT().Ping().Return(nil)
// 	user := models.User{}
// 	filter := primitive.D{
// 		{"user", user},
// 		{"iscomplete", false},
// 		{"isfail", false},
// 	}
// 	dbclient.EXPECT().Find(filter).Return(&mongo.Cursor{Current: }, nil)
	
// 	res, _ := db.Get(user, false, false)
// 	assert.Equal(t, res, "null")
// }

// func TestAddMethodNewClientFail(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	dbclient := dbclientmock.NewMockDBClient(ctrl)
// 	db := CreateTaskDBCore(dbclient)
// 	task := models.Task{ID: bson.NewObjectId()}
// 	dbclient.EXPECT().NewClient().Return(driver.ErrFilterType)
// 	dbclient.EXPECT().Connect().Return(nil)
// 	dbclient.EXPECT().Disconnect().Return(nil)
// 	dbclient.EXPECT().Ping().Return(nil)
// 	dbclient.EXPECT().InsertOne(task).Return(&mongo.InsertOneResult{InsertedID: task.ID.String()}, nil)
// 	res, _ := db.Add(task)
// 	assert.Equal(t, res, bson.ObjectId(task.ID.String()).Hex())
// }

// func TestDeleteMethod(t *testing.T) {
// 	dbClient := dbclient.NewDBClient()
// 	db := CreateTaskDBCore(dbClient)
// 	id := bson.NewObjectId()
// 	task := models.Task{ID: id, Name: "First Task"}
// 	res, _ := db.Add(task)
// 	res, _ = db.Delete(id)

// 	assert.Equal(t, res.(int64), int64(1))
// }

// func TestGetMethod(t *testing.T) {
// 	db := CreateTaskDBCore(config.MongoURL, config.MongoTest)
// 	user := models.User{Username: "Anton"}
// 	taskJSON, _ := db.Get(user, false, false)
// 	result := "[{\"_id\":\"5dee57bc2166b12990c90101\",\"Name\":\"Good Task\",\"Priority\":0,\"Timing\":0,\"Subtask\":null,\"User\":{\"Username\":\"Anton\",\"Login\":\"\",\"Password\":null},\"IsComplete\":false,\"IsFail\":false},{\"_id\":\"5dee57c62166b12990c90102\",\"Name\":\"Bad Task\",\"Priority\":0,\"Timing\":0,\"Subtask\":null,\"User\":{\"Username\":\"Anton\",\"Login\":\"\",\"Password\":null},\"IsComplete\":false,\"IsFail\":false},{\"_id\":\"5dee57d62166b12990c90103\",\"Name\":\"Task\",\"Priority\":0,\"Timing\":0,\"Subtask\":null,\"User\":{\"Username\":\"Anton\",\"Login\":\"\",\"Password\":null},\"IsComplete\":false,\"IsFail\":false}]"
// 	assert.Equal(t, taskJSON, result)
// }

// func TestGetByIDMethod(t *testing.T) {
// 	db := CreateTaskDBCore(config.MongoURL, config.MongoTest)
// 	taskJSON, _ := db.Get_By_ID("5dee57c62166b12990c90102")
// 	result := "{\"_id\":\"5dee57c62166b12990c90102\",\"Name\":\"Bad Task\",\"Priority\":0,\"Timing\":0,\"Subtask\":null,\"User\":{\"Username\":\"Anton\",\"Login\":\"\",\"Password\":null},\"IsComplete\":false,\"IsFail\":false}"
// 	assert.Equal(t, taskJSON, result)
// }
