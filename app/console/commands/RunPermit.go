package commands

import (
	"fmt"
	"log"
	"mygo/app/helpers"
	"mygo/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RunPermit() {

	// dsn := *konstanta.Connection + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := *helpers.Connection
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	user := models.User{}
	if err = helpers.TruncateTable(db, user.TableName()); err != nil {
		log.Fatal("Failed to truncate table", err)
	}

	log.Println("Table truncated successfully")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")
	// Auto Migrate (similar to Laravel migrations)
	// db.AutoMigrate(&User{})

	// Insert user
	// now := time.Now().Format("2006-01-02 15:04:05")
	// password := "12345678"
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	fmt.Println("Error hashing password:", err)
	// 	return
	// }
	// db.Create(&IsUsers{Username: "Chandra", Password: string(hashedPassword), IsActive: true, CreatedAt: now, UpdatedAt: now})

	// // Query user
	// var user User
	// db.First(&user, 1) // Find user with ID 1
	// fmt.Println("User:", user.Name)
}
