package users

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/users"
	"net/http"
	"fmt"
)

type UserService struct {
    db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
    return &UserService{db: db}
}

func GetAllUsers() ([]models.User, error) {
	var err error
	// calling repository
	users, err := repository.GetAllUsers()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return users, nil
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := repository.GetUserById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return user, err
}

func (s *UserService) CreateUser(user *models.User) error {
	userID := uuid.Must(uuid.NewV4())

	user.ID = &userID
	
    err := repository.CreateUser(*user)
    if err != nil {
        logrus.Errorf("error creating user: %s", err.Error())
        return &models.CustomError{
            Message: "Error creating user",
            Code:    http.StatusInternalServerError,
        }
    }
	fmt.Print("l'id user est : ",*user.ID)
    return nil
}


func (s *UserService) UpdateUser(id uuid.UUID, user *models.User) error {
    err := repository.UpdateUser(id, *user)
    if err != nil {
        logrus.Errorf("error updating user: %s", err.Error())
        return &models.CustomError{
            Message: "Error updating user",
            Code:    http.StatusInternalServerError,
        }
    }
	user.ID = &id
    return nil
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
    err := repository.DeleteUser(id)
    if err != nil {
        logrus.Errorf("error deleting user: %s", err.Error())
        return &models.CustomError{
            Message: "Error deleting user",
            Code:    http.StatusInternalServerError,
        }
    }
    return nil
}
