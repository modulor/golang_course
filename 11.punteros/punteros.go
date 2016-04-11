package main

import (
	"fmt"
)

func main() {
	var x, y *int
	entero := 5
	x = &entero
	y = &entero

	// imprimir direccion de memoria

	fmt.Println("X direccion de memoria: ", x)
	fmt.Println("Y direccion de memoria: ", y)

	// imprimir el valor

	fmt.Println("X valor: ", *x)
	fmt.Println("Y valor: ", *y)

	// alterar el contenido de la variable

	*x = 6

	fmt.Println("Y ahora vale: ", *y)
}
