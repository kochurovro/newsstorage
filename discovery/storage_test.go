package main

import (
	"context"
	"testing"
	"time"

	"github.com/kochurovro/storageservice/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestStorage_GetByID(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}
	collection := client.Database("testing").Collection("news")
	defer collection.Drop(ctx)
	_, err = collection.InsertOne(ctx, domain.GetNewsResponse{Title: "HW", Timestamp: 20191010, Uid: "123456"})
	if err != nil {
		t.Fatal(err)
	}

	storage := NewStorage(ctx, client)
	type compfunc func(got string) bool
	f := func(name string, compare compfunc, in string) {
		t.Run(name, func(t *testing.T) {
			out, err := storage.GetByID(in)
			if err != nil {
				if err != domain.ErrNotFound {
					t.Fatal(err)
					return
				}
			}
			if !compare(out.Title) {
				t.Fatal("wrong result")
				return
			}
		})
	}
	f("empty in should return empty out", func(g string) bool { return g == "" }, "")
	f("123456 in should return HW", func(g string) bool { return g == "HW" }, "123456")

}
