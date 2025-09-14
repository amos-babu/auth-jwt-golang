package handlers

import (
	"net/http"

	"github.com/amos-babu/auth-jwt/repository"
)

type UserHandler struct {
	Repo repository.UserRepo
}

func NewUserHandler(repo repository.UserRepo) *UserHandler {
	return &UserHandler{
		Repo: repo,
	}
}

func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

}

func (u *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {

}
