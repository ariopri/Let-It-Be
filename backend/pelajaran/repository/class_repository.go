package repository

import (
	"context"
	"database/sql"

	"github.com/ariopri/Let-It-Be/tree/main/backend/models"
)

type ClassRepo struct {
	DB *sql.DB
}

func NewClassRepo(DB *sql.DB) models.ClassRepository {
	return &ClassRepo{DB: DB}
}

func (c *ClassRepo) GetAll(ctx context.Context) ([]*models.Class, error) {
	query := "SELECT id, subject_id, teacher, rating FROM classes"
	rows, err := c.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []*models.Class
	for rows.Next() {
		var class models.Class
		if err := rows.Scan(&class.ID, &class.SubjectID, &class.Teacher, &class.Rating); err != nil {
			return nil, err
		}
		classes = append(classes, &class)
	}
	return classes, nil
}

func (c *ClassRepo) GetByID(ctx context.Context, id int64) (*models.Class, error) {
	query := "SELECT id, subject_id, teacher, rating FROM classes WHERE id = ?"
	row := c.DB.QueryRowContext(ctx, query, id)

	var class models.Class
	if err := row.Scan(&class.ID, &class.SubjectID, &class.Teacher, &class.Rating); err != nil {
		return nil, err
	}
	return &class, nil
}

func (c *ClassRepo) GetByName(ctx context.Context, name string) (*models.Class, error) {
	query := "SELECT id, subject_id, teacher, rating FROM classes WHERE name = ?"
	row := c.DB.QueryRowContext(ctx, query, name)

	var class models.Class
	if err := row.Scan(&class.ID, &class.SubjectID, &class.Teacher, &class.Rating); err != nil {
		return nil, err
	}
	return &class, nil
}

func (c *ClassRepo) GetBySubjectID(ctx context.Context, subjectID int64) ([]*models.Class, error) {
	query := "SELECT id, subject_id, teacher, rating FROM classes WHERE subject_id = ?"
	rows, err := c.DB.QueryContext(ctx, query, subjectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []*models.Class
	for rows.Next() {
		var class models.Class
		if err := rows.Scan(&class.ID, &class.SubjectID, &class.Teacher, &class.Rating); err != nil {
			return nil, err
		}
		classes = append(classes, &class)
	}
	return classes, nil
}

func (c *ClassRepo) GetBySubjectName(ctx context.Context, subjectName string) ([]*models.Class, error) {
	query := "SELECT classes.id, classes.subject_id, classes.teacher, classes.rating FROM classes INNER JOIN subjects ON classes.subject_id = subjects.id WHERE subjects.name = ?"
	rows, err := c.DB.QueryContext(ctx, query, subjectName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []*models.Class
	for rows.Next() {
		var class models.Class
		if err := rows.Scan(&class.ID, &class.SubjectID, &class.Teacher, &class.Rating); err != nil {
			return nil, err
		}
		classes = append(classes, &class)
	}
	return classes, nil
}

func (c *ClassRepo) Create(ctx context.Context, class *models.Class) (*models.Class, error) {
	query := "INSERT INTO classes (subject_id, teacher, rating) VALUES (?, ?, ?)"
	result, err := c.DB.ExecContext(ctx, query, class.SubjectID, class.Teacher, class.Rating)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	res, err := c.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *ClassRepo) Update(ctx context.Context, id int64, class *models.Class) (*models.Class, error) {
	query := "UPDATE classes SET subject_id = ?, teacher = ?, rating = ? WHERE id = ?"
	_, err := c.DB.ExecContext(ctx, query, class.SubjectID, class.Teacher, class.Rating, class.ID)
	if err != nil {
		return nil, err
	}

	res, err := c.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ClassRepo) Delete(ctx context.Context, id int64) error {
	_, err := c.GetByID(ctx, id)
	if err != nil {
		return err
	}
	query := "DELETE FROM classes WHERE id = ?"
	_, err = c.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
