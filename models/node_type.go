package models

import (
        //"github.com/astaxie/beego/orm"
)

type NodeType struct{
	Tid int
	Name string `orm:"size(255)"`
	Description string
}


func (m *NodeType) TableName() string{
	return TableName("node_type")
}
