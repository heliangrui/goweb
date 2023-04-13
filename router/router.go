package router

import (
	ginswaggerknife "gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife"
	"github.com/gin-gonic/gin"
	"go-web/core/config"
	"go-web/core/logs"
	"go-web/core/response"
	"os"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup), 0)
)

func Init() (r *gin.Engine) {
	r = gin.New()

	r.Use(logs.GinLogger, logs.GinRecovery(true))
	knifeInit(r)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	examplesNoCheckRoleRouter(r)
	examplesCheckRoleRouter(r)
	//路由不存在
	r.NoRoute(notFindFunc)
	return
}
func knifeInit(r *gin.Engine) {
	file, err := os.ReadFile("./docs/swagger.json")
	if err == nil {
		ginswaggerknife.InitSwaggerKnife(r, string(file))
	}
}

// 无需认证的路由示例
func examplesNoCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("")
	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

// 需要认证的路由示例
func examplesCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("")
	for _, f := range routerCheckRole {
		f(v1)
	}
}

func Start(r *gin.Engine) {
	server := config.GetServer()

	r.Run(":" + server.Port)
}

func notFindFunc(c *gin.Context) {
	response.Error(c, "接口不存在！")
}
