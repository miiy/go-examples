package struct2slice

import (
	"fmt"
	"reflect"
)

func StructToSliceByTag(v interface{}, t string) []interface{} {
	to := reflect.TypeOf(v)
	vo := reflect.ValueOf(v)
	if to.Kind() == reflect.Ptr {
		to = to.Elem()
		vo = vo.Elem()
	}
	var s []interface{}
	for i := 0; i < to.NumField(); i++ {
		if to.Field(i).Tag.Get(t) == "1" {
			s = append(s, fmt.Sprintf("%v", vo.Field(i)))
		}
	}
	return s
}
