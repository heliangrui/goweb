package router

import (
	"github.com/gin-gonic/gin"
	"go-web/controller"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, loginRouter)
}

func loginRouter(r *gin.RouterGroup) {
	// 分组模块
	apiGroup := r.Group("/supaiot/api/tenant")
	// 设备路由 对应设备controller
	login := controller.LoginController{}
	// 具体api和 对应设备controller相关方法
	apiGroup.POST("/userName/login", login.Login)

	apiGroup.POST("/logout", login.Logout)
}
