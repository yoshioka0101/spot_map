package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// グローバルな`gorm.DB`インスタンス
var DB *gorm.DB

// InitDBはMySQLに接続し`gorm.DB`をセットする
func InitDB() {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := os.Getenv("MYSQL_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", user, pw, dbHost, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB open error: %v", err)
	}

	// 接続確認（リトライあり）
	if err = checkConnect(5); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	fmt.Println("Database connected successfully")
}

// checkConnect は DB 接続確認（最大 `count` 回リトライ）
func checkConnect(count int) error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	for count > 0 {
		if err := sqlDB.Ping(); err != nil {
			log.Printf("DB connection failed: %v. Retrying...", err)
			time.Sleep(2 * time.Second)
			count--
		} else {
			return nil
		}
	}
	return fmt.Errorf("failed to connect to DB after multiple retries")
}

// GetDBは`gorm.DB`を取得する
func GetDB() *gorm.DB {
	return DB
}
