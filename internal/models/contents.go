package models

import (
	"github.com/hun9k/gapi/db"
	"github.com/hun9k/gapiDemo/internal/schemas"
	"gorm.io/gorm"
)

type Contents schemas.Contents

func ContentsInsert(m *Contents) error {
	return db.DB().Transaction(func(tx *gorm.DB) error {
		// create
		if r := tx.Create(&m); r.Error != nil {
			return r.Error
		}
		return nil
	})
}

func ContentsRow(id uint) (*Contents, error) {
	m := &Contents{}
	err := db.DB().Transaction(func(tx *gorm.DB) error {
		// create
		if r := tx.First(&m, id); r.Error != nil {
			return r.Error
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return m, nil
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
