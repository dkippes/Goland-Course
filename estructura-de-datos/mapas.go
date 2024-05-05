package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "#ff0000",
		"verde": "#4bf745",
		"azul":  "#0000ff",
	}

	fmt.Println(colors["rojo"])
	colors["blanco"] = "#ffffff"

	fmt.Println(colors)
	if valor, ok := colors["rojo"]; ok {
		fmt.Println(valor)
	} else {
		fmt.Println("No existe")
	}

	delete(colors, "rojo")

	for clave, valor := range colors {
		fmt.Printf("Clave: %s, Valor: %s\n", clave, valor)
	}
}
