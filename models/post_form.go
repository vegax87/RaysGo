package models

import(
)

type NewPostForm struct{
	Title string `valid:"Required;MinSize(4);MaxSize(60)"`
	Content string
	ContentType string `valid:"Required"`
	Tags string
	OverrideUri bool
	uri string
	Status string `valid:"Required"`
}