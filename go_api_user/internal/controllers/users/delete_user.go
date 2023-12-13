package users

import (
    "net/http"
    "middleware/example/internal/services/users"
    "github.com/go-chi/chi/v5"
    "github.com/gofrs/uuid"
)

// DeleteUser
// @Tags         users
// @Summary      Delete a user
// @Description  Delete a user by its ID
// @Param        id    path      string  true  "user ID"
// @Success      204    "No Content - Successfully deleted"
// @Failure      400    "Invalid ID format"
// @Failure      500    "Internal Server Error - Failed to delete user"
// @Router       /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request, service *users.UserService) {
    idStr := chi.URLParam(r, "id")
    id, err := uuid.FromString(idStr)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    if err := service.DeleteUser(id); err != nil {
        http.Error(w, "Failed to delete user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
