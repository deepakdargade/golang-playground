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

	fmt.Println("Connected!")

	rowsDeleted, err := deleteActor(201)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total actor deleted: ", rowsDeleted)

}

func deleteActor(actorID int64) (int64, error) {
	result, err := db.Exec("delete from actor where actor_id = ?", actorID)
	if err != nil {
		fmt.Println("resutl error")
		return 0, fmt.Errorf("Error: ", err)
	}

	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Error: ", err)
	}

	return id, nil
}
