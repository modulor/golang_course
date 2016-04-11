package main

import (
	"fmt"
)

type Human struct {
	name string
}

type Tutor struct {
	Human
}

func (this Human) hablar() string {
	return "wassaaaa..."
}

func (this Tutor) hablar_despedir() string {
	return this.Human.hablar() + " bye...."
}

func main() {

	// declarar tutor

	tutor := Tutor{Human{"Mario"}}

	fmt.Println("tutor.name", tutor.name)

	// otra forma

	alberto := new(Tutor)

	alberto.name = "Beto"

	fmt.Println("alberto.name: ", alberto.name)

	fmt.Println("alberto habla: ", alberto.hablar())

	fmt.Println("alberto saluda y despidete: ", alberto.hablar_despedir())

}
