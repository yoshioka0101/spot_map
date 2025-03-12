package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yoshioka0101/spot_map/internal/handler"
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

	// 汎用ハンドラのエンドポイント
	r.GET("/", handler.HelloWorldHandler)

	r.POST("/register", handler.Register)

	return r
}
