package router

import (
	"github.com/bigbigliu/web_app/app/other/apis/tools"
	"github.com/go-admin-team/go-admin-core/sdk/config"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"

	"github.com/bigbigliu/web_app/common/middleware/handler"
	_ "github.com/bigbigliu/web_app/docs/admin"
)

func InitSysRouter(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("")

	sysBaseRouter(g)

	if config.ApplicationConfig.Mode != "prod" {
		sysSwaggerRouter(g)
	}

	gen := tools.Gen{}
	r.GET("/gen/code", gen.GenCodeNew) // 生成代码

	return g
}

func sysBaseRouter(r *gin.RouterGroup) {
	r.GET("/health", handler.Ping)
}

func sysSwaggerRouter(r *gin.RouterGroup) {
	r.GET("/swagger/admin/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("admin")))
}
