package main

import (
	"fmt"
	"sync"
)

type Database struct{}

var (
	dbInstance *Database
	once       sync.Once
)

// GetDatabaseInstance guarantees dbInstance is created only once thread-safely
func GetDatabaseInstance() *Database {
	once.Do(func() {
		fmt.Println("Initializing Database Connection Pool...")
		dbInstance = &Database{}
	})
	return dbInstance
}