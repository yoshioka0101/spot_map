package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/yoshioka0101/spot_map/internal/models"
	"github.com/yoshioka0101/spot_map/internal/utils"
	"gorm.io/gorm"
)

// Registerは新規ユーザー登録のエンドポイント
func Register(c *gin.Context) {
	// User構造体の値
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// パスワードをハッシュ化(utiles/password.go)
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	// データベースに保存
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
