package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/xavicci/rsg1/models"
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

func (repo *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)
	return err
}

func (repo *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO test (id, name) VALUES ($1, $2)", test.Id, test.Name)
	return err
}

func (repo *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {

	var student models.Student
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
		return &student, nil
	}
	return &student, nil
}
func (repo *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {

	var test models.Test
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name FROM test WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		err := rows.Scan(&test.Id, &test.Name)
		if err != nil {
			return nil, err
		}
		return &test, nil
	}
	return &test, nil
}

func (repo *PostgresRepository) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO questions (id, test_id, question, answer) VALUES ($1, $2, $3, $4)", question.Id, question.TestId, question.Question, question.Answer)
	return err
}

func (repo *PostgresRepository) GetQuestion(ctx context.Context, id string) (*models.Question, error) {
	var question models.Question
	rows, err := repo.db.QueryContext(ctx, "SELECT id, test_id, question, answer FROM questions WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		err := rows.Scan(&question.Id, &question.TestId, &question.Question, &question.Answer)
		if err != nil {
			return nil, err
		}
		return &question, nil
	}
	return &question, nil
}

func (repo *PostgresRepository) GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age from students WHERE id in (SELECT student_id FROM enrollments WHERE test_id = $1", testId)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var students []*models.Student
	for rows.Next() {
		var student = models.Student{}
		if err := rows.Scan(&student.Id, &student.Name, &student.Age); err == nil {
			students = append(students, &student)
		}

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return students, nil
}
