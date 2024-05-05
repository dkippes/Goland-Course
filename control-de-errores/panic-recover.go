package main

import "fmt"

func divide(dividendo, divisor int) {
	defer func() { // No interfiere con el flujo de la funcion
		if r := recover(); r != nil { // Recupera el error, se genera un panic
			fmt.Println(r)
		}
	}()
	fmt.Println(dividendo / divisor)
}

func main() {
	divide(10, 2)
	divide(10, 0)
	divide(10, 5)
}
