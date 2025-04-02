package main

import "fmt"

const Pi float64 = 3.14159 //constantes tipadas(float64) e exportadas(pode usar fora do pacote e comeca com letra maiscula)
const Linguagem string = "GO"

func main() {
	var x int = 10
	var y float64 = 7.89
	var sum = float64(x) + (y) //cast
	fmt.Println(sum)

	fmt.Println("---------------")

	num1 := 3
	num2 := 2
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)

	fmt.Println("----------------")
	fmt.Println("CONSTANTES:")

	fmt.Println("PI =", Pi)
	fmt.Println("Linguagem =", Linguagem)

}
