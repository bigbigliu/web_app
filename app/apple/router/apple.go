package router

import (
	"github.com/gin-gonic/gin"

	"github.com/bigbigliu/web_app/app/apple/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerAppleRouter)
}

// registerAppleRouter
func registerAppleRouter(v1 *gin.RouterGroup) {
	api := apis.Apple{}
	r := v1.Group("/apple")
	{
		r.GET("/list", api.GetPage)
		r.GET("/detail/:id", api.Get)
		r.POST("/create", api.Insert)
		r.PUT("/edit", api.Update)
		r.DELETE("", api.Delete)
	}
}