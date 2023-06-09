package main

import "fmt"

var n int

// FUNÇÃO INIT SERÁ SEMPRE EXECUTADA ANTES DA FUNÇÃO MAIN.
// AO CONTRÁRIO DA FUNÇÃO MAIN QUE SÓ PODE TER UMA POR PASTA, A FUNÇÃO INIT PODE TER UMA POR ARQUIVO!
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}
