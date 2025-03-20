package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yoshioka0101/spot_map/internal/handler"
	"github.com/yoshioka0101/spot_map/internal/database"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// CORS 設定を追加
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// MySQL に接続
	database.InitDB()
	db := database.GetDB()

	// DB をコンテキストにセットするミドルウェア
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// エンドポイントの登録
	r.GET("/", handler.HelloWorldHandler)
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	return r
}
