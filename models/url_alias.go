package models

import (
       // "github.com/astaxie/beego/orm"
)

// todo
type UrlAlias struct{
	Aid int
	Source string `orm:"size(255)"`
	Alias_uri string `orm:"size(255)"`
}

func (m *UrlAlias) TableName() string{
	return TableName("uri_alias")
}

