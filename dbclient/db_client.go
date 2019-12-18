package dbclient

import (
	"context"
	"TM/models"
	"TM/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient interface{
	NewClient() error
	Connect() error
	Disconnect() error
	Ping() error
	InsertOne(task models.Task) (*mongo.InsertOneResult, error)
	Find(filter bson.D) (*mongo.Cursor, error)
	FindOne(filter bson.D) *mongo.SingleResult
	DeleteOne(filter bson.D) (*mongo.DeleteResult, error)
	UpdateOne(filter, update bson.D) (*mongo.UpdateResult, error)
}

type dbClient struct{
	client *mongo.Client
	ctx context.Context
	collection *mongo.Collection
}

func NewDBClient() DBClient{
	return &dbClient{
		ctx: context.Background(),
	}
}
func (c *dbClient) NewClient() error{
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoURL))
	c.client = client
	c.collection = client.Database(config.MongoDB).Collection(config.MongoCollection)
	return err
}
func (c *dbClient) Connect() error{
	return c.client.Connect(c.ctx)
}
func (c *dbClient) Disconnect() error{
	return c.client.Disconnect(c.ctx)
}
func (c *dbClient) Ping() error{
	return c.client.Ping(context.TODO(), nil)
}
func (c *dbClient) InsertOne(task models.Task) (*mongo.InsertOneResult, error) {
	return c.collection.InsertOne(c.ctx, task)
}
func (c *dbClient) UpdateOne(filter, update bson.D) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(c.ctx, filter, update)
}
func (c *dbClient) DeleteOne(filter bson.D) (*mongo.DeleteResult, error) {
	return c.collection.DeleteOne(c.ctx, filter)
}
func (c *dbClient) FindOne(filter bson.D) *mongo.SingleResult {
	return c.collection.FindOne(c.ctx, filter)
}
func (c *dbClient) Find(filter bson.D) (*mongo.Cursor, error) {
	return c.collection.Find(c.ctx, filter)
}