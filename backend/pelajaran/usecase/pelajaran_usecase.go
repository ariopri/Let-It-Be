package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/ariopri/Let-It-Be/tree/main/backend/models"
	"github.com/ariopri/Let-It-Be/tree/main/backend/pelajaran/repository"
)

type SubjectuseCase struct {
	SubjectRepo repository.SubjectRepo
	ClassRepo   repository.ClassRepo
	ctxTimeout  time.Duration
}

func NewSubjectUseCase(subjectRepo repository.SubjectRepo, classRepo repository.ClassRepo, ctxTimeout time.Duration) *SubjectuseCase {
	return &SubjectuseCase{
		SubjectRepo: subjectRepo,
		ClassRepo:   classRepo,
		ctxTimeout:  ctxTimeout,
	}
}

func (s *SubjectuseCase) GetAllSubjects(ctx context.Context) ([]*models.Subject, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()
	res, err := s.SubjectRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, errors.New("Timeout")
	} else if ctx.Err() == context.Canceled {
		return nil, errors.New("Canceled")
	}
	return res, nil
}

func (s *SubjectuseCase) GetSubjectByID(ctx context.Context, id int64) (*models.Subject, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()
	res, err := s.SubjectRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, errors.New("Timeout")
	} else if ctx.Err() == context.Canceled {
		return nil, errors.New("Canceled")
	}
	return res, nil
}

func (s *SubjectuseCase) GetSubjectWithClasses(ctx context.Context) ([]*models.Subject, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()
	res, err := s.SubjectRepo.GetSubjectWithClasses(ctx)
	if err != nil {
		return nil, err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, errors.New("Timeout")
	} else if ctx.Err() == context.Canceled {
		return nil, errors.New("Canceled")
	}
	return res, nil
}

func (s *SubjectuseCase) CreateSubject(ctx context.Context, subject *models.Subject) (*models.Subject, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()
	res, err := s.SubjectRepo.Create(ctx, subject)
	if err != nil {
		return nil, err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, errors.New("Timeout")
	}
	return res, nil
}

func (s *SubjectuseCase) UpdateSubject(ctx context.Context, id int64, subject *models.Subject) (*models.Subject, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()
	res, err := s.SubjectRepo.Update(ctx, id, subject)
	if err != nil {
		return nil, err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, errors.New("Timeout")
	}
	return res, nil
}

func (s *SubjectuseCase) DeleteSubject(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()
	err := s.SubjectRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return errors.New("Timeout")
	}
	return nil
}

func (s *SubjectuseCase) GetAllClasses(ctx context.Context) ([]*models.Class, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()
	res, err := s.ClassRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, errors.New("Timeout")
	}
	return res, nil
}

func (s *SubjectuseCase) GetClassByID(ctx context.Context, id int64) (*models.Class, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()
	res, err := s.ClassRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, errors.New("Timeout")
	}
	return res, nil
}

func (s *SubjectuseCase) CreateClass(ctx context.Context, class *models.Class) (*models.Class, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()
	res, err := s.ClassRepo.Create(ctx, class)
	if err != nil {
		return nil, err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, errors.New("Timeout")
	}
	return res, nil
}

func (s *SubjectuseCase) GetClassBySubjectID(ctx context.Context, id int64) ([]*models.Class, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()
	res, err := s.ClassRepo.GetBySubjectID(ctx, id)
	if err != nil {
		return nil, err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, errors.New("Timeout")
	}
	return res, nil
}

func (s *SubjectuseCase) GetClasBySubjectName(ctx context.Context, name string) ([]*models.Class, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()
	res, err := s.ClassRepo.GetBySubjectName(ctx, name)
	if err != nil {
		return nil, err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, errors.New("Timeout")
	}
	return res, nil
}

func (s *SubjectuseCase) UpdateClass(ctx context.Context, id int64, class *models.Class) (*models.Class, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()
	res, err := s.ClassRepo.Update(ctx, id, class)
	if err != nil {
		return nil, err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, errors.New("Timeout")
	}
	return res, nil
}

func (s *SubjectuseCase) DeleteClass(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()
	err := s.ClassRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return errors.New("Timeout")
	}
	return nil
}
