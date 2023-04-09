package smsgetter

import (
	"fmt"
	"io/ioutil"
	"log"
	"stageSystem/pkg/constants"
	"stageSystem/pkg/functions"
	"strings"
)

type SMSData struct { 
    Сountry string 
    Bandwidth string 
    ResponseTime string
    Provider string 
}

func GetSmsData(srcFilePath string) []SMSData {
    log.Println("Getting sms data...")
    content, err := ioutil.ReadFile(srcFilePath)
    if err != nil {
        log.Fatal(err)
    }

    rows := strings.Split(string(content), "\n")
    res := []SMSData{}
    
    for i, r := range rows {
        log.Print("reading row: ", r)
        
        if err = validateRow (r); err != nil {
            log.Printf("error: invalid string format at %d row: <%s> | %s", i, r, err)
            continue
        } 

        fields := strings.Split(r, ";")
        res = append(res, SMSData{
            fields[0],
            fields[1],
            fields[2],
            fields[3],
        })
    }

    fmt.Println(res)
    return res
}

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