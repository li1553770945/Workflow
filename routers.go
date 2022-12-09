package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if app.Config.HttpConfig.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("sessionId", store))
	// 告诉gin框架模板文件引用的静态文件去哪里找
	//r.Static("/static", "static")
	//// 告诉gin框架去哪里找模板文件
	//r.LoadHTMLGlob("templates/*")

	r.GET("test", app.WorkflowService.Test)

	//user
	//userRoutes := r.Group("/api")
	//{
	//
	//	userRoutes.POST("/user", controller.Register)
	//	userRoutes.GET("/user", controller.Me)
	//	userRoutes.PUT("/user", controller.UpdateUser)
	//	userRoutes.GET("/logout", controller.Logout)
	//	userRoutes.POST("/login", controller.Login)
	//	userRoutes.POST("/reset-password", controller.ResetPassword)
	//	//userRoutes.POST("/new-validation", controller.NewValidation)
	//
	//}

	return r
}
