package model

import "time"

type User struct {
	Id        int64     `json: "id" db: "id"`
	Name      string    `json: "name" db: "name"`
	Email     string    `json: "email" db: "emaill"`
	Password  string    `json: "-" db: "password"`
	CreatedAt time.Time `json: "created_at" db: "created_at"`
}

func NewUser(newUser User) *User {
	return &User{
		Id:        newUser.Id,
		Name:      newUser.Name,
		Email:     newUser.Email,
		Password:  newUser.Password,
		CreatedAt: time.Now(),
	}
}
