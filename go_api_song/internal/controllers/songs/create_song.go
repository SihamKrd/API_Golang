package songs

import (
    "encoding/json"
    "net/http"
    "middleware/example/internal/models"
    "middleware/example/internal/services/songs"
    "github.com/sirupsen/logrus"
)

func CreateSong(w http.ResponseWriter, r *http.Request, service *songs.SongService) {
    var song models.Song
    if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
        logrus.Errorf("Error decoding song data: %s", err.Error()) // Ajout d'un log ici
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := service.CreateSong(song); err != nil {
        logrus.Errorf("error creating song: %s", err.Error()) // Log en cas d'erreur
        http.Error(w, "Failed to create song", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(song)
}