package handler

import (
	"net/http"
	"strings"

	"gawean-be-go/internal/config"
	"gawean-be-go/internal/models"
	"gawean-be-go/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	// 🔥 CEK DULU
	var existingUser models.UserModel
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Email already registered",
		})
		return
	}

	var level models.LevelModel
	config.DB.Where("level_code = ?", "USR").First(&level)

	user := models.UserModel{
		Name:     req.Name,
		Email:    req.Email,
		Password: utils.HashPassword(req.Password),
		LevelID:  level.LevelID,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register success",
	})
}

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	c.ShouldBindJSON(&req)

	var user models.UserModel
	result := config.DB.Preload("Level").Where("email = ?", req.Email).First(&user)
	if result.Error != nil || !utils.CheckPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}

	token, _ := utils.GenerateToken(user.UserID, user.Level.LevelCode)
	c.JSON(http.StatusOK, gin.H{"token": token,
		"user": gin.H{
			"id":    user.UserID,
			"name":  user.Name,
			"email": user.Email,
			"level": user.Level.LevelName,
		},
	})
}
