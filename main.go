package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // el _ se le dice que se usara el paquete de forma indirecta
)

func main() {
	dns := "username:password@protocol(address)/dbname?param=value" // change value

	db, err := sql.Open("mysql", dns)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("DB Connections Success")
}
