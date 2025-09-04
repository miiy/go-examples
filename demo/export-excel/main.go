package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"log"
	"net/http"
	"net/url"
	"reflect"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/export", func(c *gin.Context) {
		b, err := exportUsers()
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "export error.")
			return
		}
		Header := c.Writer.Header()
		Header.Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, url.QueryEscape("用户列表.xlsx")))
		Header.Set("Content-Description", "File Transfer")
		Header.Set("Content-Transfer-Encoding", "binary")
		Header.Set("Expires", "0")
		Header.Set("Cache-Control", "must-revalidate")
		c.Data(http.StatusOK, "application/octet-stream", b)
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

type User struct {
	Id       int    `colName:"ID"`
	UserName string `colName:"用户名"`
	Gender   string `colName:"性别"`
	VisitNum int    `colName:"访问量"`
}

// exportUsers 导出用户处理
func exportUsers() ([]byte, error) {
	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Close err:", err)
		}
	}()

	// 用户信息
	users := []User{
		{
			Id:       1,
			UserName: "zhang san",
			Gender:   "F",
			VisitNum: 10,
		},
		{
			Id:       1,
			UserName: "li si",
			Gender:   "M",
			VisitNum: 20,
		},
	}

	// 表头
	header := []string{"ID", "用户名", "访问量"}
	if err := file.SetSheetRow("Sheet1", "A1", &header); err != nil {
		log.Println("SetSheetRow err:", err)
	}
	// 内容
	for rowNum, row := range users {
		// 根据标签取value
		rowSlice := structToSliceByTagValues(row, "colName", header)
		for colNum, v := range rowSlice {
			if err := setCellValue(file, "Sheet1", colNum+1, rowNum+2, v); err != nil {
				log.Println("setCellValue err:", err)
			}
		}
	}
	buf, err := file.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func structToSliceByTagValues(f interface{}, tagName string, tagValues []string) []interface{} {
	to := reflect.TypeOf(f)
	vo := reflect.ValueOf(f)
	if to.Kind() == reflect.Ptr {
		to = to.Elem()
		vo = vo.Elem()
	}

	var s []interface{}
	for _, t := range tagValues {
		find := false
		for i := 0; i < to.NumField(); i++ {
			if t == to.Field(i).Tag.Get(tagName) {
				find = true
				s = append(s, vo.Field(i).Interface())
				break
			}
		}
		if !find {
			panic("not find tag " + t)
		}
	}

	return s
}

func setCellValue(f *excelize.File, sheetName string, col, row int, value interface{}) error {
	cellName, err := excelize.CoordinatesToCellName(col, row)
	if err != nil {
		return err
	}
	err = f.SetCellValue(sheetName, cellName, value)
	if err != nil {
		return err
	}
	return nil
}
