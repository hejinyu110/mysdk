package mysdk

import "errors"

func checkEmpty(field, fieldName string)  {
	if field == "" {
		errors.New(fieldName + " is required")
	}
}

