package admin

import (
	"gapiDemo/apis/admin/content"

	"github.com/gin-gonic/gin"
	"github.com/hun9k/gapi"
)

const (
	PLATFORM = "admin"
	VERSION  = "v1"
)

var router *gin.RouterGroup

func init() {
	// group router
	router = gapi.Router().Group(PLATFORM).Group(VERSION)

	// add your modules
	content.RouterRegister(router)
}
