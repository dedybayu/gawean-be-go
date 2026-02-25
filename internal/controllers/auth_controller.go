package controllers

import (
	"net/http"
	"strings"

	"gawean-be-go/internal/config"
	"gawean-be-go/internal/models"
	"gawean-be-go/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req struct {
		Nama     string `json:"nama"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	// 🔥 CEK DULU
	var existingUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Email already registered",
		})
		return
	}

	var level models.Level
	config.DB.Where("kode = ?", "USR").First(&level)

	user := models.User{
		Nama:     req.Nama,
		Email:    req.Email,
		Password: utils.HashPassword(req.Password),
		LevelID:  level.ID,
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

	var user models.User
	result := config.DB.Preload("Level").Where("email = ?", req.Email).First(&user)
	if result.Error != nil || !utils.CheckPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Level.Kode)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
