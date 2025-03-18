package configs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/driver/mysql"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var SQL *gorm.DB

func init() {
	getDir, _ := os.Getwd()
	if strings.TrimSpace(filepath.Base(getDir)) == "configs" {
		viper.SetConfigFile("../.env")
	} else {
		viper.SetConfigFile(".env")
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := PgSQLConnection(); err != nil {
		log.Fatal(err)
	}
	// if err := MySQLConnection(); err != nil {
	// 	log.Fatal(err)
	// }
}

func MySQLConnection() error {
	var err error
	dsn := viper.GetString("DB_USERNAME") + ":" + viper.GetString("DB_PASSWORD") + "@tcp(" + viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT") + ")/" + viper.GetString("DB_DATABASE")
	SQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func PgSQLConnection() error {
	// Define your PostgreSQL database connection settings
	dsn := "host=" + viper.GetString("DB_HOST") + " user=" + viper.GetString("DB_USERNAME") + " password=" + viper.GetString("DB_PASSWORD") + " dbname=" + viper.GetString("DB_DATABASE") + " port=" + viper.GetString("DB_PORT") + " sslmode=disable"

	// Open a connection to the PostgreSQL database
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sql, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return err
	}
	SQL = sql
	return nil
	// defer db.Close()
}
