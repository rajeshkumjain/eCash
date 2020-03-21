package repository

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// database connection needs to come utils.go now

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "rajeshjain32", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "RJDELLLAPTOP", "the database server")
	user          = flag.String("user", "sa", "the database user")
	database      = flag.String("database", "eCash", "The name of the database")
)

// DB : Database object
type DB struct {
	*sql.DB
}

// NewDB : initial database connection and create the DB object
func NewDB() (*DB, error) {
	// func DBConn() (db *sql.DB) {
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
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	return &DB{db}, nil
}
