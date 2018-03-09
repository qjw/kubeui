package router

import (
	"net/http"

	"github.com/qjw/kelly"
	"github.com/qjw/kubeui/service"
)

type UpdateSecretParam struct {
	Data map[string]string `json:"data" binding:"required"`
}

func initSecret(r kelly.Router, api service.KubeApi) {
	//----secret------------------------------------------------------------------
	secretRouter := r.Group("/api/v1/secrets")
	secretRouter.GET("/:namespace", func(c *kelly.Context) {
		namespace := c.MustGetPathVarible("namespace")
		d, err := api.GetSecrets(namespace)
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

	secretRouter.GET("/:namespace/:id", func(c *kelly.Context) {
		d, err := api.GetSecret(
			c.MustGetPathVarible("namespace"),
			c.MustGetPathVarible("id"),
		)
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

	secretRouter.DELETE("/:namespace/:id", func(c *kelly.Context) {
		out, err := api.DeleteSecret(
			c.MustGetPathVarible("namespace"),
			c.MustGetPathVarible("id"),
		)
		if err != nil {
			c.Abort(500, err.Error())
			return
		}
		c.WriteIndentedJson(http.StatusOK, kelly.H{
			"code":    "0",
			"message": out,
		})
	})

	secretRouter.PUT("/:namespace/:id",
		kelly.BindMiddleware(func() interface{} { return &UpdateSecretParam{} }),
		func(c *kelly.Context) {
			param := c.GetBindParameter().(*UpdateSecretParam)
			out, err := api.UpdateSecret(
				c.MustGetPathVarible("namespace"),
				c.MustGetPathVarible("id"),
				param.Data,
			)
			if err != nil {
				c.Abort(500, err.Error())
				return
			}
			c.WriteIndentedJson(http.StatusOK, kelly.H{
				"code":    "0",
				"message": out,
			})
		})
}
