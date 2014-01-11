package models

import (
        //"github.com/astaxie/beego/orm"
        "time"
)

// todo
type Node struct{
	Nid int
	Uid int `orm:"index"`
	Type_id int `orm:"index"`
	Title string `orm:"size(255)"`
	Content string
	Summary string
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Promote int
	Top int
}



func (m *Node) TableName() string{
	return TableName("node")
}

