package film

type Film struct {
	Id        int    `json:"id"`
	Name      string `json:"Name"`
	Desc      string `json:"description,omitempty"`
	Assesment int    `json:"assesment"`
}
