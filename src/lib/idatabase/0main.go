package idatabase

import (
	"fmt"
	"sync"
)

type database struct {
	Mux   sync.RWMutex
	Words WordsTable
}

var once sync.Once
var db *database

type Option struct{}

func Go(option Option) error {
	fmt.Println("Setup db...")
	defer fmt.Println("db setup...")
	once.Do(func() {
		db = &database{
			Words: WordsTable{
				Columns: make(Columns, 0, 10),
			},
		}
	})

	return nil
}

func ConnectDB() *database {
	return db
}
