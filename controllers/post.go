package controllers

import(
	"time"
	"fmt"
	"RaysGo/models"
	"RaysGo/helpers"
	"github.com/astaxie/beego/validation"
)

type PostController struct{
	AuthController
}

func (this *PostController) AuthActions() []string{
	return []string{ "List", "New", "NewPost", "Edit", "EditPost", "Delete"}
}

func (this *PostController) View(){
	id,_ := helpers.Str2Int64(this.GetParam(":id"))

	var node *models.Node
	if node = models.GetNode(id); node == nil{
		this.Abort("404")
	}
	node.ParseContent()

	this.Data["Post"] = node
	this.Data["User"] = models.GetUser(node.Uid)
	this.Data["CanEdit"] = this.isLogin && (node.Uid == this.User().Id || this.User().Rid == models.ROLE_ADMIN)
	this.TplNames = "post/view.html"
}

// My posts
func (this *PostController) List(){
	posts := make([]*models.Node, 0)
	err := models.Engine.Where("Uid = ?", this.User().Id).Find(&posts)
	if err == nil{
		for _, post := range posts{
			post.ParseContent()
		}
		this.Data["Posts"] = posts
		this.TplNames = "post/myposts.html"
	} else {
		this.Abort("404")
	}
}

func (this *PostController) New(){
	this.Data["Title"] = "New post"
	this.TplNames = "post/new.html"
}

func (this *PostController) NewPost(){
	var(
		valid validation.Validation
		form models.NewPostForm
		post models.Node
		err error
	)

	if err = this.ParseForm(&form); err == nil{
		if ok, e := valid.Valid(form); ok && e == nil{
			post.Title = form.Title
			post.Content = form.Content
			post.ContentType = form.ContentType
			post.CreateTime = time.Now()
			post.Uid = this.User().Id
			post.Status = models.DRAFT
			post.Tid = models.GetNodeType("post")
			switch form.Status{
			case "published":
				post.Status = models.PUBLISHED
			case "private":
				post.Status = models.PRIVATE
			}

			if _, err := models.Engine.Insert(&post); err == nil{
				this.FlashNotice("Post created successfully.")
				this.SaveFlash()
				this.Redirect("/post/view/" + fmt.Sprintf("%d",post.Id), 302)
				return
			}
		}
	}

	this.SaveFlash()
	this.Data["Form"] = form
	this.Data["Title"] = "New post"
	this.TplNames = "post/new.html"
}

func (this *PostController) Edit(){
	this.TplNames = "post/edit.html"

}

func (this *PostController) EditPost(){
	this.TplNames = "post/edit.html"

}

func (this *PostController) Delete(){

}