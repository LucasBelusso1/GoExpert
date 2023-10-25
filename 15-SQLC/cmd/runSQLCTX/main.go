package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/LucasBelusso1/15-SQLC/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)

	if err != nil {
		errRb := tx.Rollback()
		if errRb != nil {
			return fmt.Errorf("error on roolback: %v, original error: $w", errRb, err)
		}
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})

		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			Price:       argsCourse.Price,
			CategoryID:  argsCategory.ID,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	courses, err := queries.ListCourses(ctx)

	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f\n",
			course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}
	// courseArgs := CourseParams{
	// 	ID:          uuid.NewString(),
	// 	Name:        "GO",
	// 	Description: sql.NullString{String: "Golang course"},
	// 	Price:       10.95,
	// }
	// categoryArgs := CategoryParams{
	// 	ID:          uuid.NewString(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend courses"},
	// }

	// courseDb := NewCourseDB(dbConn)
	// err = courseDb.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)

	// if err != nil {
	// 	panic(err)
	// }
}
