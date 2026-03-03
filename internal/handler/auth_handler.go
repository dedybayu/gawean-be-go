package handler

import (
	"net/http"
	"strings"
	"time"

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
	result := config.DB.Preload("Level").
		Where("email = ?", req.Email).
		First(&user)

	if result.Error != nil || !utils.CheckPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}

	accessToken, _ := utils.GenerateAccessToken(user.UserID, user.Level.LevelCode)
	refreshToken, _ := utils.GenerateRefreshToken(user.UserID)

	// Simpan refresh token ke database
	refresh := models.RefreshToken{
		UserID:    user.UserID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	config.DB.Create(&refresh)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": gin.H{
			"id":    user.UserID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Level.LevelCode,
		},
	})
}

func RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// cek refresh token di database
	var storedToken models.RefreshToken
	err := config.DB.
		Where("token = ? AND revoked = ?", req.RefreshToken, false).
		First(&storedToken).Error

	if err != nil || time.Now().After(storedToken.ExpiresAt) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid refresh token"})
		return
	}

	// parse token
	token, claims, err := utils.ParseToken(req.RefreshToken)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid refresh token"})
		return
	}

	userID := uint(claims["user_id"].(float64))

	// ambil role user lagi
	var user models.UserModel
	config.DB.Preload("Level").First(&user, userID)

	// generate access token baru
	newAccessToken, _ := utils.GenerateAccessToken(userID, user.Level.LevelCode)

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}

func Logout(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	if req.RefreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Refresh token required",
		})
		return
	}

	// Revoke refresh token
	result := config.DB.Model(&models.RefreshToken{}).
		Where("token = ?", req.RefreshToken).
		Update("revoked", true)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid refresh token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout success",
	})
}
