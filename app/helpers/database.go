package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var SQL *gorm.DB
var Connection *string
var SQL *gorm.DB

func init() {
	getDir, _ := os.Getwd()
	if strings.TrimSpace(filepath.Base(getDir)) == "mygo" {
		fmt.Println("up:")
		viper.SetConfigFile(".env")
	} else {
		fmt.Println("down: %s", strings.TrimSpace(filepath.Base(getDir)))

		viper.SetConfigFile("../.env")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	connStr := getConnection(
		viper.GetString("DB_CONNECTION"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_DATABASE"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
	)

	fmt.Println("connStr:", connStr)

	Connection = &connStr

	if viper.GetString("DB_CONNECTION") == "mysql" {
		sql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		sql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	}

	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return err
	}
	SQL = sql
}

func getConnection(db_conn string, db_host string, db_port string, db_database string, db_username string, db_pass string) string {
	sql_conn := ""

	fmt.Println("db_conn:", db_conn)
	fmt.Println("db_host:", db_host)
	fmt.Println("db_port:", db_port)
	fmt.Println("db_database:", db_database)
	fmt.Println("db_username:", db_username)
	fmt.Println("db_pass:", db_pass)

	if db_conn == "mysql" {
		// sql_conn = fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s", db_username, db_pass, db_host, db_port, db_database)
		sql_conn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_username, db_pass, db_host, db_port, db_database)
	} else if db_conn == "pgsql" {
		sql_conn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_username, db_pass, db_database, db_port)
	}
	return sql_conn
}

// TruncateTable truncates the specified table in the database.
func TruncateTable(db *gorm.DB, tableName string) error {
	return db.Exec("TRUNCATE TABLE `" + tableName + "`").Error
}

// func initMySql() {
// 	Connection = viper.GetString("DB_USERNAME") + ":" + viper.GetString("DB_PASSWORD") + "@tcp(" + viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT") + ")/" + viper.GetString("DB_DATABASE")
// }

// func initPgSql() {
// 	Connection = "host=" + viper.GetString("DB_HOST") + " user=" + viper.GetString("DB_USERNAME") + " password=" + viper.GetString("DB_PASSWORD") + " dbname=" + viper.GetString("DB_DATABASE") + " port=" + viper.GetString("DB_PORT") + " sslmode=disable"
// }
