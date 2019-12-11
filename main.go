package main

import (
	"fmt"
	"github.com/nameouyang/learning-go/conf"
	_ "github.com/nameouyang/learning-go/conf"
	_ "github.com/nameouyang/learning-go/lib/redis"
	_ "github.com/nameouyang/learning-go/models"
	"github.com/nameouyang/learning-go/router"
)

func main() {
	engine := router.InitRouter()
	_ = engine.Run(fmt.Sprintf(":%d", conf.ServerConf.Port))
	//fmt.Println(conf.ServerConf)
	/*
		// 创建一个默认的路由引擎
		r := gin.Default()
		// GET：请求方式；/hello：请求的路径
		// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
		r.GET("/hello", func(c *gin.Context) {
			// c.JSON：返回JSON格式的数据
			c.JSON(200, gin.H{
				"message": "Hello world!",
			})
		})
		// 启动HTTP服务，默认在0.0.0.0:8080启动服务
		_ = r.Run()*/
}
