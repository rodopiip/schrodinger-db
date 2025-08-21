package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Object struct {
	Key   string `db:"key"`
	Value string `db:"value"`
}

type App struct {
	dbConn *sql.DB
}

func getEnvVar() (string, string, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB_NAME")
	return host, port, user, password, dbName
}

func connectToDB(host string, port string, user string, password string, dbName string) *sql.DB {
	psqInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqInfo)
	if err != nil {
		panic(err)
	}
	//Verify connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}

func closeDB(db *sql.DB) error { //todo -> when do I close the db connection?
	err := db.Close()
	if err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}
	return nil

}

func storeData(key string, value string, db *sql.DB) error {
	//PUT
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS key_value_table (
		    key VARCHAR(255) PRIMARY KEY, 
		    value TEXT NOT NULL
		                                           )
		                                           `)
	if err != nil {
		return fmt.Errorf("error creating table: %s", err)
	}
	//todo -> do I need to check if an identical key already exists in the database,
	//		or is it already handled by the "PRIMARY" key word in the SQL statement?
	_, err = db.Exec(`INSERT INTO key_value_table (key, value) VALUES ($1, $2)`, key, value)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" { //todo -> solve wrapped errors issue
			return fmt.Errorf("duplicate key: %s", key)
		}
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
	if err == sql.ErrNoRows { //todo -> solve "Comparison with errors using equality operators fails on wrapped errors"
		return err, fmt.Sprintf("No data found for key: %s", key)
	} else if err != nil {
		return err, fmt.Sprintf("Error retrieving data: %s", err)
	}
	fmt.Println("Data retrieved successfully for key " + key + ": " + value)
	return err, value
	//todo -> test
}

func removeData(key string, db *sql.DB) error {
	//DELETE
	result, err := db.Exec("DELETE FROM key_value_table WHERE key = $1", key)
	if err != nil {
		return fmt.Errorf("error deleting data: %s", err)
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no data found for key: %s", key)
	}
	fmt.Println("Data deleted successfully for key " + key)
	return nil
	//todo -> test
}

func dump(db *sql.DB) error {
	rows, err := db.Query("SELECT key, value FROM key_value_table")
	if err != nil {
		return fmt.Errorf("error querying data: %s", err)
	}
	defer rows.Close() //todo -> do I need to handle an error?

	fmt.Println("Database contents:")
	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return fmt.Errorf("error scanning row: %s", err)
		}
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
	if err = rows.Err(); err != nil {
		return fmt.Errorf("error iterating rows: %s", err)
	}
	return nil
}

func openCLI() {}

func main() {
	singleton := App{}
	host, port, user, password, dbName := getEnvVar()
	singleton.dbConn = connectToDB(host, port, user, password, dbName)
	_ = storeData("peach", "this peach is red", singleton.dbConn)
	_ = storeData("apple", "the apple is tasty", singleton.dbConn)
	_ = storeData("walnut", "this walnut is tasty", singleton.dbConn)
	_ = storeData("cherry", "this cherry is red", singleton.dbConn)
	_ = storeData("watermelon", "this watermelon is ripe", singleton.dbConn)
	_, _ = retrieveData("peach", singleton.dbConn)
	//_ = removeData("peach", singleton.dbConn)
	_ = dump(singleton.dbConn)
	if err := closeDB(singleton.dbConn); err != nil {
		log.Printf("Error during database closure: %v", err)
	}
}
