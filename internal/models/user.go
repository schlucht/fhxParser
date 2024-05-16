package models

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *DBModel) CreateNewUser(name, email, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	id := uuid.New()

	stmt := `INSERT INTO users (id, name, email, password) VALUES (?,?,?,?)`

	_, err := m.DB.ExecContext(ctx, stmt, id.String(), name, email, password)
	if err != nil {
		return err
	}
	return nil
}

// User mit der Email holen
func (m *DBModel) GetUserByEmail(email string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	email = strings.ToLower(email)
	var u User
	stmt := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?`
	row := m.DB.QueryRowContext(ctx, stmt, email)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return u, err
	}
	return u, nil
}
