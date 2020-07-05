package mysdk

import (
	"errors"
)

func checkEmpty(field, fieldName string) error {
	if field == "" {
		return errors.New(fieldName + " is required")
	}
	return nil
}

