package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"web_app/app/other/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysServerMonitorRouter)
}

// 需认证的路由代码
func registerSysServerMonitorRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.ServerMonitor{}
	r := v1.Group("/server-monitor").Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("", api.ServerInfo)
	}
}
