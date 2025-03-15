package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshioka0101/spot_map/internal/database"
	"github.com/yoshioka0101/spot_map/internal/models"
	"github.com/yoshioka0101/spot_map/internal/utils"
)

// Register は新規ユーザー登録のエンドポイント
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// パスワードをハッシュ化
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	// データベースに保存
	db := database.GetDB()
	if err := db.Create(&user).Error; err != nil {
		log.Printf("Error creating user in database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login はユーザーを認証するためのエンドポイント
func Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// データベースからユーザーを取得
	db := database.GetDB()
	var user models.User
	if err := db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		log.Printf("Error finding user in database: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// パスワードを検証
	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
		log.Printf("Invalid password for user: %s", loginData.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
