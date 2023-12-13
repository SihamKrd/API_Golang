package users

import (
    "encoding/json"
    "net/http"
    "middleware/example/internal/models"
    "middleware/example/internal/services/users"
    "github.com/go-chi/chi/v5"
    "github.com/gofrs/uuid"
)

func UpdateUser(w http.ResponseWriter, r *http.Request, service *users.UserService) {
    idStr := chi.URLParam(r, "id")
    id, err := uuid.FromString(idStr)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := service.UpdateUser(id, user); err != nil {
        http.Error(w, "Failed to update user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}
