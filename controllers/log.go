package controllers

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogUser struct {
	Method string    `json:"method"`
	IP     string    `json:"ip"`
	Path   string    `json:"path"`
	Time   time.Time `json:"time"`
}

func Logger(c *gin.Context) {
	var client *mongo.Client
	var collection *mongo.Collection

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoHost := os.Getenv("MONGO_HOST")
	mongoPort := os.Getenv("MONGO_PORT")
	mongoURI := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection = client.Database("db").Collection("logger")

	logger := LogUser{}

	logger.Method = c.Request.Method
	logger.IP = c.ClientIP()
	logger.Time = time.Now()

	collection.InsertOne(context.TODO(), logger)
}
