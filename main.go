package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/jxskiss/mcli"
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
		//fmt.Printf("failed to close database connection: %w", err)
		return fmt.Errorf("failed to close database connection: %w", err)
	}
	return nil

}

func openCLI() {}

func main() {
	singleton := App{}
	host, port, user, password, dbName := getEnvVar()
	singleton.dbConn = connectToDB(host, port, user, password, dbName)
	_ = storeData("peach", "this peach is red", singleton.dbConn)
	_ = storeData("peach", "this peach is red", singleton.dbConn)
	_ = storeData("apple", "the apple is tasty", singleton.dbConn)
	_ = storeData("walnut", "this walnut is tasty", singleton.dbConn)
	_ = storeData("cherry", "this cherry is red", singleton.dbConn)
	_ = storeData("watermelon", "this watermelon is ripe", singleton.dbConn)
	_, _ = retrieveData("peach", singleton.dbConn)
	_ = removeData("peach", singleton.dbConn)
	_ = dump(singleton.dbConn)
	if err := closeDB(singleton.dbConn); err != nil {
		log.Printf("Error during database closure: %v", err)
	}
}
