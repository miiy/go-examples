package struct2slice

import (
	"reflect"
)

func StructToSlice(v interface{}) []interface{} {
	to := reflect.TypeOf(v)
	vo := reflect.ValueOf(v)
	if to.Kind() == reflect.Ptr {
		to = to.Elem()
		vo = vo.Elem()
	}
	var s []interface{}
	for i := 0; i < to.NumField(); i++ {
		s = append(s, vo.Field(i).Interface())
	}
	return s
}

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
			s = append(s, vo.Field(i).Interface())
		}
	}
	return s
}

func StructToSliceByFieldName(v interface{}, fields []string) []interface{} {
	to := reflect.TypeOf(v)
	vo := reflect.ValueOf(v)
	if to.Kind() == reflect.Ptr {
		to = to.Elem()
		vo = vo.Elem()
	}
	var s []interface{}
	for _, field := range fields {
		s = append(s, vo.FieldByName(field).Interface())
	}
	return s
}
