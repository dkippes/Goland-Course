package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	t := time.Now()
	hora := t.Hour()

	if hora < 12 {
		fmt.Println("Mañana")
	} else if hora < 17 {
		fmt.Println("Tarde")
	} else {
		fmt.Println("Noche")
	}

	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOs")
	default:
		fmt.Println("Go run -> Otro OS")
	}
}
