package sms

import (
	"fmt"
	"stageSystem/internal/constants"
	"stageSystem/internal/functions"
	"strings"
)

func validateRow(s string) error {
    // check if string has delimeter
    if strings.Count(s, ";") != (constants.SMS_ITEM_LEN - 1) {
        return fmt.Errorf("required delimiters not found") 
    }

    var (
        fields   = strings.Split(s, ";")
        a2       = fields[0]
        provider = fields[3]
    )

    // validate alpha field
    if !functions.IsValidCountryCode(a2) {
        return fmt.Errorf("string has incorrect country aplha2 code")
    }

    // validate provider
    if !functions.IsValidProvider(provider, "SMS") {
        return fmt.Errorf("string has incorrect provider name")
    }

    return nil
}