package main

import "fmt"

func main() {
	// ARITMÉTICOS (+, -, /, *, %, ++, --)
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 20 / 4
	multilicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multilicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)

	// FIM DOS ARITMÉTICOS

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// FIM DOS OPERADORES DE ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 == 1)
	fmt.Println(1 != 4)
	fmt.Println(6 > 3)
	fmt.Println(2 < 4)
	fmt.Println(2 >= 2)
	fmt.Println(6 <= 5)

	// FIM DOS RELACIONAIS

	// OPERADORES LÓGICOS &&, || e !
	fmt.Println("----------------------")
	verdadeiro, falso := true, false

	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// FIM DOS OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--     // numero - 1
	numero -= 20 // numero = numero - 20
	numero *= 3  // numero = numero * 3
	numero /= 2  // numero = numero / 2
	fmt.Println(numero)

	// FIM DOS OPERADORES UNÁRIOS
}
