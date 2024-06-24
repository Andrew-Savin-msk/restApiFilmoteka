package pgstore_test

import (
	"testing"

	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store/pgstore"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
)

// TODO: Find out why it always ends oh ok status
func TestValidateMapFields(t *testing.T) {
	allowedKeys := []string{"name", "gender", "birthdate"}

	// Good
	// Short name
	// No Keys
	// Extra keys
	// Only extra keys
	test_cases := []struct {
		Name    string
		Am      func() map[string]interface{}
		IsValid bool
	}{
		{
			Name: "valid map",
			Am: func() map[string]interface{} {
				return pgstore.TestActorMap(t)
			},
			IsValid: true,
		},
		{
			Name: "too short name",
			Am: func() map[string]interface{} {
				tmp := pgstore.TestActorMap(t)
				tmp["name"] = ""
				return tmp
			},
			IsValid: false,
		},
		{
			// Special error
			Name: "no keys",
			Am: func() map[string]interface{} {
				return map[string]interface{}{}
			},
			IsValid: true,
		},
		{
			Name: "extra keys",
			Am: func() map[string]interface{} {
				tmp := pgstore.TestActorMap(t)
				tmp["extra key"] = ""
				return tmp
			},
			IsValid: false,
		},
		{
			Name: "Only extra keys",
			Am: func() map[string]interface{} {
				return map[string]interface{}{
					"extra key": "",
				}
			},
			IsValid: false,
		},
	}

	for _, tc := range test_cases {
		err := validation.Validate(
			tc.Am(),
			validation.By(
				validation.RuleFunc(
					pgstore.ValidateMapFields(allowedKeys),
				),
			),
		)
		if tc.IsValid {
			assert.Nil(
				t,
				err,
				tc.Name,
			)
		} else {
			assert.NotNil(
				t,
				err,
				tc.Name,
			)
		}
	}
}
