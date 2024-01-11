package songs

import (
    "encoding/json"
    "net/http"
    "middleware/example/internal/models"
    "middleware/example/internal/services/songs"
    "github.com/go-chi/chi/v5"
    "github.com/gofrs/uuid"
)

// UpdateSong
// @Tags         songs
// @Summary      Update a song
// @Description  Update a song by its ID with new information
// @Param        id      path      string       true  "Song ID"
// @Param        song    body      models.Song  true  "Updated song data"
// @Success      200     {object}  models.Song  "Song successfully updated"
// @Failure      400     "Invalid ID format or invalid request body"
// @Failure      500     "Internal Server Error - Failed to update song"
// @Router       /songs/{id} [put]
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

    if err := service.UpdateSong(id, &song); err != nil {
        http.Error(w, "Failed to update song", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(song)
}
