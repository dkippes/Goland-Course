package main

import "fmt"

func main() {
	name := "Diego"
	age := 28

	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", name, age)

	// -----------------------------------
	var name2 string
	var age2 int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&name2)

	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&age2)

	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.\n", name2, age2)
	fmt.Println(greeting)

	fmt.Printf("El tipo de name es: %T\n", name2)
}
