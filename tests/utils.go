package tests

import "github.com/gistsapp/pogo/pogo"

func CreateDatabase() *pogo.Database {
	return pogo.NewDatabase("postgres", "postgres", "0.0.0.0", "5432", "pogo")
}

func ClearDatabase(db *pogo.Database) {
	db.Exec("DELETE FROM users")
}

func SeedDatabase(db *pogo.Database) {
	db.Exec("INSERT INTO users(name) VALUES('test1')")
	db.Exec("INSERT INTO users(name) VALUES('test2')")
	db.Exec("INSERT INTO users(name) VALUES('test3')")
}
