package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	dsn := mysql.Config{
		User:                 "root",
		Passwd:               "deepak",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "sakila",
		AllowNativePasswords: true,
	}

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
	fmt.Println("connected!")

	rowsUpdated, err := UpdateActor("JAMES", 201)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total actors updated: ", rowsUpdated)

}

func UpdateActor(first_name string, actorID int64) (int64, error) {
	result, err := db.Exec("update actor set first_name = ? where actor_id = ?", first_name, actorID)
	if err != nil {
		return 0, fmt.Errorf("updateActor: ", actorID)
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("updateActor: ", actorID)
	}
	return id, nil
}
