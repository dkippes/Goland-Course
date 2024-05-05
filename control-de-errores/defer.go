package main

import (
	"fmt"
	"os"
)

// Usando defer se agregan a una pila -> 1, 2, 3
func main() {
	defer fmt.Println(3) // Pospone esta ejecucion hasta que la funcion termine
	fmt.Println(1)
	fmt.Println(2)

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close() // Se ejecuta al final, antes de que termine main se ejecuta

	_, err = file.Write([]byte("Hola, Diego Kippes"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
