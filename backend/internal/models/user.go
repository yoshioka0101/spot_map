package models

import "time"

// User はユーザー情報を定義
type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name" binding:"required"`
    Email     string    `json:"email" binding:"required,email" gorm:"unique"`
    Password  string    `json:"password" binding:"required"`
    Profile   string    `json:"profile,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
