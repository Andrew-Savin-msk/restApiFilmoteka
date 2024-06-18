package model

import "time"

type Actor struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Gen       string    `json:"gender"`
	Birthdate time.Time `json:"birthdate"`
}
