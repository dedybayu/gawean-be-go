package routes

import (
	"gawean-be-go/internal/controllers"
	"gawean-be-go/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	user := r.Group("/user", middlewares.JWTAuth())
	{
		user.GET("/info", controllers.UserInfo)
		user.GET("/info-adm", middlewares.OnlyADM(), controllers.UserInfoADM)
	}
}
