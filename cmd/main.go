package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/amos-babu/auth-jwt/database"
	"github.com/amos-babu/auth-jwt/handlers"
	"github.com/amos-babu/auth-jwt/repository"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env files ", err)
	}

	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbName := os.Getenv("DATABASE_PASSWORD")
	dbPassword := os.Getenv("DATABASE_NAME")

	db, err := database.ConnectToDB(dbUsername, dbName, dbPassword)
	if err != nil {
		log.Fatal("Error Connecting to the database: ", err)
	}

	defer db.Close()

	userRepo := repository.NewRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	r := mux.NewRouter()
	r.HandleFunc("api/register", userHandler.Register).Methods("POST")
	r.HandleFunc("api/login", userHandler.Login).Methods("POST")

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server running at http://127.0.0.1:8080")
	log.Fatal(s.ListenAndServe())
}
