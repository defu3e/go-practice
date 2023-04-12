package email

import (
	"fmt"
	"io/ioutil"
	"log"
	"stageSystem/pkg/constants"
	"stageSystem/pkg/functions"
	"strings"
)

func GetEmailData () []EmailData {
	fmt.Println("\n=== Getting email data ===")

    content, err := ioutil.ReadFile(emailFile)
    functions.CheckErr(err, constants.ERR_FATAL_MODE)

    rows := strings.Split(string(content), "\n")
    res := []EmailData{}
    
    for i, r := range rows {
        //log.Print("reading row: ", r)
        if err = validateRow (r); err != nil {
            log.Printf("error: invalid string format at %d row: <%s> | %s", i, r, err)
            continue
        } 
		
        fields := strings.Split(r, ";")
        res = append(res, formatEmailFields(fields))
    }
	
    return res
}