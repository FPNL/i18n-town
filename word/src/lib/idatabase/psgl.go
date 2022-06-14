package idatabase

import (
	"context"
	"fmt"
	"github.com/jackc/pgproto3/v2"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	E_PSQL_URI = "postgres://postgres:root@i18n_idb_1:5432"
)

func PrintResult(rows pgx.Rows) {
	PrintColumnsName(rows.FieldDescriptions())
	PrintValue(rows)
}

func PrintColumnsName(columns []pgproto3.FieldDescription) {
	for i, description := range columns {
		if i == len(columns)-1 {
			fmt.Println(string(description.Name))
		} else {
			fmt.Printf("%v ", string(description.Name))
		}
	}
}

func PrintValue(rows pgx.Rows) {
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return
		}
		fmt.Println(values)
	}
}

func setupPSql() (dbpool *pgxpool.Pool, err error) {
	// urlExample := os.Getenv("DATABASE_URL")
	dbpool, err = pgxpool.Connect(context.Background(), E_PSQL_URI)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	return
}
