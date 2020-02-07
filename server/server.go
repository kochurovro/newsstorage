package server

import (
	"context"

	"github.com/kochurovro/storageservice/proto/spec"
)

type newsServer struct {
}

func (n newsServer) GetNews(ctx context.Context, request *spec.GetNewsRequest) (*spec.GetNewsResponse, error) {

	panic("implement me")
}

func NewNewsServer() spec.NewsServiceServer {
	return &newsServer{}
}
