package main

import (
	"gatewaysvr/config"
	"gatewaysvr/log"

	controller2 "gatewaysvr/controller"
	_ "gatewaysvr/docs"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"
)

//	@title			swaggo测试
//	@version		1.0
//	@description	swaggo测试API文档

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host	127.0.0.1:8002
// @BasePath
func main() {
	r := gin.Default()
	InitConfig()
	InintRoute(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + strconv.Itoa(config.GetGlobalConfig().SvrConfig.Port))
}

func InitConfig() {
	if err := config.Init(); err != nil {
		log.Fatalf("init config failed, err:%v\n", err)
	}
	log.InitLog()
	log.Info("log init success...")

}

func InintRoute(r *gin.Engine) {
	// 使用 cookie 中间件进行 session 管理
	store := cookie.NewStore([]byte("userSecret")) // 设置 session 的加密密钥
	r.Use(Cors()).Use(sessions.Sessions("UserSession", store))

	r.GET("/test/:id", controller2.Test)

	user := r.Group("/user")
	{

		user.POST("/login", controller2.Login)
		user.GET("/getLoginUser", controller2.GetLoginUser)
	}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
