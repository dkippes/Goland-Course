package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es", p.nombre)
}

func main() {
	var x int = 10
	fmt.Println(x)
	editar(&x) // Si no pasamos el puntero se hace una copia de la memoria pero no se modifica la variable en si.. por lo tanto no cambia
	fmt.Println(x)

	p := Persona{nombre: "Juan", edad: 30, correo: "juan@gmail.com"}
	p.diHola() // Es un metodo de la estructura Persona
}

func editar(x *int) {
	*x = 20
}
