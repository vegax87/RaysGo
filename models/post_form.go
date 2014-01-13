package models

import(
)

type PostForm struct{
	Title string `valid:"Required;MinSize(4);MaxSize(60)"`
	Content string
	ContentType string `valid:"Required"`
	Tags string
	OverrideUri bool
	uri string
	Status string `valid:"Required"`
}

func (this *PostForm) SetData(post *Node){
	if post != nil {
		this.Title = post.Title
		this.Content = post.Content
		this.ContentType = post.ContentType
		this.Status = GetStatusName(post.Status)

		// TODO
		this.Tags = ""
		this.OverrideUri = false
		this.uri = ""
	}
}