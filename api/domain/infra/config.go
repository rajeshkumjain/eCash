package infra

import "flag"

// database connection configuration

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "rajeshjain32", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "RJDELLLAPTOP", "the database server")
	user          = flag.String("user", "sa", "the database user")
	database      = flag.String("database", "eCash", "The name of the database")
)
