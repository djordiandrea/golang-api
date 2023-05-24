package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// var DB *gorm.DB

// func ConnectDatabase() {

// 	dsn := "root:@tcp(127.0.0.1:3306)/warehouse"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		fmt.Println("Not Connect to database")
// 		return
// 	}

// 	fmt.Println("Database connect")

// 	DB = db

// }

var DB2 *sql.DB

func ConnectDatabase2() {

	dsn := "root:@tcp(127.0.0.1:3306)/warehouse"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println("Not Connect to database")
		return
	}

	fmt.Println("Database connect")

	DB2 = db

}
