package repository

import (
	"context"
	"database/sql"

	"github.com/ariopri/Let-It-Be/tree/main/backend/models"
	"github.com/ariopri/Let-It-Be/tree/main/backend/utils/hash"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) models.UserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) GetAll(ctx context.Context) ([]models.User, error) {
	sqlStmt := `SELECT * FROM users`
	rows, err := u.db.QueryContext(ctx, sqlStmt)
	if err != nil {
		return nil, err
	}
	var listSiswa []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.NamaDepan, &user.NamaBelakang, &user.Phone)
		if err != nil {
			return nil, err
		}
		listSiswa = append(listSiswa, user)
	}
	return listSiswa, nil
}

func (u *userRepo) GetByRole(ctx context.Context, role string) ([]models.User, error) {
	sqlStmt := `SELECT * FROM users WHERE role = ?`

	rows, err := u.db.QueryContext(ctx, sqlStmt, role)
	if err != nil {
		return []models.User{}, err
	}
	users := []models.User{}
	for rows.Next() {
		user := models.User{}

		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.NamaDepan, &user.NamaBelakang, &user.Phone)
		if err != nil {
			return []models.User{}, err
		}
		users = append(users, user)
	}
	return users, nil

}

func (u *userRepo) GetById(ctx context.Context, id uint) (models.User, error) {
	sqlStmt := `SELECT * FROM users WHERE id = ?`
	row := u.db.QueryRowContext(ctx, sqlStmt, id)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.NamaDepan, &user.NamaBelakang, &user.Phone)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *userRepo) Store(ctx context.Context, user *models.User) (models.User, error) {
	sqlStmt := `INSERT INTO users (email, password, role, nama_depan, nama_belakang, phone) VALUES (?, ?, ?, ?, ?, ?)`
	user.Password, _ = hash.GeneratePassword(user.Password)
	res, err := u.db.ExecContext(ctx, sqlStmt, user.Email, user.Password, user.Role, user.NamaDepan, user.NamaBelakang, user.Phone)
	if err != nil {
		return models.User{}, err
	}

	id, _ := res.LastInsertId()

	userRes, err := u.GetById(ctx, uint(id))
	if err != nil {
		return models.User{}, err
	}
	return userRes, nil
}

func (u *userRepo) Update(ctx context.Context, id uint, user *models.User) (models.User, error) {
	userOld, _ := u.GetById(ctx, id)
	sqlStmt := `UPDATE users SET email = ?, password = ?, nama_depan = ?, nama_belakang = ? WHERE id = ?`

	if user.Password != userOld.Password {
		user.Password, _ = hash.GeneratePassword(user.Password)
	}

	_, err := u.db.ExecContext(ctx, sqlStmt, user.NamaDepan, user.NamaBelakang, user.Email, user.Password, id)
	if err != nil {
		return models.User{}, err
	}

	res, err := u.GetById(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}

func (u *userRepo) UpdateRole(ctx context.Context, id uint, user *models.User) (models.User, error) {
	userOld, _ := u.GetById(ctx, id)
	sqlStmt := `UPDATE users SET nama_depan = ?, nama_belakang = ?, email = ?, phonie = ?, password = ?, role = ? WHERE id = ?`

	if user.Password != userOld.Password {
		user.Password, _ = hash.GeneratePassword(user.Password)
	}
	_, err := u.db.ExecContext(ctx, sqlStmt, user.NamaDepan, user.NamaBelakang, user.Email, user.Phone, user.Password, user.Role, id)
	if err != nil {
		return models.User{}, err
	}
	res, err := u.GetById(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}
func (u *userRepo) Delete(ctx context.Context, id uint) error {
	sqlStmt := `DELETE FROM users WHERE id = ?`
	_, err := u.db.ExecContext(ctx, sqlStmt, id)
	if err != nil {
		return err
	}
	return nil
}
