package main

import (
	"context"
	"log"
	"runtime"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/kochurovro/storageservice/domain"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10000*time.Second)

	natsConnection, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("discovery#mian#natsconnvection#error", err)
		return
	}
	log.Println("Connected to " + nats.DefaultURL)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal("discovery#mian#dbconnvection#dbConnect", err)
		return
	}

	collection := client.Database("testing").Collection("news")
	_, err = collection.InsertOne(ctx, domain.GetNewsResponse{Title: "HW", Timestamp: 20191010, Uid: "123456"})
	if err != nil {
		log.Fatal(err)
		return
	}

	store := NewStorage(ctx, client)
	natsConnection.Subscribe("Discovery.NewsService", func(m *nats.Msg) {
		log.Println("start discovery news service")

		r := domain.GetNewsRequest{}
		err := proto.Unmarshal(m.Data, &r)
		if err != nil {
			log.Println("discovery#can not unmarshal request", err)
			natsConnection.Publish(m.Reply, nil)
		} else {
			get, err := store.GetByID(r.Id)
			if err != nil {
				log.Println("discovery#can not get news", err)
				natsConnection.Publish(m.Reply, nil)
			} else {
				data, err := proto.Marshal(&get)
				if err != nil {
					log.Println("discovery#can not marshal", err)
					natsConnection.Publish(m.Reply, nil)
				}
				natsConnection.Publish(m.Reply, data)
			}
		}
	})
	// Keep the connection alive
	runtime.Goexit()
}
