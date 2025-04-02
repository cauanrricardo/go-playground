package main

import "fmt"

func main() {
	fmt.Println("IF ELSE")
	if 1 == 1 && 2 == 2 { //&& ambos tem q ser true
		fmt.Println("Passed the check")
	}

	if 12 == 10 || 12 == 12 { //|| pelo menos um tem q ser true
		fmt.Println("Passed the check")
	}
	y := 19
	if y >= 21 {
		fmt.Println("Apto para dirigir")
	} else {
		fmt.Println("Inapto")
	}
	fmt.Println("--------------")
	fmt.Println("SWITCH CASE")

	x := 3
	switch x {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("tres")
	default:
		fmt.Println("Nao encontrado")
	}
}
