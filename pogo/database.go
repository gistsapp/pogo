package pogo

import (
	"database/sql"
)

type Database struct {
	user     string
	password string
	host     string
	port     string
	database string
}

type IDatabase interface {
	Connect() (*sql.DB, error)
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)
}

func NewDatabase(pg_user string, pg_password string, pg_host string, pg_port string, pg_database string) *Database {
	return &Database{
		pg_user,
		pg_password,
		pg_host,
		pg_port,
		pg_database,
	}
}

func (db *Database) Connect() (*sql.DB, error) {
	connStr := "user=" + db.user + " password=" + db.password + " host=" + db.host + " port=" + db.port + " dbname=" + db.database + " sslmode=disable"
	return sql.Open("postgres", connStr)
}

func (db *Database) Query(query string, args ...any) (*sql.Rows, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if len(args) == 0 {
		return conn.Query(query)
	}
	return conn.Query(query, args...)
}

func (db *Database) Exec(query string, args ...any) (sql.Result, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return conn.Exec(query, args...)
}
