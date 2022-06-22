package idatabase

import (
	"context"
	"fmt"
	"github.com/jackc/pgproto3/v2"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
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
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s",
		os.Getenv("PSQL_USERNAME"),
		os.Getenv("PSQL_PASSWORD"),
		os.Getenv("PSQL_HOST"),
		os.Getenv("PSQL_PORT"),
	)
	dbpool, err = pgxpool.Connect(context.Background(), dns)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	return
}
