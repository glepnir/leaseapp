// Package router provides ...
package router

import (
	"leaseapp/apis"

	"leaseapp/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/login", apis.Login)
	authrequired := router.Group("/apis", jwt.JWTAuth())
	{
		authrequired.POST("/addemp", apis.AddEmployee)
	}
	return router
}
