package main

import (
	"errors"
	"fmt"
)

func PrintHeloo() { //void
	fmt.Println("Hello World")
}
func Sum(a int, b int) int {
	return a + b
}
func IntDivision(numerador, denominador int) (int, int, error) { //as funções podem retornar múltiplos valores,
	var err error //inicializado como nil por padrao
	if denominador == 0 {
		err = errors.New(" Nao pode dividir por zero!\n")
		return 0, 0, err
	}
	var result int = (numerador / denominador)
	var rest int = (numerador % denominador)
	return result, rest, err
}

func main() {
	PrintHeloo() //basta chamar a funcao
	soma := Sum(10, 5)
	fmt.Println(soma)

	fmt.Println("-----------------")

	var a int = 90
	var b int = 45
	var result, rest, err = IntDivision(a, b)
	if err != nil { //ou seja, contém um erro)
		fmt.Printf(err.Error())
	} else if rest == 0 {
		fmt.Printf("O resultado da divisao inteira eh: %v\n", result)
	}
	fmt.Printf(" Numerador = %v\n Denominador = %d\n Resultado da divisao inteira = %d\n Resultado do resto = %d\n", a, b, result, rest)
}

/*
O que é nil em Go?

nil é o valor zero para:

Ponteiros

Interfaces

Maps

Slices

Channels

Funções

****No contexto de erros, nil significa "nenhum erro".

----------------
Use %d quando:

Você quer garantir que está formatando um inteiro

Precisa de formatação numérica específica (como %04d para zeros à esquerda)

Use %v quando:

Você quer uma saída genérica

Está trabalhando com tipos desconhecidos ou variados

Fazendo debugging
*/
