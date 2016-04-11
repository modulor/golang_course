package main

import "fmt"

func main() {
	
	for x:=0; x<10; x++ {
		fmt.Println("hola ",x)
	}

	i := 1
	for i<10 {
		if i > 10{
			break
		}
		fmt.Println(i)
		i++
	}

}