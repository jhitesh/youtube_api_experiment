package postgreSQL

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DBConnection struct {
	DBConn *sql.DB
}

var DBConn DBConnection

func init() {
	readCredentials()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbCredentials.Host, dbCredentials.Port, dbCredentials.Username, dbCredentials.Password, dbCredentials.Database)

	conn, connErr := sql.Open("postgres", connectionString)
	if connErr != nil {
		log.Fatalln("Could not connect to the database. ERROR: ", connErr.Error())
		return
	}

	DBConn.DBConn = conn
}
