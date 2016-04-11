package main

import (
	"fmt"
)

type User struct {
	edad             int
	nombre, apellido string
}

func main() {

	// valores por default

	var persona User
	fmt.Println(persona)

	// con valores definidos

	fmt.Println(User{edad: 11, nombre: "Mario", apellido: "Lopez"})

	// otra forma

	alberto := User{edad: 22, nombre: "Alberto", apellido: "Perez"}

	fmt.Println("Alberto: ", alberto)

	// asignando valores segun el orden en el que estan las propiedades

	noemi := User{22, "Noemi", "Barahona"}

	fmt.Println("Noemi: ", noemi)

	modulor := new(User)

	fmt.Println("New (direccion): ", modulor)

	modulor.nombre = "YESS"

	fmt.Println("Valor (nombre): ", modulor.nombre)

	modulor.nombre = "QUE ONDA"

	fmt.Println("Cambiar valor: ", modulor.nombre)

}
