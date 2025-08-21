package main

import (
	"database/sql"
	"fmt"
	"math/rand"
)

func storeSchrodingerData(key string, value string, db *sql.DB) error {
	if rand.Float64() < 0.3 {
		fmt.Println("...Data is attempted to be stored...")
		return storeData(key, value, db)
	}
	return nil
}

func retrieveSchrodingerData(db *sql.DB, key string) (error, string) {
	if rand.Float64() > 0.3 {
		fmt.Println("...Non-random data is attempted to be retrieved...")
		return retrieveData(db, SelectByKeySQL, key)
	} else {
		fmt.Println("...Random data is attempted to be retrieved...")
		return retrieveData(db, SelectByRandomKeySQL)
	}
}

// remove data
func removeSchrodingerData(db *sql.DB, key string) error {
	if rand.Float64() > 0.3 {
		fmt.Println("...Non-random data is attempted to be deleted...")
		return removeData(db, key)
	} else {
		fmt.Println("...Random data is attempted to be deleted...")
		randomKey, _ := getRandomKey(db)
		return removeData(db, randomKey)
	}
}
