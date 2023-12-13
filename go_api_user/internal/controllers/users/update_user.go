package users

import (
    "encoding/json"
    "net/http"
    "middleware/example/internal/models"
    "middleware/example/internal/services/users"
    "github.com/go-chi/chi/v5"
    "github.com/gofrs/uuid"
)

// UpdateUser
// @Tags         users
// @Summary      Update a user
// @Description  Update a user by its ID with new information
// @Param        id      path      string       true  "User ID"
// @Param        user    body      models.User  true  "Updated user data"
// @Success      200     {object}  models.User  "User successfully updated"
// @Failure      400     "Invalid ID format or invalid request body"
// @Failure      500     "Internal Server Error - Failed to update user"
// @Router       /users/{id} [put]
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
