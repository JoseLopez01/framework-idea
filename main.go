package main

import (
	"framework-web/controllers"
	"framework-web/services"
)

var appControllers = []controllers.Controller{
	controllers.AppController{},
}

var appServices = []interface{}{
	services.AppService{},
}

func main() {
	controllers.AppController{}.InitDependencies()
}
