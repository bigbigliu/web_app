package router

import (
	"github.com/go-admin-team/go-admin-core/sdk/config"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/ws"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"

	"github.com/bigbigliu/web_app/common/middleware/handler"
	_ "github.com/bigbigliu/web_app/docs/admin"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")

	sysBaseRouter(g)

	// swagger；注意：生产环境可以注释掉
	if config.ApplicationConfig.Mode != "prod" {
		sysSwaggerRouter(g)
	}

	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)
		// Refresh time can be longer than token timeout
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	// 需要认证
	sysCheckRoleRouterInit(g, authMiddleware)
	return g
}

func sysBaseRouter(r *gin.RouterGroup) {

	go ws.WebsocketManager.Start()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendAllService()

	if config.ApplicationConfig.Mode != "prod" {
		r.GET("/prod-ping", handler.Ping)
	}
	r.GET("/info", handler.Ping)
}

func sysSwaggerRouter(r *gin.RouterGroup) {
	r.GET("/swagger/admin/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("admin")))
}

func sysCheckRoleRouterInit(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	wss := r.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		wss.GET("/ws/:id/:channel", ws.WebsocketManager.WsClient)
		wss.GET("/wslogout/:id/:channel", ws.WebsocketManager.UnWsClient)
	}
}
