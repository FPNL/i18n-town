package idatabase

import "github.com/jackc/pgx/v4/pgxpool"

type IDB *pgxpool.Pool

var db *pgxpool.Pool

func Go() (err error) {
	db, err = setupPSql()
	return
}

func Connect() *pgxpool.Pool {
	if db == nil {
		panic("專案架構層級錯誤")
	}
	return db
}

func Close() {
	if db == nil {
		panic("專案架構層級錯誤")
	}

	db.Close()
}
