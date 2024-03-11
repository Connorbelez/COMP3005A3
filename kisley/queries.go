package kisley

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	// "embed"
)

type DBConn struct {
	*sql.DB
}

// CREATE TABLE students (
//   student_id SERIAL PRIMARY KEY,
//   first_name VARCHAR(100) NOT NULL,
//   last_name VARCHAR(100) NOT NULL,
//   email VARCHAR(100) NOT NULL UNIQUE,
//   enrollment_date DATE
// );
//

type student struct {
	first_name      string
	last_name       string
	email           string
	enrollment_date string
	student_id      int
}

func ConnectDb() *sql.DB {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", // Hostname uses the Docker Compose service name
		"connorbeleznay",
		"root",
		"A3",
		// os.Getenv("POSTGRES_USER"),
		// os.Getenv("POSTGRES_PASSWORD"),
		// os.Getenv("POSTGRES_DB"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewDBConn() (*DBConn, error) {
	return &DBConn{ConnectDb()}, nil
}

type UserDB interface {
	GetAllStudents() ([]student, error)
	AddStudent(first_name string, last_name string, email string, enrollment_date string) error
	UpdateStudentEmail(student_id int, email string) error
	DeleteStudent(student_id int) error
	// GetUsers(userID string) ([]User, error)
	// PutUser(u User) (string, error)
	// PutTestUsers()
	// UpdateUser(ctx context.Context, user User) error
	// Define other operations...
}

func (conn *DBConn) GetAllStudents() ([]student, error) {
	res, err := conn.Query("SELECT * FROM students")
	if err != nil {
		return nil, err
	}

	students := make([]student, 0, 3)

	if res != nil {
		for res.Next() {
			var s student
			if err := res.Scan(&s.student_id, &s.first_name, &s.last_name, &s.email, &s.enrollment_date); err != nil {
				return students, err
			}
			students = append(students, s)
		}
	}

	return students, nil
}

func (conn *DBConn) AddStudent(first_name string, last_name string, email string, enrollment_date string) error {
	_, err := conn.Query("INSERT INTO students(first_name, last_name, email, enrollment_date )", first_name, last_name, email, enrollment_date)
	return err
}

func (conn *DBConn) UpdateStudentEmail(student_id int, email string) error {
	_, err := conn.Query("UPDATE students SET email=$2 WHERE student_id =$1", student_id, email)
	return err
}
