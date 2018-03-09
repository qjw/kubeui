package router

import (
	"encoding/base64"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/qjw/kelly"
	"github.com/qjw/kubeui/config"
	"github.com/qjw/kubeui/service"
)

type Base64Param struct {
	Data string `json:"data" binding:"required" error:"base64编码内容必填"`
}

func Init(r kelly.Router, api service.KubeApi) {
	r.GET("/", func(c *kelly.Context) {
		c.Redirect(http.StatusFound, "/frontend")
	})
	r.GET("/frontend/*path", func() func(*kelly.Context) {
		fs := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: ""}
		h := http.FileServer(fs)
		return func(c *kelly.Context) {
			h.ServeHTTP(c, c.Request())
		}
	}())

	// 绑定所有的options请求来支持中间件作跨域处理
	r.OPTIONS("/*path", func(c *kelly.Context) {
		c.WriteString(http.StatusOK, "ok")
	})

	r.GET("/api/v1/menus", func() func(*kelly.Context) {
		data := config.GetMenuString()
		return func(c *kelly.Context) {
			c.WriteRawJson(http.StatusOK, data)
		}
	}())

	r.POST("/api/v1/base64",
		kelly.BindMiddleware(func() interface{} { return &Base64Param{} }),
		func(c *kelly.Context) {
			param := c.GetBindParameter().(*Base64Param)
			imgBase64Str, err := base64.StdEncoding.DecodeString(param.Data)
			if err != nil {
				panic(err)
			}
			c.WriteIndentedJson(http.StatusOK, kelly.H{
				"code":    "0",
				"message": "ok",
				"data":    string(imgBase64Str),
			})
		},
	)

	initSecret(r, api)
	initDeployment(r, api)
	initPod(r, api)
	initNamespace(r, api)
}
