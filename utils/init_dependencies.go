package utils

import (
	"reflect"
)

func InitDependencies(object interface{}) {
	element := reflect.ValueOf(object).Elem()
	for i := 0; i < element.NumField(); i++ {
		field := element.Field(i)
		if field.CanSet() {
			service, err := GetService(field.Type().Name())
			if err != nil {
				panic(err)
			}

			field.Set(reflect.ValueOf(service))
		}
	}
}
