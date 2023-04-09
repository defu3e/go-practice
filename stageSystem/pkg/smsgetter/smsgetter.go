package smsgetter

import (
	"fmt"
	"io/ioutil"
	"log"
	"stageSystem/pkg/countrycodes"
	"strings"
)

type SMSData struct { 
    Сountry string 
    Bandwidth string 
    ResponseTime string
    Provider string 
}

const LINE_PARTS = 4

var providers = map[string]struct{}{
    "Topolo": {},
    "Rond": {},
    "Kildy": {},
}

func init() {
    countrycodes.Init()
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
    if strings.Count(s, ";") != (LINE_PARTS - 1) {
        return fmt.Errorf("required delimiters not found") 
    }

    var (
        fields   = strings.Split(s, ";")
        a2       = fields[0]
        provider = fields[3]
    )
        
    // validate alpha field
    if _ , ok := countrycodes.GetByAlpha2(a2); !ok {
        return fmt.Errorf("string has incorrect country aplha2 code")
    }

    // validate provider
    if _, ok := providers[provider]; !ok {
        return fmt.Errorf("string has incorrect provider name")
    }

    return nil
}
