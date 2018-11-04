// Package router provides ...
package router

import (
	"leaseapp/apis"
	"leaseapp/middleware/jwt"
	"leaseapp/middleware/role"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.POST("/login", apis.Login)
	authrequired := router.Group("/apis")
	authrequired.Use(jwt.JWTAuth())
	authrequired.Use(authz.AuthCheckRole())
	{
		authrequired.POST("/addemp", apis.AddEmployee)
		authrequired.POST("/addrole", apis.AddCasbin)
	}
	return router
}
