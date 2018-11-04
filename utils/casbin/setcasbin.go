// Package role provides ...
package mycasbin

import (
	"github.com/casbin/mongodb-adapter"

	"leaseapp/data"

	"github.com/casbin/casbin"
	"gopkg.in/mgo.v2/bson"
)

//权限结构
type CasbinModel struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Ptype    string        `json:"ptype" bson:"ptype"`
	RoleName string        `json:"rolename" bson:"v0"`
	Path     string        `json:"path" bson:"v1"`
	Method   string        `json:"method" bson:"v2"`
}

//添加权限
func (c *CasbinModel) AddCasbin(cm CasbinModel) bool {
	e := Casbin()
	return e.AddPolicy(cm.RoleName, cm.Path, cm.Method)

}

//持久化到数据库
func Casbin() *casbin.Enforcer {
	a := mongodbadapter.NewAdapter(mongoose.MongoUrl)
	e := casbin.NewEnforcer("conf/auth_model.conf", a)
	e.LoadPolicy()
	return e
}
