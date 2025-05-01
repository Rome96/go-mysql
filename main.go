package main

import (
	"go-mysql/database"
	"log"

	_ "github.com/go-sql-driver/mysql" // el _ se le dice que se usara el paquete de forma indirecta
)

func main() {

	//conncetion to DB
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close() // se ejecuta de ultimo con defer

}
