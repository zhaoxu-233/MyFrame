package main

import (
	"exercise_code/webook/internal/repository"
	"exercise_code/webook/internal/repository/dao"
	"exercise_code/webook/internal/service"
	"exercise_code/webook/internal/web"
	"exercise_code/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db := initDB()
	u := inituser(db)
	server := initWebServer()

	//u := &web.UserHandler{}
	u.RegisterRoutes(server)
	server.Run(":8083")
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(func(c *gin.Context) {
		println("这是第三个middleware")
	})
	//自己注册middlerware
	server.Use(func(c *gin.Context) {
		println("这是第一个middleware")
	})
	server.Use(func(c *gin.Context) {
		println("这是第二个middleware")
	})
	//use方法（HandlerFunc）作用与全部路由
	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"http://localhost:3000/"},
		//不写的话，默认接受全部方法
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders: []string{"content-type"},
		//ExposeHeaders:    []string{"content-type"},
		//是否允许带cookie
		AllowCredentials: true,
		//接受全部本地的请求,使用这个方法就可以注释调AllowOrigins
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	//添加session
	//创建cookie并设置一个存储的位置
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))

	server.Use(middleware.NewLoginMiddlewareBuilder().IgnorePaths("/users/login", "users/signup").Build())

	return server
}

func inituser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	//创建数据库
	//初始化业务代码
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		//初始化错就就结束进行
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
