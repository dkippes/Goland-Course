package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	var p Persona
	p.nombre = "Juan"
	p.edad = 30
	p.correo = "juan@gmail.com"

	p1 := Persona{nombre: "Juan", edad: 30, correo: "juan@gmail.com"}

	fmt.Println(p)
	fmt.Println(p1)
}
