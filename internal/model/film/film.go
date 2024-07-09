package model

import "time"

type Film struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"description,omitempty"`
	Date      time.Time `json:"release_date"`
	Assesment float32   `json:"assesment"`
}
