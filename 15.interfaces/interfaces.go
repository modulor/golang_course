package main

import (
	"fmt"
)

type User interface {
	Permisos() int
	Nombre() string
}

// administrador

type Admin struct {
	nombre string
}

func (this Admin) Permisos() int {
	return 5
}

func (this Admin) Nombre() string {
	return this.nombre
}

// editor

type Editor struct {
	nombre string
}

func (this Editor) Permisos() int {
	return 3
}

func (this Editor) Nombre() string {
	return this.nombre
}

func auth(user User) string {

	permisos := user.Permisos()

	if permisos > 4 {
		return user.Nombre() + " tiene permisos de administrador"
	} else if permisos == 3 {
		return user.Nombre() + " tiene permisos de editor"
	}

	return "no tiene permisos..."
}

func main() {
	admin := Admin{"Mario"}
	editor := Editor{"Alberto"}
	fmt.Println(auth(admin))
	fmt.Println(auth(editor))

	// como un arreglo

	usuarios := []User{Admin{"Noemi"}, Editor{"Barahona"}}

	for _, usuario := range usuarios {
		fmt.Println(auth(usuario))
	}
}
