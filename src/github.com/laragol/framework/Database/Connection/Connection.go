package Connection

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	db2, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE")))

	if err != nil {
		println("Database Connection Error")
	} else {
		println("Database Connection Ok")
	}

	if os.Getenv("DEBUG") == "true" {
		db2.LogMode(true)
	}

	db = db2
}

func Get() *gorm.DB {
	/*db, err = gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE")))

	if os.Getenv("DEBUG") == "true" {
		db.LogMode(true)
	}*/

	return db
}
