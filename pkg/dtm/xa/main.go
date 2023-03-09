package main

import (
	"github.com/dtm-labs/client/dtmcli"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"log"
)

const (
	TransServerUrl = "http://localhost:8080"
	DtmServer      = "http://localhost:36789/api/dtmsvr"
)

type TransRequest struct {
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
		gid := uuid.New().String()
		err := dtmcli.XaGlobalTransaction(DtmServer, gid, func(xa *dtmcli.Xa) (*resty.Response, error) {
			transReq := TransRequest{
				Uid:    1,
				Amount: 30,
			}
			resp, err := xa.CallBranch(&transReq, TransServerUrl+"/TransOutXa")
			if err != nil {
				return resp, err
			}
			return xa.CallBranch(&transReq, TransServerUrl+"/TransInXa")
		})
		if err != nil {
			log.Println(err)
		}
		context.String(200, "Gid: %s", gid)
	})
	r.POST("/TransOutXa", func(context *gin.Context) {
	})
	r.POST("/TransInXa", func(context *gin.Context) {
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
