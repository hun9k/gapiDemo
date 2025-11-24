package contents

import (
	"github.com/hun9k/gapi/biz"
)

// 创建资源
type PostBody struct {
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
