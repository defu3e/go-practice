package smsgetter

import (
	"fmt"
	"io/ioutil"
	"log"
	"stageSystem/config"
	"stageSystem/pkg/constants"
	"stageSystem/pkg/functions"
	"strings"
)
var (
    smsFile = functions.GetFullFilePath(config.GoDotEnvVariable("SMS_FILE_PATH")) 
)

func GetSmsData() []SMSData {
    fmt.Println("=== Getting sms data ===")
    
    content, err := ioutil.ReadFile(smsFile) 

    functions.CheckErr(err, constants.ERR_FATAL_MODE)

    rows := strings.Split(string(content), "\n")
    res := []SMSData{}
    
    for i, r := range rows {
        //log.Print("reading row: ", r)
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
    return res
}