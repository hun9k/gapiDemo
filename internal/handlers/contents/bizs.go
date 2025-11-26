package contents

import (
	"time"

	"github.com/hun9k/gapi/biz"
	"github.com/hun9k/gapi/schema"
	"github.com/hun9k/gapiDemo/internal/models"
)

type postBody struct {
	body
}

type postRespBody struct {
	body
	schema.Model
}

// 创建资源
type body struct {
	StringField     *string    `json:"string_field" binding:""`
	NullStringField *string    `json:"null_string_field" binding:""`
	IntField        *int       `json:"int_field" binding:""`
	NullIntField    *int64     `json:"null_int_field" binding:""`
	Float64Field    *float64   `json:"float64_field" binding:""`
	NullFloatField  *float64   `json:"null_float_field" binding:""`
	BoolField       *bool      `json:"bool_field" binding:""`
	NullBoolField   *bool      `json:"null_bool_field" binding:""`
	TimeField       *time.Time `json:"time_field" binding:""`
	NullTimeField   *string    `json:"null_time_field" binding:""`
}

func postBodyToModel(pb *postBody) models.Contents {
	return models.Contents{
		StringField:     biz.PtrToType(pb.StringField),
		NullStringField: biz.PtrToNullType(pb.NullStringField),
		IntField:        biz.PtrToType(pb.IntField),
		NullIntField:    biz.PtrToNullType(pb.NullIntField),
		Float64Field:    biz.PtrToType(pb.Float64Field),
		NullFloatField:  biz.PtrToNullType(pb.NullFloatField),
		BoolField:       biz.PtrToType(pb.BoolField),
		NullBoolField:   biz.PtrToNullType(pb.NullBoolField),
		TimeField:       biz.TimePtrToTime(pb.TimeField),
		NullTimeField:   biz.StrPtrToNullTime(pb.NullTimeField),
	}
}

func modelToPostResp(m *models.Contents) postRespBody {
	body := body{
		StringField:     biz.TypeToPtrType(m.StringField),
		NullStringField: biz.NullToPtrType(m.NullStringField),
		IntField:        biz.TypeToPtrType(m.IntField),
		NullIntField:    biz.NullToPtrType(m.NullIntField),
		Float64Field:    biz.TypeToPtrType(m.Float64Field),
		NullFloatField:  biz.NullToPtrType(m.NullFloatField),
		BoolField:       biz.TypeToPtrType(m.BoolField),
		NullBoolField:   biz.NullToPtrType(m.NullBoolField),
		TimeField:       biz.TypeToPtrType(m.TimeField),
		NullTimeField:   biz.NullTimeToStrPtr(m.NullTimeField),
	}

	prb := postRespBody{}
	prb.body = body
	prb.Model = m.Model

	return prb
}
