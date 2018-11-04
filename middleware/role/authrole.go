// Package authz provides ...
package authz

import (
	"net/http"

	"leaseapp/middleware/jwt"
	"leaseapp/utils/casbin"

	"github.com/gin-gonic/gin"
)

//权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		//根据上下文获取载荷claims 从claims获得role
		claims := c.MustGet("claims").(*jwt.Customclaims)
		role := claims.Role
		e := mycasbin.Casbin()
		//检查权限
		res, err := e.EnforceSafe(role, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    "错误消息" + err.Error(),
			})
			c.Abort()
			return
		}
		if res {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "很抱歉您没有此权限",
			})
			c.Abort()
			return
		}
	}
}
