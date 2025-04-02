package main

import "fmt"

func main() {
	var num uint
	fmt.Println("Digite sua idade: ")
	fmt.Scan(&num) // para fazer a leiura de dados

	if num >= 18 {
		fmt.Println("Voce eh maior de idade")
	} else {
		fmt.Println("Voce eh menor de idade")
	}
}
