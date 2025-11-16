package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store interface {
	GetByID(id int) (*User, error)
}

type application struct {
	store Store
}

type User struct {
	ID   int
	Name string
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByID(id int) (*User, error) {
	row := r.db.QueryRow("select id, name from users where id = $1", id)

	var user User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepository := NewPostgresUserRepository(db)

	user, err := userRepository.GetByID(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User: %+v", user)
	}

}

type InMemoryRepository struct {
	users []User
}

func (r *InMemoryRepository) GetByID(id int) (*User, error) {
	return nil, nil
}
