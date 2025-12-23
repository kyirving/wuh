package middleware

import (
	"basketball/internal/conf"
	"basketball/pkg/auth"
	"basketball/pkg/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT(conf *conf.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// apiPaths := []string{"/user/login", "/user/register"}
		// for _, path := range apiPaths {
		// 	if c.Request.URL.Path == path {
		// 		c.Next()
		// 		return
		// 	}
		// }
		// 从请求头中获取token
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			util.Output(c, http.StatusUnauthorized, gin.H{
				"message": "未授权",
			}, "未授权")
			c.Abort()
			return
		}

		// 提取token
		token := bearerToken[len("Bearer "):]
		fmt.Println(token)
		// 解析token
		jwtProvider := auth.NewJWTProvider([]byte(conf.Manager.JWTSecret), conf.Manager.JWTSecret, time.Duration(conf.Manager.AccessExpire)*time.Second, time.Duration(conf.Manager.RefreshExpire)*time.Second)
		claims, err := jwtProvider.ParseToken(token)
		fmt.Println(claims)
		println(err)
		if err != nil {
			util.Output(c, http.StatusUnauthorized, gin.H{
				"message": "未授权",
			}, "未授权")
			c.Abort()
			return
		}

		fmt.Println(claims)

		c.Set("uid", claims["sub"].(float64))
		c.Set("roles", claims["roles"].([]interface{}))
		c.Next()
	}
}
