package user

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Passwd    string `json:"passwd,omitempty"`
	EncPasswd string `json:"-"`
}
