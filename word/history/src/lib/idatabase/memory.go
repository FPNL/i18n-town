package idatabase

import (
	"fmt"
	"sync"
)

type memDb struct {
	Mux sync.RWMutex
}

func setupMemory() (db *memDb, err error) {
	fmt.Println("Setup db...")
	defer fmt.Println("db setup...")

	return
}
