package songs

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/songs"
	"net/http"
)

type SongService struct {
    db *sql.DB
}

func NewSongService(db *sql.DB) *SongService {
    return &SongService{db: db}
}

func GetAllSongs() ([]models.Song, error) {
	var err error
	// calling repository
	songs, err := repository.GetAllSongs()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return songs, nil
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	song, err := repository.GetSongById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}

func (s *SongService) CreateSong(song models.Song) error {
    err := repository.CreateSong(song)
    if err != nil {
        logrus.Errorf("error creating song: %s", err.Error())
        return &models.CustomError{
            Message: "Error creating song",
            Code:    http.StatusInternalServerError,
        }
    }
    return nil
}


func (s *SongService) UpdateSong(id uuid.UUID, song models.Song) error {
    err := repository.UpdateSong(id, song)
    if err != nil {
        logrus.Errorf("error updating song: %s", err.Error())
        return &models.CustomError{
            Message: "Error updating song",
            Code:    http.StatusInternalServerError,
        }
    }
    return nil
}

func (s *SongService) DeleteSong(id uuid.UUID) error {
    err := repository.DeleteSong(id)
    if err != nil {
        logrus.Errorf("error deleting song: %s", err.Error())
        return &models.CustomError{
            Message: "Error deleting song",
            Code:    http.StatusInternalServerError,
        }
    }
    return nil
}
