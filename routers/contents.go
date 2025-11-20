package apis

import (
	"github.com/hun9k/gapiDemo/internal/contents"
)

func init() {
	// Rest API
	group := routerGroup("v1").Group("contents")
	group.POST("", contents.Post)                       // 创建一个资源
	group.DELETE(":id", contents.DeleteId)              // 删除一个资源
	group.DELETE("", contents.Delete)                   // 删除多个资源
	group.PATCH(":id/restore", contents.PatchIdRestore) // 恢复一个资源
	group.PATCH("restore", contents.PatchRestore)       // 恢复多个资源
	group.PUT(":id", contents.PutId)                    // 更新一个资源
	group.PUT("", contents.Put)                         // 更新多个资源
	group.PATCH(":id", contents.PatchId)                // 更新一个资源的部分字段
	group.PATCH("", contents.Patch)                     // 更新多个资源的部分字段
	group.GET(":id", contents.GetId)                    // 获取一个资源
	group.GET("", contents.Get)                         // 获取多个资源
}
