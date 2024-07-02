package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Actor struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Gen       string    `json:"gender"`
	Birthdate time.Time `json:"birthdate"`
}

func (a *Actor) Validate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Name, validation.Required, validation.Length(1, 150)),
		validation.Field(&a.Gen, validation.Required),
		validation.Field(&a.Birthdate, validation.By(IsDateValid()), validation.Required),
	)
}
