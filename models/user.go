package models

import (
        "time"
)

type User struct{
	Uid int `orm:"pk;auto"`
	Role_id int `orm:"index"`
	Name string `orm:"unique;size(255)"`
	Email string `orm:"unique;size(255)"`
	Password string `orm:"size(255)"`
	CreateTime  time.Time `orm:"null;auto_now_add;type(datetime);column(create_time)"`
	LoginTime  time.Time `orm:"null;auto_now_add;type(datetime);column(login_time)"`
	Picture string `orm:"null;size(255)"`
	Signature string `orm:"null;size(255)"`
	Status int `orm:"default(0)"`
}

func (m *User) TableName() string{
	return "user"
}
