package router

import (
	"github.com/gin-gonic/gin"
	v1 "others-part/api/v1"
	"others-part/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	user := r.Group("api/v1/user")
	{

		user.GET("/login", v1.LoginUser)

	}
	return r
}
