package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/joho/godotenv"
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

func closeDB(db *sql.DB) { //todo -> when do I close the db connection?
	db.Close()
}

/*
Basic Key-Value Store:
- Put(key, value) → Stores data.
- Get(key) → Retrieves data.
- Delete(key) → Removes data.
*/

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
		return fmt.Errorf("error inserting data: %s", err)
	}
	return nil
}

func retrieveData(key string, db *sql.DB) (error, string) {
	//GET
	_, err := db.Query("SELECT value FROM key_value_table WHERE key = $1", key)
	return err, "string"
}

func removeData(key string, db *sql.DB) {
	//DELETE
}

func openCLI() {}

func main() {
	singleton := App{}
	host, port, user, password, dbName := getEnvVar()
	singleton.dbConn = connectToDB(host, port, user, password, dbName)

}
