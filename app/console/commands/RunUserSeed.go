package commands

import (
	"fmt"
	"mygo/app/console/helpers"
	"mygo/app/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RunUserSeed() {

	// dsn := *helpers.Connection + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := *helpers.Connection
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")

	// Auto Migrate (similar to Laravel migrations)
	// db.AutoMigrate(&User{})

	// Insert user
	now := time.Now()
	// now := time.Now().Format(helpers.DateTimeUS)

	// fmt.Println("Error hashing password:", now)

	password := "12345678"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	db.Create(&models.User{Username: "Chandra", Password: string(hashedPassword), IsActive: true, CreatedAt: now, UpdatedAt: now})

	// // Query user
	// var user User
	// db.First(&user, 1) // Find user with ID 1
	// fmt.Println("User:", user.Name)
}
