package actor

import "time"

type Actor struct {
	Id        int       `json:"id"`
	Gen       string    `json:"gender"`
	Birthdate time.Time `json:"birthdate"`
}
