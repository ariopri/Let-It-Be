package models

import "context"

type Subject struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	Classes     []*Class `json:"classes"`
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
	GetByName(ctx context.Context, name string) (*Subject, error)
	GetSubjectWithClasses(ctx context.Context) ([]*Subject, error)
	Create(ctx context.Context, subject *Subject) (*Subject, error)
	Update(ctx context.Context, id int64, subject *Subject) (*Subject, error)
	Delete(ctx context.Context, id int64) error
}

type ClassRepository interface {
	GetAll(ctx context.Context) ([]*Class, error)
	GetByID(ctx context.Context, id int64) (*Class, error)
	GetByName(ctx context.Context, name string) (*Class, error)
	GetBySubjectID(ctx context.Context, subjectID int64) ([]*Class, error)
	GetBySubjectName(ctx context.Context, subjectName string) ([]*Class, error)
	Create(ctx context.Context, class *Class) (*Class, error)
	Update(ctx context.Context, id int64, class *Class) (*Class, error)
	Delete(ctx context.Context, id int64) error
}
