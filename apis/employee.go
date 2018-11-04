// Package apis provides ...
package apis

import (
	"leaseapp/data"
	"leaseapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"leaseapp/utils"
	"log"

	"gopkg.in/mgo.v2/bson"
)

const (
	db         = "leaseapp"
	collection = "EmployeeModel"
)

var (
	dao = model.Employee{}
)

func AddEmployee(c *gin.Context) {
	phonenum := c.PostForm("phone")
	pwd := c.PostForm("password")
	ename := c.PostForm("empname")
	emprole := c.PostForm("role")
	exit := mongoose.IsExist(db, collection, bson.M{"phone": phonenum})
	if exit {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "用户已存在",
		})
	} else {
		hashpwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalln(err)
		}
		hashedpwd := string(hashpwd)
		employee := model.Employee{
			ID:       bson.NewObjectId(),
			Phone:    phonenum,
			Password: hashedpwd,
			EmpName:  ename,
			Role:     emprole,
		}
		err = dao.AddEmployee(employee)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"msg":     "保存失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "保存成功",
		})
	}

}
func Login(c *gin.Context) {
	phonenum := c.PostForm("phone")
	pwd := c.PostForm("password")
	emp := model.Employee{}
	exist := mongoose.IsExist(db, collection, bson.M{"phone": phonenum})
	if exist {
		result, err := emp.FindUserByPhone(phonenum)
		err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(pwd))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"msg":     "用户名或密码错误",
			})
			return
		}
		gwt.GenerateToken(c, result)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "登录失败",
		})
	}
}
