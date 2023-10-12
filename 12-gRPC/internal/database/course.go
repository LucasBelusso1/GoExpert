package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryId string) (Course, error) {
	id := uuid.New().String()

	query := "INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)"
	_, err := c.db.Exec(query, id, name, description, categoryId)

	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description, CategoryID: categoryId}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		course := Course{}
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)

		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses where category_id = $1", categoryID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		course := Course{}
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)

		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}
