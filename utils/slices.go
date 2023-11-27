package utils

import (
	"reflect"
)

func RemoveIndex(s []any, index int) []any {
	return append(s[:index], s[index+1:]...)
}
func TypeofSrt(v interface{}) string {
	return reflect.TypeOf(v).String()
}
func Typeof(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}
