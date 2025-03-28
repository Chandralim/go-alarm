package models

import "time"

type User struct {
	ID        uint      `gorm:"primarykey;" json:"id"`
	Username  string    `gorm:"size:255;not null;uniqueIndex" json:"username"`
	Password  string    `gorm:"not null;" json:"password"`
	IsActive  bool      `gorm:"not null;" json:"is_active"`
	CreatedAt time.Time `gorm:"autoCreateTime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;" json:"updated_at"`
	// Fullname   string `gorm:"size:255;not null;" json:"fullname"`
	// Username   string  	`gorm:"size:30;not null;uniqueIndex" json:"username"`
	// Position  string `gorm:"size:50;not null" json:"position"`
	// Api_Token string `gorm:"null" json:"api_token"`
}

func (User) TableName() string {
	return "users"
}
