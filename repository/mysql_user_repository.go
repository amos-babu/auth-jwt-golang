package repository

import (
	"database/sql"
	"fmt"

	"github.com/amos-babu/auth-jwt/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) Register(user *model.User) (*model.User, error) {
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?);`

	result, err := u.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	fmt.Println(result)
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	fetchLastUserQuery := `SELECT id, name, email, password, created_at FROM users WHERE id = ?`

	var newUser model.User
	row := u.DB.QueryRow(fetchLastUserQuery, id)
	if err = row.Scan(
		&newUser.Id,
		&newUser.Name,
		&newUser.Email,
		&newUser.Password,
		&newUser.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (u *UserRepository) Login() (*model.User, error) {
	return nil, nil
}

func (u *UserRepository) Logout() error {
	return nil
}
