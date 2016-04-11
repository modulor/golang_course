package main

import(
	"fmt"
	"bufio"
	"os"
)

func main() {
	nombre := "Mario"
	edad := 29
	bandera := true
	precio := 100.50

	fmt.Printf("nombre %v\n",nombre)	
	fmt.Printf("edad %d\n",edad)
	fmt.Printf("bandera %t\n",bandera)
	fmt.Printf("precio %f\n",precio)

	var edad_usuario int

	fmt.Printf("\nTu edad: ")
	fmt.Scanf("%d\n",&edad_usuario)
	fmt.Printf("mi edad es %d\n",edad_usuario)

	var nombre_usuario string

	fmt.Printf("\nTu nombre: ")
	fmt.Scanf("%v\n",&nombre_usuario)
	fmt.Printf("mi nombre es %v\n",nombre_usuario)

	// con bufio

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingresa tu nombre: ")
	
	text,err := reader.ReadString('\n')
	
	if err != nil{

		fmt.Println(err)

	}else{

		fmt.Println("Hola "+text)

	}


}