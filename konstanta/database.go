package konstanta

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// var SQL *gorm.DB
var Connection *string

func init() {
	getDir, _ := os.Getwd()
	if strings.TrimSpace(filepath.Base(getDir)) == "konstanta" {
		viper.SetConfigFile("../.env")
	} else {
		viper.SetConfigFile(".env")
	}

	db_conn := viper.GetString("DB_CONNECTION")
	if db_conn == "mysql" {

	} else if db_conn == "pgsql" {

	}

	connStr := getConnection(
		viper.GetString("DB_CONNECTION"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_DATABASE"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
	)

	Connection = &connStr
}

func getConnection(db_conn string, db_host string, db_port string, db_database string, db_username string, db_pass string) string {
	sql_conn := ""

	if db_conn == "mysql" {
		sql_conn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_username, db_pass, db_host, db_port, db_database)
	} else if db_conn == "pgsql" {
		sql_conn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_username, db_pass, db_database, db_port)
	}
	return sql_conn
}

// func initMySql() {
// 	Connection = viper.GetString("DB_USERNAME") + ":" + viper.GetString("DB_PASSWORD") + "@tcp(" + viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT") + ")/" + viper.GetString("DB_DATABASE")
// }

// func initPgSql() {
// 	Connection = "host=" + viper.GetString("DB_HOST") + " user=" + viper.GetString("DB_USERNAME") + " password=" + viper.GetString("DB_PASSWORD") + " dbname=" + viper.GetString("DB_DATABASE") + " port=" + viper.GetString("DB_PORT") + " sslmode=disable"
// }
