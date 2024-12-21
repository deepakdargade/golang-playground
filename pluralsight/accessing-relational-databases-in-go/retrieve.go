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

	actors, err := GetActor(201)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Actor Found: ", actors)
}

type Actor struct {
	actor_id   int64
	first_name string
	last_name  string
}

func GetActor(actorID int64) ([]Actor, error) {
	var actors []Actor

	result, err := db.Query("select actor_id, first_name, last_name from actor where actor_id = ?", actorID)

	if err != nil {
		return nil, fmt.Errorf("GetActor %v, %v", actorID, err)
	}

	defer result.Close()

	for result.Next() {
		var actor Actor
		if err := result.Scan(&actor.actor_id, &actor.first_name, &actor.last_name); err != nil {
			return nil, fmt.Errorf("GetActor %v: %v", actorID, err)
		}
		actors = append(actors, actor)

		if err := result.Err(); err != nil {
			return nil, fmt.Errorf("GetActor %v: %v", actorID, err)
		}
	}

	return actors, nil
}
