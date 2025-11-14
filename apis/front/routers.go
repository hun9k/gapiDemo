package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hun9k/gapi"
)

const (
	PLATFORM = "front"
	VERSION  = "v1"
)

var router *gin.RouterGroup

func init() {
	// group router
	router = gapi.Router().Group(PLATFORM).Group(VERSION)
}
