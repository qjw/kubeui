package router

import (
	"net/http"

	"github.com/qjw/kelly"
	"github.com/qjw/kubeui/service"
)

func initNamespace(r kelly.Router, api service.KubeApi) {
	//-----namespace-------------------------------------------------------
	r.GET("/api/v1/namespaces", func(c *kelly.Context) {
		d, err := api.GetNamespaces()
		if err != nil {
			c.Abort(500, err.Error())
			return
		}
		c.WriteIndentedJson(http.StatusOK, kelly.H{
			"code":    "0",
			"message": "ok",
			"data":    d,
		})
	})
}
