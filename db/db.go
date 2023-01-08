package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5454
	user     = "abubakirkais"
	password = ""
	dbname   = "brokerage"
)

func EstablishDBConnection() {

	// prepare the DB connection string
	postgresConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// sql.open() validates the arguments BUT doesnt make an actual connection to the db
	// NOTE: sql.Open() function call never creates a connection to the database. Instead, it simply validates the arguments provided.
	db, err := sql.Open("postgres", postgresConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Ping() opens a connection to the DB
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Established a successful DataBase connection!")
}
