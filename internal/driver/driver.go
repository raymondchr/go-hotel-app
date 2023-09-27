package driver

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDBConn = 10
const maxIdleDBConn = 5
const maxOpenDBConnLifetime = 5 * time.Minute

func ConnectSQL(dsn string) (*DB, error) {
	db, err := NewDatabase(dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxIdleTime(maxIdleDBConn)
	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetConnMaxLifetime(maxOpenDBConnLifetime)

	dbConn.SQL = db

	err = testDB(db)
	if err != nil {
		log.Fatal(err)
	}

	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		log.Println("cannot ping the database")
		return err
	}

	return nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	//create a new database connection
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Println("Error in opening connection to a database")
		return nil, err
	}

	//check if the database is alive
	if err = db.Ping(); err != nil {
		log.Println("Error getting response from database")
		return nil, err
	}

	return db, nil
}
