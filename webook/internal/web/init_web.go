package web

import "github.com/gin-gonic/gin"

func RegisterRoutes() *gin.Engine {
	server := gin.Default()
	RegisterUserRoutes(server)
	return server
}
func RegisterUserRoutes(server *gin.Engine) {
	u := &UserHandler{}
	server.POST("/user/signup", u.SignUp)
	server.POST("/user/login", u.Login)
	server.POST("/user/edit", u.Edit)
	server.GET("/user/profile", u.Profile)
	server.POST("/users/logout", u.LogOut)
}
