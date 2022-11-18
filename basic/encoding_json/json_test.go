package encoding_json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

// 使用 json.Unmarshal 对 interface{} 类型反序列化时，当数字的位数大于6位时，会变成科学计数法
// 方法一和方法二适用于能去确定传的参数
// 在验签的时候反序列化数据后签名对应不上，可使用方法三
func TestInterfaceUnmarshal(t *testing.T) {
	var jsonStr = `{"id":7656942, "title": "标题", "price": 1.50}`

	param := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &param)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(param) // map[id:7.656942e+06 price:1.5 title:标题]

	// 方法一：强制类型转换类型
	fmt.Println(int(param["id"].(float64))) // 7656942

	// 方法二：定义 struct 反序列化
	type Article struct {
		Id    int64   `json:"id"`
		Title string  `json:"title"`
		Price float64 `json:"price"`
	}
	var article Article
	err = json.Unmarshal([]byte(jsonStr), &article)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(article) // {7656942 标题 1.5}

	// 方法三：decoder.UseNumber
	d := json.NewDecoder(bytes.NewReader([]byte(jsonStr)))
	d.UseNumber()
	err = d.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(param) // map[id:7656942 price:1.50 title:标题]

}
