package main

import (
	"framework-web/controllers"
	"framework-web/services"
	"framework-web/utils"
	"github.com/gin-gonic/gin"
)

var appControllers = []controllers.Controller{
	controllers.AppController{},
}

func main() {
	utils.RegisterService(services.AppService{})

	engine := gin.Default()
	group := engine.Group("/api")
	for _, controller := range appControllers {
		controller.InitDependencies()
		controller.InitRoutes(group)
	}

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
