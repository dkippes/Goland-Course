package main

import (
	"fmt"
	"strconv"
)

func main() {

	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(int32(integer16) + integer32)

	s := "100"
	i, _ := strconv.Atoi(s) // El _ es un valor que no queremos guardar

	fmt.Println(i)

	n := 42
	sn := strconv.Itoa(n)
	fmt.Println(sn + sn)
}
