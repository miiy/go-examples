package main

import (
	"app/protobuf/address"
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

func main() {
	var cb address.ContactBook

	p1 := address.Person{
		Name:   "张三",
		Gender: address.GenderType_MALE,
		Number: "123456",
	}
	fmt.Println(p1)
	cb.Persons = append(cb.Persons, &p1)
	// 序列化
	data, err := proto.Marshal(&p1)
	if err != nil {
		fmt.Printf("marshal failed,err:%v\n", err)
		return
	}
	os.WriteFile("./proto.dat", data, 0644)

	data2, err := os.ReadFile("./proto.dat")
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	var p2 address.Person
	proto.Unmarshal(data2, &p2)
	fmt.Println(p2)

}
