package store

import (
	"context"
	"database/sql"
)

type password struct {
	text *string
	hash []byte
}

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created_at"`
	IsActive  bool     `json:"is_active"`
	RoleID    int64    `json:"role_id"`
	Role      Role     `json:"role"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, tx *sql.Tx, user *User) error {
	query := `
		INSERT INTO users (username, password, email, role_id) VALUES 
    ($1, $2, $3, (SELECT id FROM roles WHERE name = $4))
    RETURNING id, created_at
	`

	// ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	// defer cancel()

	role := user.Role.Name
	if role == "" {
		role = "user"
	}

	err := tx.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password.hash,
		user.Email,
		role,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		switch {
		// case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
		// 	return ErrDuplicateEmail
		// case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
		// 	return ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}
