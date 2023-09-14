package middleware

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(path ...string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path...)
	return l
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	//用go的方式解析元数据，编码，解码
	gob.Register(time.Now())
	return func(c *gin.Context) {
		//不需要校验
		for _, path := range l.paths {
			if c.Request.URL.Path == path {
				return
			}
		}
		//if c.Request.URL.Path == "/users/login" ||
		//	c.Request.URL.Path == "/users/signup" {
		//	return
		//}

		sess := sessions.Default(c)
		//表示session的查询语句，看是基于什么实现的，基于redis实现的session就表示查寻redis
		id := sess.Get("userId")
		if id == nil {
			//没有登录
			c.String(http.StatusOK, "eeeeeeeeeeeeeeee")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//需要上边代码先判断通过，表示用户登录，采取获取上次的更新时间
		update_time := sess.Get("update_time")
		sess.Set("userId", id)
		//通过回去cookie中seesion信息判断maxage的时间
		//ck, err := c.Request.Cookie("ssid")
		//ck.MaxAge
		sess.Options(sessions.Options{
			MaxAge: 60,
		})
		now := time.Now()
		//update_time为nil表示还没有刷新过
		if update_time == nil {
			sess.Set("update_time", now)
			if err := sess.Save(); err != nil {
				panic(err)
			}
		}
		//如果updatetime存在
		updateTimeVal, _ := update_time.(time.Time)
		//一分钟刷新一次
		if now.Sub(updateTimeVal) > time.Second*10 {
			sess.Set("update_time", now)
			if err := sess.Save(); err != nil {
				panic(err)
			}
		}
	}
}

/*
	update_time := sess.Get("update_time")


		now := time.Now().UnixMilli()
		//update_time为nil表示还没有刷新过
		if update_time == nil {
			sess.Set("update_time", now)
			sess.Options(sessions.Options{
				MaxAge: 30,
			})
			sess.Save()
			return
		}
		//如果updatetime存在
		updateTimeVal, _ := update_time.(int64)
		if now-updateTimeVal > 30*1000 {
			sess.Set("update_time", now)
			sess.Options(sessions.Options{
				MaxAge: 30,
			})
			sess.Save()
			return
		}
*/
