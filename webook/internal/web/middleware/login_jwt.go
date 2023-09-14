package middleware

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
}

func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path ...string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path...)
	return l
}

func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
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
		//设置jwt的校验
		//jwt放在sesion里面
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//生成的token是bearer xxxx所以需要切片处理一下
		segs := strings.Split(tokenHeader, " ")
		if len(segs) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//获取token的字符串
		tokenstr := segs[1]
		//对获取到的字符串进行解码，拿到真正的token
		token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
			return []byte("MM653MID5HDZ3LLAG57SB294YBHS76UU"), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//err为nil，token不为nil
		if token == nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
