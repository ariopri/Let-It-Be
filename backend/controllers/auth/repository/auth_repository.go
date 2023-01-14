package repository

import (
	"context"
	"database/sql"

	"github.com/ariopri/Let-It-Be/tree/main/backend/models"
	"github.com/ariopri/Let-It-Be/tree/main/backend/utils/hash"
)

type authRepo struct {
	db *sql.DB
}

func NewAuthRepo(DB *sql.DB) models.AuthRepository {
	return &authRepo{db: DB}
}

func (a *authRepo) findById(ctx context.Context, id int64) (models.RegisterUser, error) {
	sqlStmt := `SELECT * FROM users WHERE id = ?`
	row := a.db.QueryRowContext(ctx, sqlStmt, id)

	var registerResp models.RegisterUser

	err := row.Scan(&registerResp.ID, &registerResp.NamaDepan, &registerResp.NamaBelakang, &registerResp.Email, &registerResp.Phone, &registerResp.Password, &registerResp.Role)
	if err != nil {
		return models.RegisterUser{}, err
	}

	return registerResp, nil
}

func (a *authRepo) Login(ctx context.Context, login *models.LoginUser) (models.User, error) {
	sqlStmt := `SELECT * FROM users WHERE email = ?`

	row := a.db.QueryRowContext(ctx, sqlStmt, login.Email, login.Password)

	var user models.User

	err := row.Scan(&user.ID, &user.NamaDepan, &user.NamaBelakang, &user.Email, &user.Phone, &user.Password, &user.Role)
	if err != nil {
		return models.User{}, err
	}

	err = hash.ValidatePassword(user.Password, login.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (a *authRepo) Register(ctx context.Context, register *models.RegisterUser) (models.RegisterUser, error) {
	sqlStmt := `INSERT INTO users(nama_depan, nama_belakang, email, no_telp, password) VALUES(?, ?, ?, ?, ?)`

	register.Password, _ = hash.GeneratePassword(register.Password)

	res, err := a.db.ExecContext(ctx, sqlStmt, register.NamaDepan, register.NamaBelakang, register.Email, register.Password)
	if err != nil {
		return models.RegisterUser{}, err
	}

	id, _ := res.LastInsertId()

	userRes, err := a.findById(ctx, id)
	if err != nil {
		return models.RegisterUser{}, err
	}

	return userRes, nil
}
