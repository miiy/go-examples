package main

import (
	"database/sql"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

const (
	TransServerUrl = "http://localhost:8080"
	DtmServer      = "http://localhost:36789/api/dtmsvr"
)

type TransReq struct {
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
	r.GET("/Trans", func(context *gin.Context) {
		req := TransReq{
			Uid:    1,
			Amount: 50,
		}
		msg := dtmcli.NewMsg(DtmServer, uuid.New().String()).
			//Add(TransServerUrl+"/transOut", req).
			Add(TransServerUrl+"/TransIn", req)
		msg.DoAndSubmitDB(TransServerUrl+"/QueryPrepareB", db, func(tx *sql.Tx) error {

		})
		err := msg.Prepare(TransServerUrl + "Query")
		if err != nil {
			log.Println(err)
		}
		log.Println("busi trans submit")
		err = msg.Submit()
		if err != nil {
			log.Println(err)
		}
		context.String(200, "Gid: %s", msg.Gid)
	})
	r.POST("/Query", func(context *gin.Context) {
	})
	r.POST("/TransOut", func(context *gin.Context) {
	})
	r.POST("/TransIn", func(context *gin.Context) {
		context.String(200, "")
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
