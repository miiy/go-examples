package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"hash/crc32"
	"time"
)

const (
	SnowflakeNode = 1

	dbShardNum = 32
	tableShardNum = 1024
)

func tableName(key, tableName, dataSourceKey string, tableShardNum, dbShardNum int) (finalTableName, finalDataSourceKey string) {
	tablePosition := crc32.ChecksumIEEE([]byte(key)) % uint32(tableShardNum)
	dbPosition := tablePosition / uint32(tableShardNum / dbShardNum)
	finalTableName = fmt.Sprintf("%s_%d", tableName, tablePosition)
	finalDataSourceKey = fmt.Sprintf("%s_%d", dataSourceKey, dbPosition)
	return
}

func main()  {
	sNode, err := snowflake.NewNode(SnowflakeNode)
	if err != nil {
		panic(err)
	}

	uTableName := "users"
	dataSourceKey := "some-proxy/users_db"

	for i := 0; i < 200; i++ {
		go func() {
			id := sNode.Generate().String()
			fmt.Println(tableName(id, uTableName, dataSourceKey, tableShardNum, dbShardNum))
		}()
	}
	time.Sleep(1 * time.Second)
}
