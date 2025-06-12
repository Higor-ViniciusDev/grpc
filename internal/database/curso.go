package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Curso struct {
	DB          *sql.DB
	ID          string
	Nome        string
	Descricao   string
	CategoriaID string
}

func NovoCurso(db *sql.DB) *Curso {
	return &Curso{
		DB: db,
	}
}

func (c *Curso) Create(nome, descricao string, categoriaID string) (*Curso, error) {
	id := uuid.New().String()

	query := "INSERT INTO cursos (id, nome, descricao, categoria_id) VALUES ($1, $2, $3, $4)"
	_, err := c.DB.Exec(query, id, nome, descricao, categoriaID)
	if err != nil {
		return nil, err
	}

	return &Curso{
		ID:          id, // This should be replaced with the actual ID from the database
		Nome:        nome,
		Descricao:   descricao,
		CategoriaID: categoriaID,
	}, nil
}

func (c *Curso) FindAll() ([]*Curso, error) {
	query := "SELECT id, nome, descricao, categoria_id FROM cursos"
	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cursos []*Curso
	for rows.Next() {
		curso := &Curso{}
		if err := rows.Scan(&curso.ID, &curso.Nome, &curso.Descricao, &curso.CategoriaID); err != nil {
			return nil, err
		}
		cursos = append(cursos, curso)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cursos, nil
}

func (c *Curso) FindByCategoriaID(categoriaID string) ([]Curso, error) {
	query := "SELECT id, nome, descricao, categoria_id FROM cursos WHERE categoria_id = $1"
	rows, err := c.DB.Query(query, categoriaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cursos []Curso
	for rows.Next() {
		curso := Curso{}
		if err := rows.Scan(&curso.ID, &curso.Nome, &curso.Descricao, &curso.CategoriaID); err != nil {
			return nil, err
		}
		cursos = append(cursos, curso)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cursos, nil
}
