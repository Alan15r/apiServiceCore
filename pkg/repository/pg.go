package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func NewConnPool(config *Database) (*pgxpool.Pool, error) {
	cfg, _ := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.UserName,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	))

	if pool, err := pgxpool.ConnectConfig(context.Background(), cfg); err == nil {
		return pool, nil
	} else {
		return nil, err
	}
}
