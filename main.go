package main

import (
	"go-mysql/controllers"
	"go-mysql/database"
	"go-mysql/models"
	"log"

	_ "github.com/go-sql-driver/mysql" // el _ se le dice que se usara el paquete de forma indirecta
)

func main() {

	// new_contact := models.Contact{
	// 	Name:  "Pepito2",
	// 	Email: "pepito2@gmail.com",
	// 	Phone: "4444",
	// }

	update_contact := models.Contact{
		Id:    8,
		Name:  "juanchito",
		Email: "juanchito@gmail.com",
		Phone: "302884884",
	}

	//conncetion to DB
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	controllers.ListContact(db)
	controllers.GetContactById(db, 3)
	// controllers.CrearContact(db, new_contact)
	controllers.UpdateContact(db, update_contact)

	defer db.Close() // se ejecuta de ultimo con defer

}
