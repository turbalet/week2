package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"se07.com/pkg/models"
)

type SnippetModel struct {
	DB *pgxpool.Pool
}

const (
	snippetInsert = "INSERT INTO snippets (title, content, created, expires) VALUES($1, $2, $3,$4) RETURNING id"
)

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
