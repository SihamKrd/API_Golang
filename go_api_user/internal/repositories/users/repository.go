package users

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.ID, &data.Name, &data.Username, &data.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return users, err
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var user models.User
	err = row.Scan(&user.ID, &user.Name, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func CreateUser(user models.User) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    // Génération d'un nouvel UUID pour la chanson
    userID := uuid.Must(uuid.NewV4())

    _, err = db.Exec("INSERT INTO users (id, name, username, email) VALUES (?, ?, ?, ?)",
						userID.String(), user.Name, user.Username, user.Email)
	return err
}

func UpdateUser(id uuid.UUID, user models.User) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    statement, err := db.Prepare("UPDATE users SET name=?, username=?, email=? WHERE id=?")
    if err != nil {
        return err
    }
    defer statement.Close()

    _, err = statement.Exec(user.Name, user.Username, user.Email, id.String())
    return err
}

func DeleteUser(id uuid.UUID) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    statement, err := db.Prepare("DELETE FROM users WHERE id=?")
    if err != nil {
        return err
    }
    defer statement.Close()

    _, err = statement.Exec(id.String())
    return err
}