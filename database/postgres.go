package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/go/protogrpc/models"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil

}

func (this *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := this.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)
	return err
}


func (this *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	rows, err := this.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func ()  {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var students = models.Student{}
	for rows.Next() {
		err := rows.Scan(&students.Id, &students.Name, &students.Age)
		if err != nil {
			return nil, err
		}
		return &students, nil
	}

	return &students, nil
}

func (this *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := this.db.ExecContext(ctx, "INSERT INTO tests (id, name) VALUES ($1, $2)", test.Id, test.Name)
	return err
}

func (this *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	rows, err := this.db.QueryContext(ctx, "SELECT id, name FROM tests WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func ()  {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var test = models.Test{}
	for rows.Next() {
		err := rows.Scan(&test.Id, &test.Name)
		if err != nil {
			return nil, err
		}
		return &test, nil
	}

	return &test, nil
}

func (this *PostgresRepository) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := this.db.ExecContext(ctx, "INSERT INTO questions (id, answer, question, test_id) VALUES ($1, $2, $3, $4)", question.Id, question.Answer, question.Question, question.TestId)
	return err
}