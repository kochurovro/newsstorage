package main

import (
	"context"
	"log"

	"github.com/kochurovro/storageservice/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	c   *mongo.Client
	ctx context.Context
}

func (s *Storage) GetByID(uid string) (r domain.GetNewsResponse, err error) {
	collection := s.c.Database("testing").Collection("news")
	filter := bson.D{{Key: "uid", Value: uid}}
	err = collection.FindOne(s.ctx, filter).Decode(&r)
	if err != nil {
		log.Println("store#FindOne", err)
		if err == mongo.ErrNoDocuments {
			return r, domain.ErrNotFound
		}
	}
	return r, err
}

func NewStorage(ctx context.Context, c *mongo.Client) *Storage {
	return &Storage{c: c, ctx: ctx}
}
