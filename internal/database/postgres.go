package database

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func OpenPostgres(dbString string) *sqlx.DB {
	db, err := sqlx.Open("pgx", dbString)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}
