package main

import (
	"encoding/gob"
	"gatewaysvr/config"
	controller "gatewaysvr/controller"
	_ "gatewaysvr/docs"
	"gatewaysvr/log"
	"gatewaysvr/rpcservice"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yimeng436/OJ/pkg/pb"
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

// @host	127.0.0.1:8000
// @BasePath
func main() {
	gob.Register(pb.UserVo{})
	r := gin.Default()
	InitConfig()
	rpcservice.InitSvrConn()
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

	r.GET("/test/:id", controller.Test)

	user := r.Group("/user")
	{

		user.POST("/login", controller.Login)
		user.GET("/getLoginUser", controller.GetLoginUser)
	}

	question := r.Group("/question")
	{
		question.POST("/add", controller.AddQuestion)
		question.GET("/get", controller.CheckLogin, controller.GetQuestion)
		question.GET("/get/:id", controller.GetQuestionByPathId)
		question.POST("/list", controller.ListQuestion)
		question.POST("/submit/do", controller.CheckLogin, controller.DoSubmit)
		question.POST("/submit/query", controller.CheckLogin, controller.QueryQuestionSubmit)
	}

}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
