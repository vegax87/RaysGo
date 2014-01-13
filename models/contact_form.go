package models

import(

)

type ContactForm struct{
	UserName string `valid:"Required"`
	Email string `valid:"Required;Email"`
	Title string `valid:"Required"`
	Content string `valid:"Required"`
}