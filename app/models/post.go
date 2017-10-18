package models

type Post struct {
	Id         int64
	Title      string
	Content    string
	Created_by int64
	Archived   bool
	Published  bool
}
