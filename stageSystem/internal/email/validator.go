package email

import (
	"fmt"
	"stageSystem/internal/constants"
	"stageSystem/internal/functions"
	"strings"
)

func validateRow(s string) error {
	if !functions.IsCorrectParts(s, constants.EMAIL_ITEM_LEN) {
	 return fmt.Errorf("required delimiters not found") 
	}

	var (
		fields   = strings.Split(s, ";")
		a2       = fields[0]
		provider = fields[1]
	)

	if !functions.IsValidCountryCode(a2) {
		return fmt.Errorf("string has incorrect country aplha2 code")
	}

	if !functions.IsValidProvider(provider, "EMAIL") {
		return fmt.Errorf("string has incorrect provider name")
	}

	return nil
}