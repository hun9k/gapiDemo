package contents

import (
	"github.com/hun9k/gapi/biz"
)

// 创建资源
type PostBody struct {
	StringField     *string  `json:"string_field,omitempty"  form:"string_field,omitempty"  binding:""`
	NullStringField *string  `json:"null_string_field,omitempty"  form:"null_string_field"  binding:"null_string_field"`
	IntField        *int     `json:"int_field,omitempty"  form:"int_field"  binding:"int_field"`
	NullIntField    *int64   `json:"null_int_field,omitempty"  form:"null_int_field"  binding:"null_int_field"`
	FloatField      *float64 `json:"float_field,omitempty"  form:"float_field"  binding:"float_field"`
	NullFloatField  *float64 `json:"null_float_field,omitempty"  form:"null_float_field"  binding:"null_float_field"`
	BoolField       *bool    `json:"bool_field,omitempty"  form:"bool_field"  binding:"bool_field"`
	NullBoolField   *bool    `json:"null_bool_field,omitempty"  form:"null_bool_field"  binding:"null_bool_field"`
	TimeField       *string  `json:"time_field,omitempty"  form:"time_field"  binding:"time_field"`
	NullTimeField   *string  `json:"null_time_field,omitempty"  form:"null_time_field"  binding:"null_time_field"`
	DateField       *string  `json:"date_field,omitempty"  form:"date_field"  binding:"date_field"`
	TimeFiled       *string  `json:"time_filed,omitempty"  form:"time_filed"  binding:"time_filed"`
}

// 删除一个资源
type DeleteIdQuery struct {
	biz.QueryStrId
}

// 删除多个资源
type DeleteQuery struct {
	biz.QueryStr
}

// 恢复一个资源
type RestoreIdQuery struct {
	biz.QueryStrId
}

// 恢复多个资源
type RestoreQuery struct {
	biz.QueryStr
}

// 更新一个资源
type PutIdQuery struct {
	biz.QueryStrId
}
type PutIdBody struct {
}

// 更新多个资源
type PutQuery struct {
	biz.QueryStr
}
type PutBody struct {
}

// 更新一个资源的部分字段
type PatchIdQuery struct {
	biz.QueryStrId
}
type PatchIdBody struct {
}

// 更新多个资源的部分字段
type PatchQuery struct {
	biz.QueryStr
}
type PatchBody struct {
}

// 获取一个资源
type GetIdQuery struct {
	biz.QueryStrId
}

// 获取多个资源
type GetQuery struct {
	biz.QueryStr
}
