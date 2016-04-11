package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {
	
	go mi_nombre_lento("Mario")

	// funcion anonima

	go func(){
		var wait string
		fmt.Scanln(&wait)		
	}()

	fmt.Println("esperando...")
}

func mi_nombre_lento(name string){

	letras := strings.Split(name, "")

	for _, letra := range(letras){

		time.Sleep(1000 * time.Millisecond)

		fmt.Println(letra)

	}

}