package router

import (
	"github.com/bigbigliu/web_app/app/other/apis"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/tools/transfer"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysServerMonitorRouter)
}

func registerSysServerMonitorRouter(v1 *gin.RouterGroup) {
	api := apis.ServerMonitor{}
	r := v1.Group("/server-monitor")
	{
		r.GET("", api.ServerInfo)
	}
	v1.GET("/metrics", transfer.Handler(promhttp.Handler()))
}
