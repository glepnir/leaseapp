// Package gwt provides ...
package gwt

import (
	"leaseapp/middleware/jwt"
	"leaseapp/models"
	"log"
	"time"

	"net/http"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginResult struct {
	Token string `json:"token"`
	model.Employee
}

//生成token
func GenerateToken(c *gin.Context, emp model.Employee) {
	j := &jwt.JWT{
		[]byte("martin"),
	}
	claims := jwt.Customclaims{
		emp.EmpName,
		emp.Phone,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), //签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), //签名过期时间 一小时
			Issuer:    "martin",                        //签名发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
	}
	log.Println(token)
	data := LoginResult{
		Token:    token,
		Employee: emp,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功",
		"data":   data,
	})
	return
}
