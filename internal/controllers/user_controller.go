package controllers

import (
	"gawean-be-go/internal/config"
	"gawean-be-go/internal/models"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	var user models.User
	config.DB.Preload("Level").First(&user, c.GetUint("user_id"))
	c.JSON(200, user)
}

func UserInfoADM(c *gin.Context) {
	var users []models.User
	config.DB.Preload("Level").Find(&users)
	c.JSON(200, users)
}
