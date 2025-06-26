package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sajid414/Go-react-auth/controller"
	"github.com/sajid414/Go-react-auth/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	private := r.Group("/private")
	private.Use(middleware.Authenticate)
	private.GET("/refreshtoken", controller.RefreshToken)

}
