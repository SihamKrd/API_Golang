package users

import (
    "encoding/json"
    "net/http"
    "middleware/example/internal/models"
    "middleware/example/internal/services/users"
    "github.com/sirupsen/logrus"
)

func CreateUser(w http.ResponseWriter, r *http.Request, service *users.UserService) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        logrus.Errorf("Error decoding user data: %s", err.Error()) // Ajout d'un log ici
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := service.CreateUser(user); err != nil {
        logrus.Errorf("error creating user: %s", err.Error()) // Log en cas d'erreur
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}