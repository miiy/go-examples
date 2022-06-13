package struct2slice

import "testing"

type TestStudent struct {
	Name  string `a:"1" b:"1"`
	Age   int    `a:"0" b:"1"`
	Class string `a:"1" b:"0"`
}

func TestStructToSlice(t *testing.T) {
	b := &TestStudent{
		Name:  "zhangsan",
		Age:   19,
		Class: "class1",
	}
	s1 := StructToSliceByTag(b, "a")
	t.Log(s1)
	s2 := StructToSliceByTag(b, "b")
	t.Log(s2)
}
