package main

import (
	"github.com/hun9k/gapi/app"
	_ "github.com/hun9k/gapiDemo/routers"
)

func main() {
	// db.DB().AutoMigrate(&schemas.Contents{})
	// 应用运行
	app.Run()
}
