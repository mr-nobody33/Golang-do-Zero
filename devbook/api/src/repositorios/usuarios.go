package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Insere um usuário no banco de dados
func (u usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
