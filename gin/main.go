package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() //创建一个服务器
	//路由注册
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello go")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "print post method")
	})

	//参数路由
	r.GET("/users/:name", func(c *gin.Context) {
		//获取参数路由中参数的方法
		name := c.Param("name")
		c.String(http.StatusOK, "这是参数路由"+name)
	})
	//通配符路由
	r.GET("/view/*.html", func(c *gin.Context) {
		//获取html文件名
		page := c.Param(".html")
		c.String(http.StatusOK, "这是通配符路由 "+page)
	})
	//查询参数
	r.GET("order", func(c *gin.Context) {
		//c.Query("id")
		c.String(http.StatusOK, "这是查询参数"+c.Query("id"))
	})

	r.Run(":8089") // 监听并在 0.0.0.0:8080 上启动服务 //必须在最后
}
