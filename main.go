package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/jxskiss/mcli"
	"log"
	"os"
)

const (
	CreateTableSQL       = `CREATE TABLE IF NOT EXISTS key_value_table (key VARCHAR(255) PRIMARY KEY, value TEXT NOT NULL)`
	InsertSQL            = `INSERT INTO key_value_table (key, value) VALUES ($1, $2)`
	SelectByKeySQL       = `SELECT value FROM key_value_table WHERE key = $1`
	SelectByRandomKeySQL = `SELECT value FROM key_value_table ORDER BY RANDOM() LIMIT 1`
	DeleteSQL            = `DELETE FROM key_value_table WHERE key = $1`
	SelectRandomKeySQL   = `SELECT key FROM key_value_table ORDER BY RANDOM() LIMIT 1`
	SelectAllSQL         = `SELECT key, value FROM key_value_table`
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
	fmt.Println("Successfully connected to database!")
	return db
}

func closeDB(db *sql.DB) error { //todo -> when do I close the db connection?
	err := db.Close()
	if err != nil {
		//fmt.Printf("failed to close database connection: %w", err)
		return fmt.Errorf("failed to close database connection: %w", err)
	}
	fmt.Println("Successfully disconnected from database!")
	return nil

}

//func openCLI() {}

func main() {
	singleton := App{}
	host, port, user, password, dbName := getEnvVar()
	singleton.dbConn = connectToDB(host, port, user, password, dbName)
	/*
		// Define a helper function to handle errors
		handleError := func(operation string, err error) {
			if err != nil {
				log.Printf("Error during %s: %v", operation, err)
			}
		}

		// Use the helper with each operation
		handleError("storing peach", storeSchrodingerData("peach", "this peach is red", singleton.dbConn))
		handleError("storing apple", storeSchrodingerData("apple", "the apple is tasty", singleton.dbConn))
		handleError("storing walnut", storeSchrodingerData("walnut", "this walnut is tasty", singleton.dbConn))
		handleError("storing cherry", storeSchrodingerData("cherry", "this cherry is red", singleton.dbConn))
		handleError("storing watermelon", storeSchrodingerData("watermelon", "this watermelon is ripe", singleton.dbConn))
	*/
	err := storeSchrodingerData("peach", "this peach is red", singleton.dbConn)
	//err, value := retrieveSchrodingerData(singleton.dbConn, "peach")
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("Success")
	}
	/*
		handleError("removing peach", removeSchrodingerData(singleton.dbConn, "peach"))
		handleError("dumping database", dump(singleton.dbConn))
	*/
	if err := closeDB(singleton.dbConn); err != nil {
		log.Printf("Error during database closure: %v", err)
	}
}
