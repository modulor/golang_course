package main

import "fmt"

func main() {

	// slice normal inicializado
	
	var matriz = []int{1,2,3,4}

	fmt.Println(matriz)

	// "slicing"

	arreglo := [4]int{11,22,33,44}

	slice := arreglo[:]

	fmt.Println(slice)
	
}