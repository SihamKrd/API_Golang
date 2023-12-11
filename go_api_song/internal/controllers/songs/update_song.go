package songs

import (
    "encoding/json"
    "net/http"
    "middleware/example/internal/models"
    "middleware/example/internal/services/songs"
    "github.com/go-chi/chi/v5"
    "github.com/gofrs/uuid"
)

func UpdateSong(w http.ResponseWriter, r *http.Request, service *songs.SongService) {
    idStr := chi.URLParam(r, "id")
    id, err := uuid.FromString(idStr)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    var song models.Song
    if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := service.UpdateSong(id, song); err != nil {
        http.Error(w, "Failed to update song", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(song)
}
