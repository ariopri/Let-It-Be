package main

import (
	"database/sql"

	"github.com/ariopri/Let-It-Be/tree/main/backend/db/seeder"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:tanahdamai@tcp(localhost:3306)/letitbe")
	if err != nil {
		panic(err)
	}
	seeder.Seed(db)
	defer db.Close()
}
