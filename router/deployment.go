package router

import (
	"net/http"

	"github.com/qjw/kelly"
	"github.com/qjw/kubeui/service"
)

func initDeployment(r kelly.Router, api service.KubeApi) {
	//----deployment------------------------------------------------------------------
	deployRouter := r.Group("/api/v1/deployments")
	deployRouter.GET("/:namespace", func(c *kelly.Context) {
		namespace := c.MustGetPathVarible("namespace")
		d, err := api.GetDeployments(namespace)
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

	deployRouter.GET("/:namespace/:id", func(c *kelly.Context) {
		namespace := c.MustGetPathVarible("namespace")
		id := c.MustGetPathVarible("id")
		c.WriteIndentedJson(http.StatusOK, kelly.H{
			"code":      "0",
			"message":   "ok",
			"id":        id,
			"namespace": namespace,
		})
	})
}
