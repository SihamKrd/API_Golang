package songs

import (
    "net/http"
    "middleware/example/internal/services/songs"
    "github.com/go-chi/chi/v5"
    "github.com/gofrs/uuid"
)

func DeleteSong(w http.ResponseWriter, r *http.Request, service *songs.SongService) {
    idStr := chi.URLParam(r, "id")
    id, err := uuid.FromString(idStr)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    if err := service.DeleteSong(id); err != nil {
        http.Error(w, "Failed to delete song", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
