package email

import (
	"fmt"
	"io/ioutil"
	"log"
	"stageSystem/pkg/constants"
	"stageSystem/pkg/functions"
	"strconv"
	"strings"
)

type EmailData struct { 
   Country string
   Provider string
   DeliveryTime int 
}

func GetEmailData (srcFilePath string) []EmailData {
	log.Println("\n=== Getting email data ===")
    content, err := ioutil.ReadFile(srcFilePath)
    if err != nil {
        log.Fatal(err)
    }

    rows := strings.Split(string(content), "\n")
    res := []EmailData{}
    
    for i, r := range rows {
        log.Print("reading row: ", r)
        if err = validateRow (r); err != nil {
            log.Printf("error: invalid string format at %d row: <%s> | %s", i, r, err)
            continue
        } 
		
        fields := strings.Split(r, ";")
        res = append(res, formatEmailFields(fields))
    }
	
    return res
}

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


func formatEmailFields (fields []string) EmailData {
	DeliveryTime,_     := strconv.Atoi(fields[2])

	return EmailData{
		fields[0],
		fields[1],
		DeliveryTime,
	}
}
