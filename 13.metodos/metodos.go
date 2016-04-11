package main

import (
	"fmt"
)

type User struct {
	edad             int
	nombre, apellido string
}

func (this *User) nombre_completo() string {
	return this.nombre + " " + " " + this.apellido
}

func (this *User) cambiar_nombre(nuevo_nombre string) {
	this.nombre = nuevo_nombre
}

func main() {

	mario := new(User)

	mario.nombre = "Mario"
	mario.apellido = "Lopez"

	fmt.Println("nombre_completo: ", mario.nombre_completo())

	// cambiar nombre

	mario.cambiar_nombre("Alberto")

	fmt.Println("cambiar_nombre: ", mario.nombre_completo())

}
