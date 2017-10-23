package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title      string
	Content    string
	Created_by int64
	Archived   bool
	Published  bool
}
