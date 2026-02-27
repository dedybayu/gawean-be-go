package routes

import (
	"gawean-be-go/internal/config"
	"gawean-be-go/internal/handler"
	"gawean-be-go/internal/middlewares"
	"gawean-be-go/internal/repository"
	"gawean-be-go/internal/service"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {

	// 🔹 Dependency Injection
	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// ===== AUTH ROUTES =====
	auth := r.Group("/auth")
	{
		auth.POST("/register", handler.Register) // kalau belum dipisah service
		auth.POST("/login", handler.Login)
	}

	// ===== USER ROUTES =====
	user := r.Group("/user", middlewares.JWTAuth())
	{
		user.GET("/info", userHandler.UserInfo)
		user.GET("/info-adm", middlewares.OnlyADM(), userHandler.UserInfoADM)
	}
}