package repository

import (
	"database/sql"
	"fmt"

	"github.com/BoruTamena/internal/core/port/repository"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	dbname   = "user_db"
	user     = "postgres"
	password = "root"
)

type database struct {
	*sql.DB
}

func NewDB() (repository.DataBase, error) {

	dns := fmt.Sprintf("host=%v port=%d user=%v password=%v dbname=%v sslmode=disable", host, port, user, password, dbname)

	con, err := sql.Open("postgres", dns)

	if err != nil {

		return nil, err
	}

	return database{
		con,
	}, nil
}

func (db database) GetDB() *sql.DB {
	return db.DB
}

func (db database) Close() error {
	return db.DB.Close()
}
