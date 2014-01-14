package models

import (
	"strings"
)

// CanRegistered checks if the username or e-mail is available.
func CanRegistered(userName string, email string) (bool, bool, error) {

	var users []User
	if err := Engine.Where("Name = ? or Email = ?", userName, email).Find(&users); err != nil {
		return false, false, err
	}

	e1 := true
	e2 := true

	if len(users) > 0 {
		for _, m := range users {
			if e1 && m.Name == userName {
				e1 = false
			}
			if e2 && m.Email == email {
				e2 = false
			}
		}
	}

	return e1, e2, nil
}

// check if exist user by username or email
func HasUser(user *User, username string) bool {
	var has bool
	if strings.IndexRune(username, '@') == -1 {
		user.Name = username
		has, _ = Engine.Get(&user)
	} else {
		user.Email = username
		has, _ = Engine.Get(&user)
	}
	if has {
		return true
	}
	return false
}

func GetUser(id int64) *User {
	user := &User{Id: id}
	if has, e := Engine.Get(user); has && e == nil {
		return user
	}

	return nil
}

func (this *User) IsAdmin() bool {
	return this.Rid == ROLE_ADMIN
}
