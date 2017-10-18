package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "Sobin"
	my_app := "Data"
	return c.Render(greeting, my_app)
}
