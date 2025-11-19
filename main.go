package main

import (
	_ "github.com/hun9k/gapiDemo/apis"

	"github.com/hun9k/gapi"
)

func main() {
	// 应用运行
	app := gapi.App()

	if err := app.Run(); err != nil {
		gapi.Log().Error(err.Error())
	}
}
