package models

import (
	"github.com/gofrs/uuid"
)

type Song struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Genre  string `json:"genre"`
	}
