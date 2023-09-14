package main

import (
	"exercise_code/webook/internal/repository"
	"exercise_code/webook/internal/repository/dao"
	"exercise_code/webook/internal/service"
	"exercise_code/webook/internal/web"
	"exercise_code/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
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
		AllowHeaders: []string{"content-type", "Authorization"},
		//前端想要获取到token需要设置exposeheader参数
		ExposeHeaders: []string{"x-jwt-token"},
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
	//store := cookie.NewStore([]byte("secret"))
	//store := memstore.NewStore([]byte("MM653MID5HDZ3LLAG57SB294YBHS76UU"), []byte("XGYBI4C7GDFYKNXAKP0GAQBXLX7EBFUB"))
	/*
		第一个参数表示最大空闲链接数量
		第二个就是网络协议tcp，不太可能用upd
		第三个、四个就是连接信息和密码
		第五个、六个就是两个key
	*/
	store, err := redis.NewStore(16, "tcp", "localhost:6379", "", []byte("MM653MID5HDZ3LLAG57SB294YBHS76UU"))
	if err != nil {
		panic(err)
	}
	//mystroe := &sqlx_store.Store{}
	server.Use(sessions.Sessions("ssid", store))

	//server.Use(middleware.NewLoginMiddlewareBuilder().IgnorePaths("/users/login", "/users/signup").Build())
	server.Use(middleware.NewLoginJWTMiddlewareBuilder().IgnorePaths("/users/login", "/users/signup").Build())

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
