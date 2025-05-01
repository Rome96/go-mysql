package controllers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

// list all contacts of the db
func ListContact(db *sql.DB) {
	//select all contacts
	query := "select * from contact"

	// run consult
	contac, err := db.Query(query)

	//manage error consult
	if err != nil {
		log.Fatal(err)
	}

	//iterate over the contacts and display them
	fmt.Println("\nLIST CONTACTS:")
	fmt.Println("---------------------------------------------------------------")

	for contac.Next() {
		//Instance of the contact model
		model_contact := models.Contact{}

		err := contac.Scan(&model_contact.Id, &model_contact.Name, &model_contact.Email, &model_contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Phone: %s\n",
			model_contact.Id, model_contact.Name, model_contact.Email, model_contact.Phone)
		fmt.Println("----------------------------------------------------------------")

	}

	defer contac.Close() // This line runs last, thanks to defer

}
