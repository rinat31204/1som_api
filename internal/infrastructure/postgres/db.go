package postgres

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	dsn := os.Getenv("postgres://som_user:som_password_pwd@postgres:5432/som?sslmode=disable")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
