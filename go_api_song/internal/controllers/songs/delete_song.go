package songs

import (
    "net/http"
    "middleware/example/internal/services/songs"
    "github.com/go-chi/chi/v5"
    "github.com/gofrs/uuid"
)

// DeleteSong
// @Tags         songs
// @Summary      Delete a song
// @Description  Delete a song by its ID
// @Param        id    path      string  true  "Song ID"
// @Success      204    "No Content - Successfully deleted"
// @Failure      400    "Invalid ID format"
// @Failure      500    "Internal Server Error - Failed to delete song"
// @Router       /songs/{id} [delete]
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
