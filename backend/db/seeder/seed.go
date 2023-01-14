package seeder

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func Seed(db *sql.DB) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Password@123"), bcrypt.DefaultCost)
	_, err := db.Exec("INSERT INTO users (email, password, role, nama_depan, nama_belakang) VALUES (?, ?, ?, ?, ?)", "inggris3@letitbe.com", string(hashedPassword), "mentor", "Bahasa", "inggris3")
	if err != nil {
		panic(err)
	}

}
