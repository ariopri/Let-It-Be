package models

import "context"

type Subject struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type Class struct {
	ID        int64   `json:"id"`
	SubjectID int64   `json:"subject_id"`
	Teacher   string  `json:"teacher"`
	Rating    float64 `json:"rating"`
}

type SubjectRepository interface {
	GetAll(ctx context.Context) ([]*Subject, error)
	GetByID(ctx context.Context, id int64) (*Subject, error)
	Create(ctx context.Context, subject *Subject) (*Subject, error)
	Update(ctx context.Context, id int64, subject *Subject) (*Subject, error)
	Delete(ctx context.Context, id int64) error
}

type ClassRepository interface {
	GetAll(ctx context.Context) ([]*Class, error)
	GetByID(ctx context.Context, id int64) (*Class, error)
	GetBySubjectID(ctx context.Context, subjectID int64) ([]*Class, error)
	Create(ctx context.Context, class *Class) (*Class, error)
	Update(ctx context.Context, id int64, class *Class) (*Class, error)
	Delete(ctx context.Context, id int64) error
}
