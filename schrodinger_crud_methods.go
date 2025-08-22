package main

import (
	"database/sql"
	"fmt"
)

func storeSchrodingerData(key string, value string, db *sql.DB, randomGenerator RandomNumbersGenerator) error {
	if randomGenerator.Random() > 0.3 {
		fmt.Println("...Data is attempted to be stored...")
		return storeData(key, value, db)
	}
	fmt.Println("...Data has not been stored...")
	return nil
}

func retrieveSchrodingerData(db *sql.DB, randomGenerator RandomNumbersGenerator, args ...interface{}) (error, string) {
	if randomGenerator.Random() > 0.3 {
		fmt.Println("...Non-random data is attempted to be retrieved...")
		return retrieveData(db, SelectByKeySQL, args...)
	} else {
		fmt.Println("...Random data is attempted to be retrieved...")
		return retrieveData(db, SelectByRandomKeySQL)
	}
}

// remove data
func removeSchrodingerData(db *sql.DB, key string, randomGenerator RandomNumbersGenerator) error {
	if randomGenerator.Random() > 0.3 {
		fmt.Println("...Non-random data is attempted to be deleted...")
		return removeData(db, key)
	} else {
		fmt.Println("...Random data is attempted to be deleted...")
		randomKey, _ := getRandomKey(db)
		return removeData(db, randomKey)
	}
}
