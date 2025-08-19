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
	defer db.Close()
	//Verify connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}

/*
Basic Key-Value Store:
- Put(key, value) → Stores data.
- Get(key) → Retrieves data.
- Delete(key) → Removes data.
*/

func storeData(key string, value string, db *sql.DB) {
	//PUT
}

func retrieveData(key string, db *sql.DB) {
	//GET
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
