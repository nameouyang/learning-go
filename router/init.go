package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nameouyang/learning-go/biz"
	"github.com/nameouyang/learning-go/conf"
	"github.com/nameouyang/learning-go/controller/v1"
	_ "github.com/nameouyang/learning-go/docs"
	"github.com/nameouyang/learning-go/middleware"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

func init() {
	fmt.Println("init router")
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(conf.ServerConf.RunMode)
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  conf.CORSConf.AllowAllOrigins,
		AllowMethods:     conf.CORSConf.AllowMethods,
		AllowHeaders:     conf.CORSConf.AllowHeaders,
		ExposeHeaders:    conf.CORSConf.ExposeHeaders,
		AllowCredentials: conf.CORSConf.AllowCredentials,
		MaxAge:           conf.CORSConf.MaxAge * time.Hour,
	}))
	uploadService := biz.UploadService{}
	r.StaticFS(conf.ServerConf.UploadImagePath, http.Dir(uploadService.GetImgFullPath()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiV1 := r.Group("api/v1")
	{
		authController := new(v1.AuthController)
		apiV1.POST("/auth/signup", authController.Signup)
		// 账号登录
		apiV1.POST("/auth/signin", authController.Signin)
		userController := new(v1.UserController)
		apiV1.Use(middleware.JWTAuth())
		{
			// 账户注销
			apiV1.POST("/auth/signout", authController.Signout)
			// 查看用户信息
			apiV1.GET("/user", userController.Retrieve)
			// 修改用户名称
			apiV1.PATCH("/user/name", userController.AlterName)
			// 修改用户密码
			apiV1.PATCH("/user/pass", userController.AlterPass)
			// 修改用户头像
			apiV1.PATCH("/user/avatar", userController.AlterAvatar)

			taskController := v1.TaskController{}

			// 获取任务列表
			apiV1.GET("/task", taskController.List)
			// 新增任务
			apiV1.POST("/task", taskController.Create)
			// 获取任务详情
			apiV1.GET("/task/:taskId", taskController.Retrieve)
			// 修改任务参数
			apiV1.PUT("/task/:taskId", taskController.Update)
			// 删除任务
			apiV1.DELETE("/task/:taskId", taskController.Destroy)

		}
	}

	return r
}
