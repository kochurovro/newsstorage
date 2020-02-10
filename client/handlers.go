package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/proto"
	"github.com/kochurovro/storageservice/domain"
	"github.com/nats-io/nats.go"
)

func getNewsHandler(ctx context.Context, n *nats.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Query("id")
		if uid == "" {
			c.JSON(400, gin.H{
				"message": "empty id",
			})
			log.Println("client#get#empty id")
			return
		}
		reqData := domain.GetNewsRequest{Id: uid}
		req, err := proto.Marshal(&reqData)
		if err != nil {
			log.Println("client#getNews#marshalRequest", err)
			c.JSON(500, gin.H{
				"message": "internal error",
			})
			return
		}
		msg, err := n.Request("Discovery.NewsService", req, 100000*time.Millisecond)
		if err != nil {
			log.Println("client#getNews#request", err)
			c.JSON(500, gin.H{
				"message": "internal error",
			})
			return
		}
		if msg == nil || len(msg.Data) == 0 {
			log.Println("client#getNews#emptyMsg")
			c.Status(404)
			return
		}

		data := domain.GetNewsResponse{}
		err = proto.Unmarshal(msg.Data, &data)
		if err != nil {
			log.Println("client#getNews#unmarshalResponse", err)
			c.JSON(500, gin.H{
				"message": "internal error",
			})
			return
		}

		c.JSON(200, data.String())
	}
}
