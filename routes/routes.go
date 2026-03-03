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

	userRepo := repository.NewUserRepository(config.DB)
	refreshRepo := repository.NewRefreshTokenRepository(config.DB)

	authService := service.NewAuthService(userRepo, refreshRepo)
	authHandler := handler.NewAuthHandler(authService)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.RefreshToken)
		auth.POST("/logout", middlewares.JWTAuth(), authHandler.Logout)
	}

	// ===== USER ROUTES =====
	user := r.Group("/user", middlewares.JWTAuth())
	{
		user.GET("/info", userHandler.UserInfo)
		user.GET("/info-adm", middlewares.OnlyADM(), userHandler.UserInfoADM)
	}
}