// Package model provides ...
package model

import (
	"leaseapp/data"

	"gopkg.in/mgo.v2/bson"
)

const (
	db         = "leaseapp"
	collection = "EmployeeModel"
)

type Employee struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Phone    string        `json:"phone" bson:"phone"`
	Password string        `json:"-" bson:"passoword"`
	EmpName  string        `json:"empname" bson:"empname"`
	Role     string        `json:"role" bson:"role"`
}

//新增用户
func (e *Employee) AddEmployee(emp Employee) error {
	return mongoose.Insert(db, collection, emp)
}

//查找用户
func (e *Employee) FindUserByPhone(phone string) (Employee, error) {
	var result Employee
	err := mongoose.FindOne(db, collection, bson.M{
		"phone": phone,
	}, nil, &result)
	return result, err
}
