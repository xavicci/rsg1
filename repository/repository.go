package repository

import (
	"context"

	"github.com/xavicci/rsg1/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, test *models.Test) error
	GetQuestion(ctx context.Context, id string) (*models.Question, error)
	SetQuestion(ctx context.Context, question *models.Question) error
	SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error
	GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error)
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

func SetTest(ctx context.Context, test *models.Test) error {
	return implementation.SetTest(ctx, test)
}

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return implementation.GetTest(ctx, id)
}

func SetQuestion(ctx context.Context, question *models.Question) error {
	return implementation.SetQuestion(ctx, question)
}

func GetQuestion(ctx context.Context, id string) (*models.Question, error) {
	return implementation.GetQuestion(ctx, id)
}

func SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	return implementation.SetEnrollment(ctx, enrollment)
}

func GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	return implementation.GetStudentsPerTest(ctx, testId)
}
