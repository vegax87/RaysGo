package models

import (
       // "github.com/astaxie/beego/orm"
)

// TODO
type Category struct{
	Cid int
	Uid int `orm:"index"`
	Name string `orm:"size(255)"`
	Description string
}


func (m *Category) TableName() string{
	return TableName("category")
}
