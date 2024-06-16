package user

type Request struct {
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
}

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Passwd    string `json:"passwd,omitempty"`
	EncPasswd string `json:"-"`
}
