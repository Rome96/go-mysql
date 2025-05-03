package main

import (
	"bufio"
	"fmt"
	"go-mysql/controllers"
	"go-mysql/database"
	"go-mysql/models"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql" // el _ se le dice que se usara el paquete de forma indirecta
)

func main() {

	// new_contact := models.Contact{
	// 	Name:  "Pepito2",
	// 	Email: "pepito2@gmail.com",
	// 	Phone: "4444",
	// }

	// update_contact := models.Contact{
	// 	Id:    8,
	// 	Name:  "juanchito",
	// 	Email: "juanchito@gmail.com",
	// 	Phone: "302884884",
	// }

	//conncetion to DB
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	// controllers.ListContact(db)
	// controllers.GetContactById(db, 3)
	// controllers.CrearContact(db, new_contact)
	// controllers.UpdateContact(db, update_contact)
	// controllers.DeleteContact(db, 7)

	defer db.Close() // se ejecuta de ultimo con defer

	//MENU

	for {
		fmt.Println("\nMenú:")
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener contacto por ID")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		// Leer la opción seleccionada por el usuario
		var option int
		fmt.Scanln(&option)

		// Ejecutar la opción seleccionada
		switch option {
		case 1:
			controllers.ListContact(db)
		case 2:
			fmt.Print("Ingrese el ID del contacto: ")
			var idContact int
			fmt.Scanln(&idContact)
			controllers.GetContactById(db, 3)
		case 3:
			new_contact := inputContactDetails()
			controllers.CrearContact(db, new_contact)
		case 4:
			update_contact := inputContactDetails()
			controllers.UpdateContact(db, update_contact)
		case 5:
			fmt.Print("Ingrese el ID del contacto que quiere eliminar: ")
			var idContact int
			fmt.Scanln(&idContact)
			controllers.DeleteContact(db, idContact)
		case 6:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
		}
	}

}

// Función para ingresar los detalles del contacto desde la entrada estándar
func inputContactDetails() models.Contact {
	// Leer la entrada del usuario utilizando bufio
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	fmt.Print("Ingrese el nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el correo electrónico del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el número de teléfono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
