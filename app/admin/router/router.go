package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
)

func InitExamplesRouter(r *gin.Engine) *gin.Engine {

	// 无需认证的路由
	examplesNoCheckRoleRouter(r)

	return r
}

// 无需认证的路由示例
func examplesNoCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")
	for _, f := range routerNoCheckRole {
		f(v1)
	}
}
