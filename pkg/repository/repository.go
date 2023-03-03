package repository

import "github.com/jackc/pgx/v4/pgxpool"

type Repository struct {
	Hellow
}

type Hellow interface {
	RepositoryHellow(name string) error
}

func NewRepository(dbPool *pgxpool.Pool) *Repository {
	return &Repository{
		Hellow: NewHellowPG(dbPool),
	}
}
