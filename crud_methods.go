package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
)

func storeData(key string, value string, db *sql.DB) error {
	//PUT
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS key_value_table (
		    key VARCHAR(255) PRIMARY KEY, 
		    value TEXT NOT NULL
		                                           )
		                                           `)
	if err != nil {
		//fmt.Printf("error creating table: %s", err)
		return fmt.Errorf("error creating table: %s", err)
	}
	//todo -> do I need to check if an identical key already exists in the database,
	//		or is it already handled by the "PRIMARY" key word in the SQL statement? -> DONE
	_, err = db.Exec(`INSERT INTO key_value_table (key, value) VALUES ($1, $2)`, key, value)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" { //todo -> solve wrapped errors issue -> DONE
			//fmt.Printf("duplicate key: %s", key)
			return fmt.Errorf("duplicate key: %s", key)
		}
		//fmt.Printf("error inserting data: %s", err)
		return fmt.Errorf("error inserting data: %s", err)
	}
	fmt.Println("Data stored successfully for key " + key)
	return nil
	//todo -> test
}

func retrieveData(key string, db *sql.DB) (error, string) {
	//GET
	var value string
	err := db.QueryRow("SELECT value FROM key_value_table WHERE key = $1", key).Scan(&value)
	if errors.Is(err, sql.ErrNoRows) { //todo -> solve "Comparison with errors using equality operators fails on wrapped errors" -> DONE
		//fmt.Printf("No data found for key: %s", key)
		return err, fmt.Sprintf("No data found for key: %s", key)
	} else if err != nil {
		//fmt.Printf("Error retrieving data: %s", err)
		return err, fmt.Sprintf("Error retrieving data: %s", err)
	}
	fmt.Println("Data retrieved successfully for key " + key + ": " + value)
	return err, value
	//todo -> test -> DONE
}

func removeData(key string, db *sql.DB) error {
	//DELETE
	result, err := db.Exec("DELETE FROM key_value_table WHERE key = $1", key)
	if err != nil {
		//fmt.Printf("error deleting data: %s", err)
		return fmt.Errorf("error deleting data: %s", err)
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		//fmt.Printf("no data found for key: %s", key)
		return fmt.Errorf("no data found for key: %s", key)
	}
	fmt.Println("Data deleted successfully for key " + key)
	return nil
	//todo -> test -> DONE
}

func dump(db *sql.DB) error {
	rows, err := db.Query("SELECT key, value FROM key_value_table")
	if err != nil {
		//fmt.Printf("error querying data: %s", err)
		return fmt.Errorf("error querying data: %s", err)
	}
	defer rows.Close() //todo -> do I need to handle an error?

	fmt.Println("Database contents:")
	for rows.Next() { //todo -> do the rows have to be closed after iteration (closing the channel)?
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			//fmt.Printf("error scanning row: %s", err)
			return fmt.Errorf("error scanning row: %s", err)
		}
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
	if err = rows.Err(); err != nil {
		//fmt.Printf("error iterating rows: %s", err)
		return fmt.Errorf("error iterating rows: %s", err)
	}
	return nil
}
