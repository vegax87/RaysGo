package models

import ()

type ContactForm struct {
	UserName string `valid:"Required"`
	Email    string `valid:"Required;Email"`
	Title    string `valid:"Required"`
	Content  string `valid:"Required"`
}

// TODO
func SendContactMail(title string, content string, fromEmail string, fromName string) error{
	return nil
}
