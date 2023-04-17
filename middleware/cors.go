package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		cors.New(cors.Config{
			//AllowAllOrigins: true, 和下边一个意思
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"Origin"},
			ExposeHeaders: []string{"Content-Length", "Authorization"},
			//AllowCredentials: true,//是否发送cookie请求 无需配置
			//AllowOriginFunc: func(origin string) bool {
			//	return origin == "https://github.com"
			//}, //这行的意思是如果访问的是不允许跨域的话返回的是什么
			MaxAge: 12 * time.Hour, //预请求保存时间 12个小时
		})
	}
}
