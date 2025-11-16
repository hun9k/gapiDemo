package content

import (
	"gapiDemo/apis/_commons"
	"gapiDemo/apis/_schemas"
)

// 创建一个资源
type ReqPost struct {
	_schemas.Contents
}
type RespPost struct {
	_commons.Resp
}

// 删除一个资源
type ReqDeleteId struct {
	_commons.ReqId
}
type RespDeleteId struct {
	_commons.Resp
}

// 删除多个资源
type ReqDelete struct {
	_commons.Req
}
type RespDelete struct {
	_commons.Resp
}

// 恢复一个资源
type ReqPatchIdRestore struct {
	_commons.ReqId
}
type RespPatchIdRestore struct {
	_commons.Resp
}

// 恢复多个资源
type ReqPatchRestore struct {
	_commons.Req
}
type RespPatchRestore struct {
	_commons.Resp
}

// 更新一个资源
type ReqPutId struct {
	_commons.ReqId
}
type RespPutId struct {
	_commons.Resp
}

// 更新多个资源
type ReqPut struct {
	_commons.Req
}
type RespPut struct {
	_commons.Resp
}

// 更新一个资源的部分字段
type ReqPatchId struct {
	_commons.ReqId
}
type RespPatchId struct {
	_commons.Resp
}

// 更新多个资源的部分字段
type ReqPatch struct {
	_commons.Req
}
type RespPatch struct {
	_commons.Resp
}

// 获取一个资源
type ReqGetId struct {
	_commons.ReqId
}
type RespGetId struct {
	_commons.Resp
}

// 获取多个资源
type ReqGet struct {
	_commons.Req
}
type RespGet struct {
	_commons.Resp
}
