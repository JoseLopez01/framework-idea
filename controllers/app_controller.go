package controllers

import (
	"framework-web/utils"
	"github.com/gin-gonic/gin"
	"net/http"

	"framework-web/services"
)

type AppController struct {
	Service services.AppService
}

func (controller AppController) InitRoutes(engine *gin.RouterGroup) {
	engine.GET("/message", func(context *gin.Context) {
		message := controller.Service.GetMessage()
		context.JSON(http.StatusOK, gin.H{"message": message})
	})
}

func (controller AppController) InitDependencies(dependencies ...interface{}) {
	utils.InitDependencies(&controller)
}
