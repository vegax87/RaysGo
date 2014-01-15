package models

import (
	"RaysGo/helpers"
	"fmt"
)

func InitModel() {
	fmt.Println("init model")
}

func InsertRole() {
	roles := make([]Role, 3)
	roles[0].Id = 1
	roles[0].Name = "admin"
	roles[0].Description = "administrator"

	roles[1].Id = 2
	roles[1].Name = "authenticated"
	roles[1].Description = "authenticated user"

	roles[2].Id = 3
	roles[2].Name = "anonymous"
	roles[2].Description = "anonymous user"

	Engine.Insert(&roles)
}

func InsertUser() {
	admin := User{
		Name:     "admin",
		Email:    "admin@gmail.com",
		Password: helpers.EncryptPassword("admin", nil),
		IRole:     Role{Id : ROLE_ADMIN},
		Status:   ACTIVE,
	}

	Engine.Insert(&admin)

	user := User{
		Name:     "hello",
		Email:    "raysmond@gmail.com",
		Password: helpers.EncryptPassword("111111", nil),
		IRole:     Role{Id : ROLE_AUTHENTICATED},
		Status:   ACTIVE,
	}

	Engine.Insert(&user)
}
