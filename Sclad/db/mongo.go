package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Mongo struct {
	client *mongo.Client
	dbname string
}

func (m *Mongo) InitDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	m.client, _ = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL"))) //"mongodb://localhost:27017"
	m.dbname = os.Getenv("MONGO_DBNAME")
}

func (m *Mongo) CreatePost(post Post) error {
	collection := m.client.Database(m.dbname).Collection("items")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err := collection.InsertOne(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongo) GetPost(id string) (Post, error) {
	var post Post
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := m.client.Database(m.dbname).Collection("items")

	err := collection.FindOne(ctx, Post{PostId: id}).Decode(&post)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (m *Mongo) GetUserPosts(userId string) ([]Post, error) {
	var posts []Post
	collection := m.client.Database(m.dbname).Collection("items")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"createdAt", -1}})

	cursor, err := collection.Find(ctx, bson.M{"authorId": userId}, findOptions)

	if err != nil {
		return posts, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}
	if err := cursor.Err(); err != nil {
		return posts, err
	}
	return posts, nil
}

func (m *Mongo) Check(id string) bool {
	collection := m.client.Database(m.dbname).Collection("items")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	count, _ := collection.CountDocuments(ctx, bson.M{"id": id})
	if count >= 1 {
		return true
	}
	return false
}

func (m *Mongo) IsConnected() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return m.client.Ping(ctx, nil)
}
