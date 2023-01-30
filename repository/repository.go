package repository

import (
	"context"

	"github.com/go/protogrpc/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, text *models.Test) error
	SetQuestion(ctx context.Context, question *models.Question) error
}

var implementations Repository

func SetRepository(repository Repository) {
	implementations = repository
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementations.SetStudent(ctx, student)
}

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return implementations.GetTest(ctx, id)
}

func SetTest(ctx context.Context, test *models.Test) error {
	return implementations.SetTest(ctx, test)
}

func SetQuestion(ctx context.Context, question *models.Question) error {
	return implementations.SetQuestion(ctx, question)
}