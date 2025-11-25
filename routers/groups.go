package routers

import (
	"github.com/hun9k/gapi/http"

	"github.com/gin-gonic/gin"
)

// 使用该方法，可以保证path组上的中间件应用于全部组内路由
func group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return http.Handler().Group(relativePath, handlers...)

	// ues middleware for group path
	// return gapi.RouterGroup(path, func(ctx *gin.Context) {}, func(ctx *gin.Context) {})
}
