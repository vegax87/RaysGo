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
	
	this.Data["Post"] = node
	this.Data["User"] = models.GetUser(node.Uid)
	this.Data["CanEdit"] = this.canEditPost(node)
	this.TplNames = "post/view.html"
}

// My posts
func (this *PostController) List(){
	this.TplNames = "post/myposts.html"
	this.Data["Title"] = "My posts"

	posts := make([]*models.Node, 0)
	err := models.Engine.Where("Uid = ?", this.User().Id).Find(&posts)
	if err == nil{
		for _, post := range posts{
			post.ParseContent()
		}
		this.Data["Posts"] = posts
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
		form models.PostForm
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
			post.Status = models.GetStatus(form.Status)
			post.Tid = models.GetNodeType("post")

			if _, err := models.Engine.Insert(&post); err == nil{
				this.FlashNotice("Post created successfully.")
				this.SaveFlash()
				this.Redirect("/post/view/" + fmt.Sprintf("%d",post.Id), 302)
				return
			}
		} else {
			for _, m := range valid.Errors {
				this.FlashError(m.Key + " : " + m.Message)
			}
		}
	}

	this.SaveFlash()
	this.Data["Form"] = form
	this.Data["TYPE_"+form.ContentType] = true
	this.Data["STATUS_" + form.Status] = true
	this.Data["Title"] = "New post"
	this.TplNames = "post/new.html"
}

func (this *PostController) Edit(){
	id,_ := helpers.Str2Int64(this.GetParam(":id"))
	var post *models.Node
	if post = models.GetNode(id); post == nil {
		this.Abort("404")
	}
	if !this.canEditPost(post) {
		this.FlashError("Sorry, you don't have the permission to edit the post!")
		this.SaveFlash()
		this.Redirect("/post/view/" + fmt.Sprintf("%d", post.Id), 302)
	}

	form := models.PostForm{}
	form.SetData(post)

	this.Data["Title"] = "Edit - " + post.Title
	this.Data["Post"] = post
	this.Data["Form"] = form
	this.Data["TYPE_" + form.ContentType] = true
	this.Data["STATUS_" + form.Status] = true
	this.TplNames = "post/edit.html"

}

func (this *PostController) EditPost(){
	var(
		post *models.Node
		form models.PostForm
		valid validation.Validation
		err error
	)

	id,_ := helpers.Str2Int64(this.GetParam(":id"))
	if post = models.GetNode(id); post == nil || !this.canEditPost(post) {
		this.Abort("404")
	}

	if err = this.ParseForm(&form); err != nil {
		this.Abort("404")
	}

	result := false
	if ok, e := valid.Valid(form); ok && e == nil{
		post.Title = form.Title
		post.Content = form.Content
		post.ContentType = form.ContentType
		post.CreateTime = time.Now()
		post.Uid = this.User().Id
		post.Status = models.GetStatus(form.Status)
		post.Tid = models.GetNodeType("post")

		if _, err = models.Engine.Id(id).Update(post); err == nil{
			result = true
			this.FlashNotice("Post updated successfully.")
		} else {
			this.FlashError("Post updated failed.")
		}

	} else {
		for _, m := range valid.Errors {
			this.FlashError(m.Key + " : " + m.Message)
		}
	}
	this.SaveFlash()
	if result {
		this.Redirect("/post/view/" + fmt.Sprintf("%d", post.Id), 302)
	}
	this.Data["Title"] = "Edit - " + post.Title
	this.Data["Post"] = post
	this.Data["Form"] = form
	this.TplNames = "post/edit.html"
}

func (this *PostController) canEditPost(post *models.Node) bool{
	if this.isLogin && (this.User().Id == post.Uid || this.User().Rid == models.ROLE_ADMIN){
		return true
	}
	return false
}

func (this *PostController) Delete(){
	id,_ := helpers.Str2Int64(this.GetParam(":id"))
	if post := models.GetNode(id); post != nil && this.canEditPost(post) {
		models.Engine.Id(id).Delete(&models.Node{})
		if _, err := models.Engine.Id(id).Delete(new(models.Node)); err == nil {
			this.FlashNotice("Post was deleted successfully.")
			this.SaveFlash()
			this.Redirect("/myposts", 302)
		} else {
			this.FlashNotice("Deleting post failed.")
			this.SaveFlash()
			// TODO : return to the former uri
		}
		
	}

	this.Abort("404")
}