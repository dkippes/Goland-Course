package main

import "fmt"

func main() {
	saludo := hello("Diego")
	fmt.Println(saludo)

	sum, mult := calc(4, 5)
	fmt.Println("La suma es: ", sum)
	fmt.Println("La multiplicacion es: ", mult)
}

func hello(name string) string {
	return "Hola, " + name
}

func calc(a, b int) (sum, mult int) {
	sum = a + b
	mult = a * b
	return
}
