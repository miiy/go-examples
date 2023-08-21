package main

import "fmt"

type Book struct {
	Name string
}

func main() {
	books := []*Book{
		{"book1"},
		{"book2"},
		{"book3"},
	}
	for _, v := range books {
		v.Name = "test"
	}

	for _, v := range books {
		fmt.Println(v.Name)
	}

	//
	books2 := &[]Book{
		{"book1"},
		{"book2"},
		{"book3"},
	}
	for i, _ := range *books2 {
		(*books2)[i].Name = "test"
	}
	for _, v := range *books2 {
		fmt.Println(v.Name)
	}

	//
	books3 := []Book{
		{"book1"},
		{"book2"},
		{"book3"},
	}
	for i, _ := range books3 {
		books3[i].Name = "test"
	}

}
