package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// data source name properties
	dsn := mysql.Config{
		User:                 "root",
		Passwd:               "deepak",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "sakila",
		AllowNativePasswords: true,
	}

	// Get a database handle
	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("connected")

	actorID, err := addActor("JOE", "BERRY")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID of the Actor: %v\n", actorID)

}

func addActor(firstName string, lastname string) (int64, error) {
	// return 1, nil
	result, err := db.Exec("insert into actor(first_name, last_name) values(?, ?)", firstName, lastname)
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}

	return id, nil
}
