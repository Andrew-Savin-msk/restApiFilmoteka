package model_test

import (
	"testing"
	"time"

	actor "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/actor"
	"github.com/stretchr/testify/assert"
)

func TestActorValidate(t *testing.T) {
	test_cases := []struct {
		name    string
		a       func() *actor.Actor
		isValid bool
	}{
		{
			name: "valid",
			a: func() *actor.Actor {
				a := actor.TestActor(t)
				return a
			},
			isValid: true,
		},
		{
			name: "empty name",
			a: func() *actor.Actor {
				a := actor.TestActor(t)
				a.Name = ""
				return a
			},
			isValid: false,
		},
		{
			name: "empty password",
			a: func() *actor.Actor {
				a := actor.TestActor(t)
				a.Gen = ""
				return a
			},
			isValid: false,
		},
		{
			name: "default birthdate",
			a: func() *actor.Actor {
				a := actor.TestActor(t)
				a.Birthdate = time.Time{}
				return a
			},
			isValid: false,
		},
	}

	for _, tc := range test_cases {
		if tc.isValid {
			assert.NoError(t, tc.a().Validate(), tc.name)
		} else {
			assert.Error(t, tc.a().Validate(), tc.name)
		}
	}
}
