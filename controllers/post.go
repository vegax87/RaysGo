package controllers

import (
	"RaysGo/helpers"
	"RaysGo/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"strings"
	"time"
)

type PostController struct {
	AuthController
}

// Actions only available for login users
func (this *PostController) AuthActions() []string {
	return []string{"List", "New", "NewPost", "Edit", "EditPost", "Delete"}
}

// View a post
func (this *PostController) View() {
	id, _ := helpers.Str2Int64(this.GetParam(":id"))

	var node *models.Node
	if node = models.GetNode(id); node == nil {
		this.Abort("404")
	}

	if tags, err := models.GetNodeTags(node.Uid, node.Id); err == nil {
		this.Data["Tags"] = *tags
	}

	this.Data["Title"] = node.Title
	this.Data["Post"] = node
	this.Data["User"] = models.GetUser(node.Uid)
	this.Data["CanEdit"] = this.canEditPost(node)
	this.TplNames = "post/view.html"
}

// My posts
func (this *PostController) List() {
	this.TplNames = "post/myposts.html"
	this.Data["Title"] = "My posts"

	posts := make([]*models.Node, 0)
	err := models.Engine.Where("Uid = ?", this.User().Id).Find(&posts)
	if err == nil {
		this.Data["Posts"] = posts
		if tags, err := models.GetUserTags(this.User().Id); err == nil {
			this.Data["Tags"] = *tags
		}
	} else {
		this.Abort("404")
	}
}

// Create new post
func (this *PostController) New() {
	this.Data["Title"] = "New post"
	this.TplNames = "post/new.html"
}

// Post implemented creating new post
func (this *PostController) NewPost() {
	var (
		valid validation.Validation
		form  models.PostForm
		post  models.Node
		err   error
	)

	if err = this.ParseForm(&form); err == nil {
		if ok, e := valid.Valid(form); ok && e == nil {
			post.Title = form.Title
			post.Content = form.Content
			post.ContentType = form.ContentType
			post.CreateTime = time.Now()
			post.Uid = this.User().Id
			post.Status = models.GetStatus(form.Status)
			post.Tid = models.GetNodeType("post")

			tag := strings.TrimSpace(form.Tags)
			tags := make([]string, 0)
			tagarr := strings.Split(tag, ",")
			for _, v := range tagarr {
				if t := strings.TrimSpace(v); t != "" {
					tags = append(tags, t)
				}
			}

			if _, err := models.Engine.Insert(&post); err == nil {
				models.AddTags(this.User().Id, post.Id, tags)
				this.FlashNotice("Post created successfully.")
				this.SaveFlash()
				this.Redirect("/post/view/"+fmt.Sprintf("%d", post.Id), 302)
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
	this.Data["STATUS_"+form.Status] = true
	this.Data["Title"] = "New post"
	this.TplNames = "post/new.html"
}

// Edit a post
func (this *PostController) Edit() {
	id, _ := helpers.Str2Int64(this.GetParam(":id"))
	var post *models.Node
	if post = models.GetNode(id); post == nil {
		this.Abort("404")
	}
	if !this.canEditPost(post) {
		this.FlashError("Sorry, you don't have the permission to edit the post!")
		this.SaveFlash()
		this.Redirect("/post/view/"+fmt.Sprintf("%d", post.Id), 302)
	}

	form := models.PostForm{}
	form.SetData(post)

	tagStr := ""
	if tags, err := models.GetNodeTags(post.Uid, post.Id); err == nil {
		comma := ""
		for _, tag := range *tags {
			tagStr = tagStr + comma + tag.Name
			comma = ", "
		}
		form.Tags = tagStr
	}

	this.Data["Title"] = "Edit - " + post.Title
	this.Data["Post"] = post
	this.Data["Form"] = form
	this.Data["TYPE_"+form.ContentType] = true
	this.Data["STATUS_"+form.Status] = true
	this.TplNames = "post/edit.html"

}

// Post implemented editing a post
func (this *PostController) EditPost() {
	var (
		post  *models.Node
		form  models.PostForm
		valid validation.Validation
		err   error
	)

	id, _ := helpers.Str2Int64(this.GetParam(":id"))
	if post = models.GetNode(id); post == nil || !this.canEditPost(post) {
		this.Abort("404")
	}

	if err = this.ParseForm(&form); err != nil {
		this.Abort("404")
	}

	result := false
	if ok, e := valid.Valid(form); ok && e == nil {
		post.Title = form.Title
		post.Content = form.Content
		post.ContentType = form.ContentType
		post.CreateTime = time.Now()
		post.Uid = this.User().Id
		post.Status = models.GetStatus(form.Status)
		post.Tid = models.GetNodeType("post")

		tag := strings.TrimSpace(form.Tags)
		tags := make([]string, 0)
		tagarr := strings.Split(tag, ",")
		for _, v := range tagarr {
			if t := strings.TrimSpace(v); t != "" {
				tags = append(tags, t)
			}
		}

		if _, err = models.Engine.Id(id).Update(post); err == nil {
			models.AddTags(this.User().Id, post.Id, tags)
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
		this.Redirect("/post/view/"+fmt.Sprintf("%d", post.Id), 302)
	}
	this.Data["Title"] = "Edit - " + post.Title
	this.Data["Post"] = post
	this.Data["Form"] = form
	this.TplNames = "post/edit.html"
}

func (this *PostController) canEditPost(post *models.Node) bool {
	if this.isLogin && (this.User().Id == post.Uid || this.User().Rid == models.ROLE_ADMIN) {
		return true
	}
	return false
}

func (this *PostController) Delete() {
	id, _ := helpers.Str2Int64(this.GetParam(":id"))
	post := models.GetNode(id)
	if post != nil && this.canEditPost(post) {
		if err := models.DelNode(id); err == nil {
			this.FlashNotice("Post was deleted successfully.")
			this.SaveFlash()
			this.Redirect("/myposts", 302)
			return
		}
	}

	this.FlashNotice("Deleting post failed.")
	this.SaveFlash()
	// TODO : return to the former uri
	this.Abort("404")
}

func (this *PostController) Comment() {
	id, _ := helpers.Str2Int64(this.GetParam(":id"))
	post := models.GetNode(id)
	if post == nil {
		this.Abort("404")
	}
	// save comment
	this.Redirect("/post/view/" + fmt.Sprintf("%d", id), 302)
}

// TODO
func (this *PostController) Tag(){
	name := this.GetParam(":name")
	tag := models.CategoryTerm{Name : name, Uid : this.User().Id}
	if has, err := models.Engine.Get(&tag); !has || err != nil{
		this.Abort("404")
	}
	posts := make([]models.Node, 0)
	models.Engine.Join("inner","node_category_term","node_category_term.nid = node.id").Where("node_category_term.tid = ?", tag.Id).Find(&posts)
	
	if tags, err := models.GetUserTags(this.User().Id); err == nil {
		this.Data["Tags"] = *tags
	}

	this.Data["Posts"] = posts 
	this.Data["Tag"] = tag
	this.Data["Title"] = tag.Name
	this.TplNames = "post/tag.html"
}
