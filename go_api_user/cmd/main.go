package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/users"
	services "middleware/example/internal/services/users"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"
)

func main() {
	db, err := helpers.OpenDB()
    if err != nil {
        logrus.Fatalf("error while opening database: %s", err.Error())
    }
    defer helpers.CloseDB(db)

	r := chi.NewRouter()
	
	userService := services.NewUserService(db)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", users.GetUsers)
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			users.CreateUser(w, r, userService)  // Route POST pour créer un utilisateur
		})
		r.Route("/{id}", func(r chi.Router) {
			r.Use(users.Ctx)
			r.Get("/", users.GetUser)
			r.Put("/", func(w http.ResponseWriter, r *http.Request) { // Route PUT pour mettre à jour un utilisateur spécifique
				users.UpdateUser(w, r, userService)
			})
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL,
    		username VARCHAR(255) NOT NULL UNIQUE,
    		email VARCHAR(255) NOT NULL UNIQUE
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
