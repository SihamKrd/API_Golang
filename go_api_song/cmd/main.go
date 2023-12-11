package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/songs"
	services "middleware/example/internal/services/songs"
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

	songService := services.NewSongService(db)

	r.Route("/songs", func(r chi.Router) {
		r.Get("/", songs.GetSongs)
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			songs.CreateSong(w, r, songService)  // Route POST pour créer une chanson
		})
		r.Route("/{id}", func(r chi.Router) {
			r.Use(songs.Ctx)
			r.Get("/", songs.GetSong)
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
		`CREATE TABLE IF NOT EXISTS songs (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			title VARCHAR(255) NOT NULL,
			artist VARCHAR(255) NOT NULL,
			album VARCHAR(255) NOT NULL,
			genre VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
