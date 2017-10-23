package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"thought_books/app/models"
)

type Post struct {
	*revel.Controller
}

func (p Post) Index() revel.Result {
	var post models.Post
	fmt.Println("**************")
	fmt.Println(post)
	return p.Render(post)
}

func (p Post) New() revel.Result {
	return p.Render()
}

func (p Post) Create(title, content string) revel.Result {
	post := models.Post{Title: title, Content: content}

	p.Validation.Required(title).Message("Should enter title!")
	p.Validation.MinSize(content, 3).Message("Should add valid content")
	if p.Validation.HasErrors() {
		p.Validation.Keep()
		p.FlashParams()
		return p.Redirect(Post.New)
	}
	p.Create(post)
	fmt.Println("+++++++++++++++")
	return p.Redirect(Post.Index)
}
