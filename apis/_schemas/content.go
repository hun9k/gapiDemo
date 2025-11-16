package _schemas

import "gorm.io/gorm"

type Contents struct {
	gorm.Model
	Subject  string `json:"subject,omitempty" gorm:"subject"`
	AuthorId uint   `json:"author_id,omitempty" gorm:"author_id"`
}
