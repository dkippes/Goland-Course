package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Guarda contactos en un archivo JSON
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	enconder := json.NewEncoder(file)
	err = enconder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

// Carga contactos de un archivo JSON
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Slice de contactos
	var contacts []Contact

	// Cargo los contactos desde el archivo
	err := loadContactsFromFile(&contacts)
	if err != nil {
		println("Error al cargar los contactos:", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("=== Gestor de contactos ===\n",
			"1. Agregar contacto\n",
			"2. Mostrar contactos\n",
			"3. Salir\n",
			"Opción: ")

		// Leer la opción del usuario
		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			println("Error al leer la opción:", err)
			return
		}

		switch option {
		case 1:
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Teléfono: ")
			c.Phone, _ = reader.ReadString('\n')

			contacts = append(contacts, c)

			// Guardo los contactos en el archivo
			if err := saveContactsToFile(contacts); err != nil {
				println("Error al guardar los contactos:", err)
			}
		case 2:
			fmt.Println("=== Contactos ===")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n", index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("=== Fin de la lista ===")
		case 3:
			return
		default:
			fmt.Println("Opción no válida")
		}
	}
}
