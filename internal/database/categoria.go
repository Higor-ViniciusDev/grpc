package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Categoria struct {
	db        *sql.DB
	ID        string
	Nome      string
	Descricao string
}

func NewCategoria(db *sql.DB) *Categoria {
	return &Categoria{
		db: db,
	}
}

func (c *Categoria) Create(nome, descricao string) (Categoria, error) {
	id := uuid.New().String()

	query := "INSERT INTO categorias (id, nome, descricao) VALUES ($1, $2, $3)"
	_, err := c.db.Exec(query, id, nome, descricao)

	if err != nil {
		return Categoria{}, err
	}

	return Categoria{
		ID:        id,
		Nome:      nome,
		Descricao: descricao,
	}, nil
}

func (c *Categoria) FindAll() ([]Categoria, error) {
	query := "SELECT id, nome, descricao FROM categorias"
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorias []Categoria
	for rows.Next() {
		var categoria Categoria
		if err := rows.Scan(&categoria.ID, &categoria.Nome, &categoria.Descricao); err != nil {
			return nil, err
		}
		categorias = append(categorias, categoria)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categorias, nil

}

func (c *Categoria) FindByCursoID(id string) (Categoria, error) {
	query := "SELECT id, nome, descricao FROM categorias WHERE id = (SELECT categoria_id from cursos where ID = $1)"
	row := c.db.QueryRow(query, id)

	var categoria Categoria
	if err := row.Scan(&categoria.ID, &categoria.Nome, &categoria.Descricao); err != nil {
		if err == sql.ErrNoRows {
			return Categoria{}, nil // No category found
		}
		return Categoria{}, err
	}

	return categoria, nil
}
