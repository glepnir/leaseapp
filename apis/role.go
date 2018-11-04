// Package apis provides ...
package apis

import (
	"leaseapp/utils/casbin"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

var (
	casbins = mycasbin.CasbinModel{}
)

func AddCasbin(c *gin.Context) {
	rolename := c.PostForm("rolename")
	path := c.PostForm("path")
	method := c.PostForm("method")
	ptype := "p"
	casbin := mycasbin.CasbinModel{
		ID:       bson.NewObjectId(),
		Ptype:    ptype,
		RoleName: rolename,
		Path:     path,
		Method:   method,
	}
	isok := casbins.AddCasbin(casbin)
	if isok {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "保存成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "保存失败",
		})
	}

}
