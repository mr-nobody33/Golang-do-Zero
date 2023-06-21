package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndreco := enderecos.TipoDeEndereco("Rodovia dos Imigrantes")
	fmt.Println(tipoEndreco)
}
