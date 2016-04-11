package main

import (
	"fmt"
	"strconv"
)

func main() {

	edad := 30

	// int a string

	edad_str := strconv.Itoa(edad)

	fmt.Println("Mi edad es: "+edad_str)


	anio := "1986"

	// string a int

	anio_actual,_ := strconv.Atoi(anio)

	fmt.Println("Año de nacimiento: "+anio)

	fmt.Println(anio_actual)
	
	fmt.Println(anio_actual+edad)

}