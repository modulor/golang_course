package main

import (
	"fmt"
)

func main() {

	// copia el minimo de elementos en ambos arreglos

	slice := []int{99, 2, 3, 4}
	copia := make([]int, 1)

	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)

	// copia la misma cantidad de elementos

	arreglo1 := []int{22, 28}
	arreglo2 := []int{99, 11}

	copy(arreglo2, arreglo1)

	fmt.Println(arreglo1)

	fmt.Println(arreglo2)

	// pasando como parametro la longitud de otro slice

	hola := []int{4, 3, 2}
	mundo := make([]int, len(hola))

	copy(mundo, hola)

	fmt.Println("mundo: ", mundo)
}
