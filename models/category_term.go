package models

import (
        //"github.com/astaxie/beego/orm"
)

// todo
type CategoryTerm struct{
	Tid int
	Pid int `orm:"index"` // parent id
	Cid int `orm:"index"` // category id
	Uid int `orm:"index"` 
	Name string `orm:"size(255)"`
	Weigth int
}


func (m *CategoryTerm) TableName() string{
	return TableName("category_term")
}
