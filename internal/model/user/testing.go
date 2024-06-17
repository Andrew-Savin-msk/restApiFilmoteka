package model

import (
	"testing"
)

func TestUser(t *testing.T) *User {
	return &User{
		Email:  "yandex@mail.ru",
		Passwd: "qwerty",
	}
}
