package utils

import (
	"fmt"
	"reflect"
)

func InitDependencies(object interface{}, dependencies ...interface{}) {
	element := reflect.ValueOf(object).Elem()
	objectType := element.Type()
	for i := 0; i < element.NumField(); i++ {
		fmt.Println(i, objectType.Field(i).Name, objectType.Field(i).Type.String())
	}
}
