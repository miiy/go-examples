package main

import (
	"github.com/dtm-labs/client/dtmcli"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

const (
	TransServerUrl = "http://localhost:8080"
	DtmServer      = "http://localhost:36789/api/dtmsvr"
)

type Request struct {
	Uid    float64
	Amount float64
}

func main() {
	r := gin.Default()
	r.GET("/")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/trans", func(context *gin.Context) {
		req := Request{
			Uid:    1,
			Amount: 50,
		}
		saga := dtmcli.NewSaga(DtmServer, uuid.New().String()).
			Add(TransServerUrl+"/TransOut", TransServerUrl+"/TransOutCompensate", req).
			Add(TransServerUrl+"/TransIn", TransServerUrl+"/TransInCompensate", req)
		err := saga.Submit()
		if err != nil {
			log.Println(err)
		}
		context.String(200, "Gid: %s", saga.Gid)
	})
	r.POST("/TransOut", func(context *gin.Context) {
	})
	r.POST("/TransOutCompensate", func(context *gin.Context) {
	})
	r.POST("/TransIn", func(context *gin.Context) {
	})
	r.POST("/TransInCompensate", func(context *gin.Context) {
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
