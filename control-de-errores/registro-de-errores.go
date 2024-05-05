package main

import (
	"log"
	"os"
)

func main() {
	// log.Fatal("Este es un mensaje de error fatal") // Se detiene la ejecución del programa
	// log.Panic("Este es un mensaje de error de pánico") // Se detiene la ejecución del programa
	log.SetPrefix("main(): ")
	log.Print("Este es un mensaje de log")
	log.Println("Este es un mensaje de log con salto de línea")

	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Este es un mensaje de log en un archivo")
}
