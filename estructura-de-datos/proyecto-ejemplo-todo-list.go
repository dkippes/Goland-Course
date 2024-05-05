package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (lt *ListaTareas) agregarTarea(t Tarea) {
	lt.tareas = append(lt.tareas, t)
}

func (lt *ListaTareas) marcarCompletado(i int) {
	lt.tareas[i].completado = true
}

func (lt *ListaTareas) editarTarea(i int, t Tarea) {
	lt.tareas[i] = t
}

func (lt *ListaTareas) eliminarTarea(i int) {
	lt.tareas = append(lt.tareas[:i], lt.tareas[i+1:]...)
}

func main() {
	lista := ListaTareas{}

	leer := bufio.NewReader(os.Stdin)
	for {
		var opcion int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir")

		fmt.Print("Opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("Nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Descripción de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada")
		case 2:
			var i int
			fmt.Println("Índice de la tarea a marcar como completada: ")
			fmt.Scanln(&i)
			lista.marcarCompletado(i)
			fmt.Println("Tarea marcada como completada")
		case 3:
			var i int
			var t Tarea
			fmt.Println("Índice de la tarea a editar: ")
			fmt.Scanln(&i)
			fmt.Println("Nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Descripción de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(i, t)
			fmt.Println("Tarea editada")
		case 4:
			var i int
			fmt.Println("Índice de la tarea a eliminar: ")
			fmt.Scanln(&i)
			lista.eliminarTarea(i)
			fmt.Println("Tarea eliminada")
		case 5:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida")
		}

		fmt.Println("Lista de tareas:")
		fmt.Println("===============")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("===============")
	}
}
