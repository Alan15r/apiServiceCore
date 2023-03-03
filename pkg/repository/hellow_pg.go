package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type HellowPG struct {
	db *pgxpool.Pool
}

func NewHellowPG(db *pgxpool.Pool) *HellowPG {
	return &HellowPG{db: db}
}

func (hpg *HellowPG) RepositoryHellow(name string) error {
	return nil
}
