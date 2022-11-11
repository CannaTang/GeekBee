package api

import (
	"GeekBee5/api/middleware"
	"GeekBee5/dao"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	go dao.LoadAll()
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/register", register)   // 注册
	r.POST("/login", login)         // 登录
	r.POST("/modify", modifyUser)   // 修改
	r.POST("/find", FindPassword)   // 找回密码
	r.POST("/Comment", MakeComment) // 评论

	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}

	r.Run(":8088") // 跑在 8088 端口上
}
