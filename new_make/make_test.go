package main

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
}

func TestOrgAddDept(t *testing.T) {
	var a1 = make([]Person, 0)
	var a2 []Person
	fmt.Println(len(a1), cap(a1))
	fmt.Println(len(a2), cap(a2))
	a1 = append(a1, Person{
		name: "a",
	})
	a2 = append(a2, Person{
		name: "a",
	})
	fmt.Println(a1)
	fmt.Println(a2)

	var a3 = make([]*Person, 0)
	var a4 []*Person
	fmt.Println(len(a3), cap(a3))
	fmt.Println(len(a4), cap(a4))
	a3 = append(a3, &Person{name: "a"})
	a4 = append(a4, &Person{name: "a"})
	fmt.Println(a3)
	fmt.Println(a4)

	//var a3 Person
	//fmt.Println(a3)
	//fmt.Println(a3.Menu)
	//a3.Menu = append(a3.Menu, org.OrgOrgMenu{
	//	MId:    0,
	//	MName:  "",
	//	ResId:  0,
	//	Pid:    0,
	//	MLevel: 0,
	//	Seq:    0,
	//})
	//fmt.Println(a3)
}
