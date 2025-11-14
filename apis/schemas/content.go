package models

import "gorm.io/gorm"

type Contents struct {
	gorm.Model
	Subject  string
	Summary  string
	Content  string
	AuthorId uint
}
