package router

import (
	"github.com/gin-gonic/gin"

	"github.com/bigbigliu/web_app/app/orange/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerOrangeRouter)
}

// registerOrangeRouter
func registerOrangeRouter(v1 *gin.RouterGroup) {
	api := apis.Orange{}
	r := v1.Group("/orange")
	{
		r.GET("/list", api.GetPage)
		r.GET("/detail/:id", api.Get)
		r.POST("/create", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}
