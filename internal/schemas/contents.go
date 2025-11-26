package schemas

import (
	"database/sql"
	"time"

	"github.com/hun9k/gapi/schema"
)

type Contents struct {
	StringField     string              `gorm:"string_field"`
	NullStringField sql.Null[string]    `gorm:"null_string_field"`
	IntField        int                 `gorm:"int_field"`
	NullIntField    sql.Null[int64]     `gorm:"null_int_field"`
	Float64Field    float64             `gorm:"float64_field"`
	NullFloatField  sql.Null[float64]   `gorm:"null_float_field"`
	BoolField       bool                `gorm:"bool_field"`
	NullBoolField   sql.Null[bool]      `gorm:"null_bool_field"`
	TimeField       time.Time           `gorm:"time_field"`
	NullTimeField   sql.Null[time.Time] `gorm:"null_time_field"`
	schema.Model
}
