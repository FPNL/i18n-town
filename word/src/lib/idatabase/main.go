package idatabase

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

type Model interface{}

func Go() (err error) {
	db, err = setupPSql()
	if err != nil {
		fmt.Println("database done")
	}
	return
}

func Connect() *pgxpool.Pool {
	if db == nil {
		panic("專案架構級別錯誤")
	}
	return db
}

func Close() {
	if db == nil {
		panic("專案架構級別錯誤")
	}
	db.Close()
}
