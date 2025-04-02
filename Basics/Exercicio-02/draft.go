package main

import (
	"fmt"
	"unicode/utf8" //pois caracteres acentuados ocupam mais bytes no UTF-8.
)

func main() {

	fmt.Println("INTEIROS:")

	var a int = 10
	fmt.Println(a)

	var x uint = 100 //uint numeor sosmente positivo
	fmt.Println(x)

	z := 350
	fmt.Println(z)

	fmt.Println("-----------------")
	fmt.Println("PONTOS FLUTUANTES:")

	var b float64 = 5.5
	fmt.Println(b)

	var c = 5.9 //por padrao vai em float64
	fmt.Println(c)

	preco := 19.99 // o Go infere que 'preco' é um float64
	fmt.Println(preco)

	fmt.Println("-----------------")
	fmt.Println("STRINGS:")

	var nome = "Cauan"
	fmt.Println(nome)

	var name string = "Clara"
	fmt.Println(name)

	nick := "Thebirl009"
	fmt.Println(nick)

	palavra := "Cauan"
	fmt.Println(len(palavra)) //len conta quantos bytes tem a palara - 1 byte = 1 palavra

	palavraAcentuada := "Olá"
	fmt.Println(utf8.RuneCountInString(palavraAcentuada)) //conta os caracteres reais(acentos,etc)

	fmt.Println("-----------------")
	fmt.Println("BOOLEANOS:")

	var isTrue bool = true
	fmt.Println(isTrue)

	isOn := true
	fmt.Println(isOn)
}

/*Aqui estão os principais tipos inteiros em Go:

int8: 8 bits, valores de -128 a 127.

int16: 16 bits, valores de -32,768 a 32,767.

int32: 32 bits, valores de -2,147,483,648 a 2,147,483,647.

int64: 64 bits, valores de -9,223,372,036,854,775,808 a 9,223,372,036,854,775,80

--------------------------
E para os tipos sem sinal (não negativos):

uint8: 8 bits, valores de 0 a 255.

uint16: 16 bits, valores de 0 a 65,535.

uint32: 32 bits, valores de 0 a 4,294,967,295.

uint64: 64 bits, valores de 0 a 18,446,744,073,709,551,615.

------------------

Resumo:

float64 é mais preciso e tem uma faixa de valores maior do que float32.

float32 ocupa menos memória, mas é menos preciso.
*/
