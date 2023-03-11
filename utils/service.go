package utils

import (
	"fmt"
	"reflect"
)

type RegisteredServices = map[string]interface{}

var AppServices = make(RegisteredServices)

func RegisterService(service interface{}) {
	name := reflect.TypeOf(service).Name()
	AppServices[name] = service
}

func GetService(key string) (interface{}, error) {
	service, ok := AppServices[key]
	if !ok {
		return nil, fmt.Errorf("no service registered with token %s", key)
	}

	return service, nil
}
