package main

import (
	"fmt"
	"rsc.io/quote"
)

// Declaracion de constantes
const Pi = 3.14

const (
	Domingo = iota + 1 // Inicia desde 0
	Lunes
	Martes
	Miecoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Hello())

	// Declaracion de variables
	var firstName, email = "Diego", "kippes.diego@gmail.com"
	var (
		lastName = "Kippes"
		age      int
	)

	fechaNacimiento := "05/10/1995" // Solo dentro de las funciones se puede

	firstName = "Diego"
	lastName = "Kippes"
	age = 28

	fmt.Println(firstName, lastName, age, email, fechaNacimiento)
	fmt.Println(Pi, Domingo)
}
