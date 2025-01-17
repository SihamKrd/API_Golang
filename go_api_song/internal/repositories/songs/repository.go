package songs

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.Id, &data.Title, &data.Artist, &data.Album, &data.Genre)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return songs, err
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.Id, &song.Title, &song.Artist, &song.Album, &song.Genre)
	if err != nil {
		return nil, err
	}
	return &song, err
}

func CreateSong(song models.Song) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    _, err = db.Exec("INSERT INTO songs (id, title, artist, album, genre) VALUES (?, ?, ?, ?, ?)", 
                     song.Id.String(), song.Title, song.Artist, song.Album, song.Genre)
	return err
}

func UpdateSong(id uuid.UUID, song models.Song) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    statement, err := db.Prepare("UPDATE songs SET title=?, artist=?, album=?, genre=? WHERE id=?")
    if err != nil {
        return err
    }
    defer statement.Close()

    _, err = statement.Exec(song.Title, song.Artist, song.Album, song.Genre, id.String())
    return err
}

func DeleteSong(id uuid.UUID) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    statement, err := db.Prepare("DELETE FROM songs WHERE id=?")
    if err != nil {
        return err
    }
    defer statement.Close()

    _, err = statement.Exec(id.String())
    return err
}
