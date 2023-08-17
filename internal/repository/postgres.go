package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

type Config struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	SslMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	os.Unsetenv("PGLOCALEDIR")
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s database=%s user=%s password=%s sslmode=%s", cfg.Host, cfg.Port,
		cfg.Database, cfg.User, cfg.Password, cfg.SslMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
