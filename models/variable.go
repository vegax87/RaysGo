package models

import (
        //"github.com/astaxie/beego/orm"
)

// todo
type Variable struct{
	Key string `orm:"pk;size(255)"`
	Value string 
}

func (m *Variable) TableName() string{
	return "variable"
}

