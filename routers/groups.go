package routers

import (
	"github.com/hun9k/gapi/http"
)

// 使用该方法，可以保证path组上的中间件应用于全部组内路由
func group(path string) *http.RouterGroup {
	return http.Group(path)

	// ues middleware for group path
	// return gapi.RouterGroup(path, func(ctx *gin.Context) {}, func(ctx *gin.Context) {})
}
