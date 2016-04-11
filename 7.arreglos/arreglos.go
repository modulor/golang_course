package main

import "fmt"

func main() {
	
	arreglo := [3]int{1,3}

	// tamanio arreglo: len(arreglo)

	for x:=0; x<len(arreglo); x++ {
		fmt.Println(arreglo[x])
	}

	// bidimensional

	var matriz [2][2]int

	fmt.Println(matriz)

}