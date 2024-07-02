package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestActor(t *testing.T) *Actor {

	birth, err := time.Parse("01-02-2006", "01-01-2001")
	assert.NoError(t, err)
	assert.NotNil(t, birth)

	return &Actor{
		Name:      "John Poul",
		Gen:       "man",
		Birthdate: birth,
	}
}
