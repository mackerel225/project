package main

import (
	"database/sql"
)

func main() {
	injectionTest("Nottingham")
}

func injectionTest(city string) {
	db, err := sql.Open("postgres", "postgresql://test:test@test")
	if err != nil {
		// return err
	}

	var count int

	row := db.QueryRow("SELECT COUNT(*) FROM t WHERE city=" + city) //nolint:safesql
	if err := row.Scan(&count); err != nil {
		// return err
	}

	row = db.QueryRow("SELECT COUNT(*) FROM t WHERE city=?", city)
	if err := row.Scan(&count); err != nil {
		// return err
	}

	return
}