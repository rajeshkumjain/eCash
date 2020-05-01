package infra

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

// DB : Database object
type DB struct {
	*sql.DB
}

// NewDB : initial database connection and create the DB object
func NewDB() (*DB, error) {
	flag.Parse()
	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
		fmt.Printf(" database:%s\n", *database)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", *server, *user, *password, *database)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	db, err := sql.Open("sqlserver", connString)

	db.SetConnMaxLifetime(1)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.Stats()


	if err != nil {
		log.Fatal("*** ERROR *** \n Failed to connect to Database: ", err.Error())
	}
	return &DB{db}, nil
}
