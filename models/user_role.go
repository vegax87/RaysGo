package models

import (
        //"github.com/astaxie/beego/orm"
)

// todo
type UserRole struct{
	Rid int `orm:"pk"`
	Name string `orm:"size(255)"`
	Description string
}


func (m *UserRole) TableName() string{
	return "user_role"
}

