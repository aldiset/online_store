package controllers

import (
	"context"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogUser struct {
	RequestID string    `json:"request_id"`
	Method    string    `json:"method"`
	IP        string    `json:"ip"`
	Path      string    `json:"path"`
	Time      time.Time `json:"time"`
}

func Logger(c *gin.Context) {
	var client *mongo.Client
	var collection *mongo.Collection

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection = client.Database("db").Collection("logger")

	request_id, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	logger := LogUser{}

	logger.RequestID = string(request_id)
	logger.Method = c.Request.Method
	logger.IP = c.ClientIP()
	logger.Time = time.Now()

	collection.InsertOne(context.TODO(), logger)
}
