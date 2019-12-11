package dbcore

import (
	"TM/config"
	"TM/models"
	"context"
	"encoding/json"
	// "fmt"
	// "log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	id "gopkg.in/mgo.v2/bson"
)

type TaskDBCore interface {
	Add(task models.Task) (interface{}, error)
	Delete(taskID id.ObjectId) (interface{}, error)
	Change(task models.Task) (interface{}, error)
	Get(user models.User, iscomplete bool, isfail bool) (string, error)
	Get_By_ID(ID string) (string, error)
}

type taskDBCore struct {
	MongoURL string
	Name string
}

func CreateTaskDBCore(url string, dbName string) TaskDBCore {
	return &taskDBCore{MongoURL: url, Name: dbName}
}

func Make_Client() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoURL))
	if err != nil {
		return &mongo.Client{}, err
	}
	return client, nil
}

func Connect(ctx context.Context, client *mongo.Client) error {
	err := client.Connect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func Disconnect(ctx context.Context, client *mongo.Client) error{
	err := client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func Check_Connection(client *mongo.Client) error {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	return nil
}

func insert_task(ctx context.Context, collection *mongo.Collection, task models.Task) (interface{}, error) {
	result, err := collection.InsertOne(ctx, task)
	id.ObjectId(result.InsertedID.(string)).Hex()
	return id.ObjectId(result.InsertedID.(string)).Hex(), err
}

func find_tasks(ctx context.Context, collection *mongo.Collection, user models.User, iscomplete bool, isfail bool) ([]models.Task, error) {
	filter := bson.D{
		{"user", user},
		{"iscomplete", iscomplete},
		{"isfail", isfail},
	}
	option := options.Find()
	cursor, err := collection.Find(ctx, filter, option)
	if err != nil {
		return []models.Task{}, err
	}
	var tasks []models.Task
	for cursor.Next(ctx) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return []models.Task{}, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (db *taskDBCore) Add(task models.Task) (interface{}, error){
	ctx := context.Background()
	client, err := Make_Client()
	if err != nil{
		return "", err
	}
	err = Connect(ctx, client)
	if err != nil{
		return "", err
	}
	defer Disconnect(ctx, client)
	err = Check_Connection(client)
	if err != nil{
		return "", err
	}

	collection := client.Database(db.Name).Collection("Tasks")
	return insert_task(ctx, collection, task)
}

func (db *taskDBCore) Delete(taskID id.ObjectId) (interface{}, error) {
	ctx := context.Background()
	client, err := Make_Client()
	if err != nil{
		return "", err
	}
	err = Connect(ctx, client)
	if err != nil{
		return "", err
	}
	defer Disconnect(ctx, client)
	err = Check_Connection(client)
	if err != nil{
		return "", err
	}

	collection := client.Database(db.Name).Collection("Tasks")
	filter := bson.D{{"_id", taskID}}
	res, err := collection.DeleteOne(ctx, filter)
	return res.DeletedCount, err
}

func (db *taskDBCore) Change(task models.Task) (interface{}, error) {
	ctx := context.Background()
	client, err := Make_Client()
	if err != nil{
		return "", err
	}
	err = Connect(ctx, client)
	if err != nil{
		return "", err
	}
	defer Disconnect(ctx, client)
	err = Check_Connection(client)
	if err != nil{
		return "", err
	}

	collection := client.Database(db.Name).Collection("Tasks")

	filter := bson.D{{"_id", task.ID}}

	update := bson.D{
		{
			"$set",
			bson.D{
				{"name", task.Name},
				{"priority", task.Priority},
				{"iscomplete", task.IsComplete},
				{"timing", task.Timing},
				{"isfail", task.IsFail},
			},
		},
	}
	res, err := collection.UpdateOne(ctx, filter, update)
	return res.ModifiedCount, err
}

func (db *taskDBCore) Get(user models.User, iscomplete bool, isfail bool) (string, error) {
	ctx := context.Background()
	client, err := Make_Client()
	if err != nil{
		return "", err
	}
	err = Connect(ctx, client)
	if err != nil{
		return "", err
	}
	defer Disconnect(ctx, client)
	err = Check_Connection(client)
	if err != nil{
		return "", err
	}

	collection := client.Database(db.Name).Collection("Tasks")
	tasks, err := find_tasks(ctx, collection, user, iscomplete, isfail)
	if err != nil {
		return "", err 
	}
	bytes, err := json.Marshal(tasks)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (db *taskDBCore) Get_By_ID(ID string) (string, error) {
	ctx := context.Background()
	client, err := Make_Client()
	if err != nil{
		return "", err
	}
	err = Connect(ctx, client)
	if err != nil{
		return "", err
	}
	defer Disconnect(ctx, client)
	err = Check_Connection(client)
	if err != nil{
		return "", err
	}

	collection := client.Database(db.Name).Collection("Tasks")

	filter := bson.D{{"_id", id.ObjectIdHex(ID)}}
	result := collection.FindOne(ctx, filter)

	var task models.Task
	err = result.Decode(&task)
	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(task)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
