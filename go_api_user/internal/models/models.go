package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
    ID       *uuid.UUID `json:"id"`
    Name     string     `json:"name"`
    Username string     `json:"username"`
    Email    string     `json:"email"`
}    
