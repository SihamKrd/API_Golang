package users

import (
    "encoding/json"
    "net/http"
    "middleware/example/internal/models"
    "middleware/example/internal/services/users"
    "github.com/sirupsen/logrus"
    "fmt"
)
// CreateUser
// @Tags         users
// @Summary      Create a new user
// @Description  Create a new user with the provided information
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User Data"
// @Success      201    {object}  models.User
// @Failure      400    "Invalid request body"
// @Failure      500    "Internal Server Error"
// @Router       /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request, service *users.UserService) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        logrus.Errorf("Error decoding user data: %s", err.Error()) // Ajout d'un log ici
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := service.CreateUser(&user); err != nil {
        logrus.Errorf("error creating user: %s", err.Error()) // Log en cas d'erreur
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    body, _ := json.Marshal(user)
	_, _ = w.Write(body)
    //json.NewEncoder(w).Encode(user)
    fmt.Print("(controller)l'id user est : ",user.ID)
}