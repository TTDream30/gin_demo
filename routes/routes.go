package routes

import (
	"gin_demo2/controller"
	"gin_demo2/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/uploadImage", controller.UploadImage)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
