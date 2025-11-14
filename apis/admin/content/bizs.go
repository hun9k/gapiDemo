package content

// 创建一个资源
type ReqPost struct {
}
type RespPost struct {
}

// 删除一个资源
type ReqDeleteId struct {
}
type RespDeleteId struct {
}

// 删除多个资源
type ReqDelete struct {
}
type RespDelete struct {
}

// 恢复一个资源
type ReqPatchIdRestore struct {
}
type RespPatchIdRestore struct {
}

// 恢复多个资源
type ReqPatchRestore struct {
}
type RespPatchRestore struct {
}

// 更新一个资源
type ReqPutId struct {
}
type RespPutId struct {
}

// 更新多个资源
type ReqPut struct {
}
type RespPut struct {
}

// 更新一个资源的部分字段
type ReqPatchId struct {
}
type RespPatchId struct {
}

// 更新多个资源的部分字段
type ReqPatch struct {
}
type RespPatch struct {
}

// 获取一个资源
type ReqGetId struct {
}
type RespGetId struct {
}

// 获取多个资源
type ReqGet struct {
}
type RespGet struct {
}
