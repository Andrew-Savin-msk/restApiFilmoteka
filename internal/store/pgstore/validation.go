package pgstore

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

// ValidateMapFields is a custom validation rule which makes shure
// that map have only allowed keys
func ValidateMapFields(allowedKeys []string) validation.RuleFunc {
	return func(value interface{}) error {
		m, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("incalid validating type")
		}

		// Check if we have at least one allowed key
		hasKey := false
		allowedKeysMap := map[string]int{}
		for _, k := range allowedKeys {
			if _, ok := m[k]; ok {
				hasKey = true
			}
			allowedKeysMap[k]++
		}

		if !hasKey {
			return fmt.Errorf("map must contain at least one key from pool")
		}

		// Check if we have only allowed keys
		for k, _ := range m {
			if _, ok := allowedKeysMap[k]; !ok {
				return fmt.Errorf("key is not allowed for inserting to database")
			}
		}
		return nil
	}
}
