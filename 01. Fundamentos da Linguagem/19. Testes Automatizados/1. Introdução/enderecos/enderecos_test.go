package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "rua"},
		{"Avenida Paulista", "avenida"},
		{"Rodovia dos Imigrantes", "rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "estrada"},
		{"RUA DOS BOBOS", "rua"},
		{"AVENIDA REBOUÇAS", "avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
		}
	}
}
