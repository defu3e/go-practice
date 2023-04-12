package voice

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func GetVoiceData() []VoiceData {
    fmt.Println("\n=== Getting voice data ===")
    content, err := ioutil.ReadFile(voiceFile)
    if err != nil {
        log.Fatal(err)
    }

    rows := strings.Split(string(content), "\n")
    res := []VoiceData{}
    
    for i, r := range rows {
        //log.Print("reading row: ", r)
        if err = validateRow (r); err != nil {
            log.Printf("error: invalid string format at %d row: <%s> | %s", i, r, err)
            continue
        } 
		
        fields := strings.Split(r, ";")
        res = append(res, FormatVoiceFields(fields))
    }
	
    return res
}