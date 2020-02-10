package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func main() {
	ctx := context.Background()
	// Create NATS client connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	defer natsConnection.Close()
	log.Println("Connected to " + nats.DefaultURL)

	r := gin.Default()
	r.GET("/getNews", getNewsHandler(ctx, natsConnection))
	r.Run()

}
