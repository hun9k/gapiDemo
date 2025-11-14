package content

import (
	"github.com/gin-gonic/gin"
)

func RouterRegister(router *gin.RouterGroup) {
	// Rest API
	content := router.Group("contents")
	content.POST("", Post)                       // 创建一个资源
	content.DELETE(":id", DeleteId)              // 删除一个资源
	content.DELETE("", Delete)                   // 删除多个资源
	content.PATCH(":id/restore", PatchIdRestore) // 恢复一个资源
	content.PATCH("restore", PatchRestore)       // 恢复多个资源
	content.PUT(":id", PutId)                    // 更新一个资源
	content.PUT("", Put)                         // 更新多个资源
	content.PATCH(":id", PatchId)                // 更新一个资源的部分字段
	content.PATCH("", Patch)                     // 更新多个资源的部分字段
	content.GET(":id", GetId)                    // 获取一个资源
	content.GET("", Get)                         // 获取多个资源
}
