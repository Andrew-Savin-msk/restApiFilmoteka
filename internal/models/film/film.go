package film

import "time"

type Film struct {
	Id        int       `json:"id"`
	Name      string    `json:"Name"`
	Desc      string    `json:"description,omitempty"`
	Date      time.Time `json:"release_date"`
	Assesment int       `json:"assesment"`
}
