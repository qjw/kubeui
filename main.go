package main

import (
	"github.com/qjw/kelly"
	"github.com/qjw/kelly/middleware"
	R "github.com/qjw/kubeui/router"
	"github.com/qjw/kubeui/service"
)

//go:generate go-bindata-assetfs -o bindata.go -pkg router ./frontend/...

func main() {
	// 方便前端调试，开启cors
	router := kelly.New(middleware.Cors(&middleware.CorsConfig{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type"},
	}))

	api := service.CreateCliApi()
	R.Init(router, api)
	router.Run(":8888")
}
