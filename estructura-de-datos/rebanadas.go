package main

import "fmt"

func main() {
	diasSemana := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}

	diaRebanada := diasSemana[0:5]
	fmt.Println(diaRebanada)

	diaRebanada = append(diaRebanada, "Viernes", "Sabado", "Otro dia")
	fmt.Println(diaRebanada)

	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)

	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	nombres := make([]string, 5, 10)
	nombres[0] = "Juan"
	nombres[1] = "Pedro"
	fmt.Println(nombres)

	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)

	copy(rebanada2, rebanada1)
	fmt.Println(rebanada1)
	fmt.Println(rebanada2)
}
