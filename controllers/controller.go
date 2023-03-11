package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	InitRoutes(engine *gin.RouterGroup)
	InitDependencies(dependencies ...interface{})
}
