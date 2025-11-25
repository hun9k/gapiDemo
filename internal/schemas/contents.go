package schemas

import (
	"database/sql"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Contents struct {
	StringField     string          `gorm:"string_field"`
	NullStringField sql.NullString  `gorm:"null_string_field"`
	IntField        int             `gorm:"int_field"`
	NullIntField    sql.NullInt64   `gorm:"null_int_field"`
	FloatField      float64         `gorm:"float_field"`
	NullFloatField  sql.NullFloat64 `gorm:"null_float_field"`
	BoolField       bool            `gorm:"bool_field"`
	NullBoolField   sql.NullBool    `gorm:"null_bool_field"`
	TimeField       time.Time       `gorm:"time_field"`
	NullTimeField   sql.NullTime    `gorm:"null_time_field"`
	DateField       datatypes.Date  `gorm:"date_field"`
	TimeFiled       datatypes.Time  `gorm:"time_filed"`
	gorm.Model
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
