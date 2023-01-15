package repository

import (
	"context"
	"database/sql"

	"github.com/ariopri/Let-It-Be/tree/main/backend/models"
)

type SubjectRepo struct {
	DB        *sql.DB
	ClassRepo *ClassRepo
}

func NewSubjectRepo(DB *sql.DB, classRepo *ClassRepo) models.SubjectRepository {
	return &SubjectRepo{
		DB:        DB,
		ClassRepo: classRepo,
	}
}

func (s *SubjectRepo) GetAll(ctx context.Context) ([]*models.Subject, error) {
	query := "SELECT id, name, image, description FROM subjects"
	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []*models.Subject
	for rows.Next() {
		var subject models.Subject
		if err := rows.Scan(&subject.ID, &subject.Name, &subject.Image, &subject.Description); err != nil {
			return nil, err
		}

		subjects = append(subjects, &subject)
	}

	return subjects, nil
}

func (s *SubjectRepo) GetByID(ctx context.Context, id int64) (*models.Subject, error) {
	query := "SELECT id, name, image, description FROM subject WHERE id = ?"
	row := s.DB.QueryRowContext(ctx, query, id)

	var subject models.Subject
	if err := row.Scan(&subject.ID, &subject.Name, &subject.Image, &subject.Description); err != nil {
		return nil, err
	}
	return &subject, nil
}

func (s *SubjectRepo) GetByName(ctx context.Context, name string) (*models.Subject, error) {
	query := "SELECT id, name, image, description FROM subject WHERE name = ?"
	row := s.DB.QueryRowContext(ctx, query, name)

	var subject models.Subject
	if err := row.Scan(&subject.ID, &subject.Name, &subject.Image, &subject.Description); err != nil {
		return nil, err
	}
	return &subject, nil
}

func (s *SubjectRepo) GetSubjectWithClasses(ctx context.Context) ([]*models.Subject, error) {
	query := `
		SELECT 
			s.id, s.name, s.image, s.description, 
			c.id, c.subject_id, c.teacher, c.rating 
		FROM subjects s 
		LEFT JOIN classes c ON s.id = c.subject_id`

	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// map untuk menyimpan subject yang sudah diambil dari database
	subjectsMap := make(map[int64]*models.Subject)
	for rows.Next() {
		var subject models.Subject
		var class models.Class
		if err := rows.Scan(
			&subject.ID, &subject.Name, &subject.Image, &subject.Description,
			&class.ID, &class.SubjectID, &class.Teacher, &class.Rating); err != nil {
			return nil, err
		}

		// jika subject belum ada dalam map, tambahkan subject baru
		if _, ok := subjectsMap[subject.ID]; !ok {
			subjectsMap[subject.ID] = &subject
			subject.Classes = []*models.Class{&class}
		} else {
			// jika subject sudah ada dalam map, tambahkan class ke slice class dari subject tersebut
			subjectsMap[subject.ID].Classes = append(subjectsMap[subject.ID].Classes, &class)
		}
	}

	// tampung semua subject dari map ke slice
	var subjects []*models.Subject
	for _, subject := range subjectsMap {
		subjects = append(subjects, subject)
	}

	return subjects, nil
}

func (s *SubjectRepo) Create(ctx context.Context, subject *models.Subject) (*models.Subject, error) {
	query := "INSERT INTO subjects (name, image, description) VALUES (?,?,?)"
	result, err := s.DB.ExecContext(ctx, query, subject.Name, subject.Image, subject.Description)
	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	subject.ID = lastID

	for _, class := range subject.Classes {
		class.SubjectID = subject.ID
		_, err := s.ClassRepo.Create(ctx, class)
		if err != nil {
			return nil, err
		}
	}
	return subject, nil
}

func (s *SubjectRepo) Update(ctx context.Context, id int64, subject *models.Subject) (*models.Subject, error) {
	query := "UPDATE subject SET name = ?, image = ?, description = ? WHERE id = ?"
	_, err := s.DB.ExecContext(ctx, query, subject.Name, subject.Image, subject.Description, id)
	if err != nil {
		return nil, err
	}
	res, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SubjectRepo) Delete(ctx context.Context, id int64) error {
	_, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	query := "DELETE FROM subject WHERE id = ?"
	_, err = s.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
