package router

import (
	"net/http"

	"github.com/qjw/kelly"
	"github.com/qjw/kubeui/service"
)

func initPod(r kelly.Router, api service.KubeApi) {
	//----pod-------------------------------------------------------------------
	podRouter := r.Group("/api/v1/pods")
	podRouter.GET("/:namespace", func(c *kelly.Context) {
		namespace := c.MustGetPathVarible("namespace")
		d, err := api.GetPods(namespace)
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
