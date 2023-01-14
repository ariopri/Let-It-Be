package sqlite3

import (
	"context"
	"database/sql"

	"github.com/ariopri/Let-It-Be/tree/main/backend/models"
)

type kursus struct {
	db *sql.DB
}

func NewKursusRepo(db *sql.DB) models.KursusRepository {
	return &kursus{db: db}
}

func (k *kursus) GetAll(ctx context.Context) ([]models.KursusResp, error) {
	sqlStmt := `SELECT * FROM kursus`

	rows, err := k.db.QueryContext(ctx, sqlStmt)
	if err != nil {
		return []models.KursusResp{}, err
	}

	listKursus := []models.KursusResp{}

	for rows.Next() {
		kursus := models.KursusResp{}

		err := rows.Scan(&kursus.ID, &kursus.Nama, &kursus.User.ID, &kursus.Deskripsi, &kursus.Modul, &kursus.CreatedAt)
		if err != nil {
			return []models.KursusResp{}, err
		}

		listKursus = append(listKursus, kursus)
	}

	return listKursus, nil
}

func (k *kursus) GetById(ctx context.Context, id int64) (models.KursusResp, error) {
	sqlStmt := `SELECT * FROM kursus WHERE id = ?`

	row := k.db.QueryRowContext(ctx, sqlStmt, id)

	kursus := models.KursusResp{}

	err := row.Scan(&kursus.ID, &kursus.Nama, &kursus.User.ID, &kursus.Deskripsi, &kursus.Modul, &kursus.CreatedAt)
	if err != nil {
		return models.KursusResp{}, err
	}

	return kursus, nil
}

func (k *kursus) Store(ctx context.Context, kursus *models.Kursus) (models.KursusResp, error) {
	sqlStmt := `INSERT INTO kursus(nama, users_id, deskripsi, modul) VALUES(?, ?, ?, ?)`

	row, err := k.db.ExecContext(ctx, sqlStmt, kursus.Nama, kursus.UsersId, kursus.Deskripsi, kursus.Modul)
	if err != nil {
		return models.KursusResp{}, err
	}

	id, _ := row.LastInsertId()

	res, err := k.GetById(ctx, id)
	if err != nil {
		return models.KursusResp{}, err
	}

	return res, nil
}

func (k *kursus) Update(ctx context.Context, id int64, kursus *models.Kursus) (models.KursusResp, error) {
	sqlStmt := `UPDATE kursus SET nama = ?, users_id = ?, deskripsi = ?, modul = ? WHERE id = ?`

	_, err := k.db.ExecContext(ctx, sqlStmt, kursus.Nama, kursus.UsersId, kursus.Deskripsi, kursus.Modul, id)
	if err != nil {
		return models.KursusResp{}, err
	}

	res, err := k.GetById(ctx, id)
	if err != nil {
		return models.KursusResp{}, err
	}

	return res, nil
}

func (k *kursus) Delete(ctx context.Context, id int64) error {
	_, err := k.GetById(ctx, id)
	if err != nil {
		return err
	}

	sqlStmt := `DELETE FROM kursus WHERE id = ?`

	_, err = k.db.ExecContext(ctx, sqlStmt, id)
	if err != nil {
		return err
	}

	return nil
}
