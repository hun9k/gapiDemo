package schemas

import (
	"github.com/hun9k/gapi/schema"
	"gorm.io/gorm"
)

type Contents struct {
	Subject  string `gorm:"subject"`
	AuthorId uint   `gorm:"author_id"`
	schema.Model
}

// hooks

// before save
func (m *Contents) BeforeSave(tx *gorm.DB) error {
	return nil
}

// create
func (m *Contents) BeforeCreate(tx *gorm.DB) error {
	return nil
}

func (m *Contents) AfterCreate(tx *gorm.DB) error {
	return nil
}

// update
func (m *Contents) BeforeUpdate(tx *gorm.DB) error {
	return nil
}

func (m *Contents) AfterUpdate(tx *gorm.DB) error {
	return nil
}

// after save
func (m *Contents) AfterSave(tx *gorm.DB) error {
	return nil
}

// delete
func (m *Contents) BeforeDelete(tx *gorm.DB) error {
	return nil
}

func (m *Contents) AfterDelete(tx *gorm.DB) error {
	return nil
}

// find
func (m *Contents) AfterFind(tx *gorm.DB) error {
	return nil
}
