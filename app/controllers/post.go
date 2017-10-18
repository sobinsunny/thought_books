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
	return p.Redirect(Post.Index)
}
