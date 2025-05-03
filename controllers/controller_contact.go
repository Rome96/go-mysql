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

	//contac.Next() se usa para moverse a la siguiente fila en los resultados de la consulta.
	for contac.Next() {
		//Instance of the contact model
		model_contact := models.Contact{}

		//validate nulls
		var valueEmails sql.NullString

		//Scan llena el model_contact con los datos de la fila actual.
		err := contac.Scan(&model_contact.Id, &model_contact.Name, &valueEmails, &model_contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		//validate nulls - return bool
		if valueEmails.Valid {
			model_contact.Email = valueEmails.String
		} else {
			model_contact.Email = "Not email"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Phone: %s\n",
			model_contact.Id, model_contact.Name, model_contact.Email, model_contact.Phone)
		fmt.Println("----------------------------------------------------------------")

	}

	defer contac.Close() // This line runs last, thanks to defer
}

// GET CONTACT BY ID
func GetContactById(db *sql.DB, contactID int) {
	query := "SELECT * FROM contact WHERE id = ?"
	row := db.QueryRow(query, contactID)
	model_contact := models.Contact{}
	var valueEmails sql.NullString

	err := row.Scan(&model_contact.Id, &model_contact.Name, &valueEmails, &model_contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("Not found contact ID -> %d", contactID)
		}
		log.Fatal(err)
	}

	fmt.Println("\nLIST OF A CONTACT:")
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Phone: %s\n",
		model_contact.Id, model_contact.Name, model_contact.Email, model_contact.Phone)
	fmt.Println("----------------------------------------------------------------")
}

// CrearContact adds the specified contac to the database,
// returning the contact ID of the new entry
func CrearContact(db *sql.DB, contact models.Contact) {
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"
	result, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)

	if err != nil {
		log.Fatalf("Create contact: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Create contact: %v", err)
	}

	fmt.Printf("ID of added Contact: %v\n", id)
}

// UPDATE CONTACT
func UpdateContact(db *sql.DB, contact models.Contact) {
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)

	if err != nil {
		log.Fatalf("Create contact: %v", err)
	}

	fmt.Printf("Updated Contact success: %v\n", contact)
}
