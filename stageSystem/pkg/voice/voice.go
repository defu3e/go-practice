package voice

import (
	"fmt"
	"io/ioutil"
	"log"
	"stageSystem/pkg/constants"
	"stageSystem/pkg/functions"
	"strconv"
	"strings"
)

type VoiceData struct { 
    Сountry string 
	Bandwidth int
    ResponseTime int 
    Provider string 
	ConnectionStability float32
	TTFB int
	VoicePurity int
	MedianOfCallsTime int
}

func GetVoiceData(srcFilePath string) []VoiceData {
    log.Println("\n=== Getting voice data ===")
    content, err := ioutil.ReadFile(srcFilePath)
    if err != nil {
        log.Fatal(err)
    }

    rows := strings.Split(string(content), "\n")
    res := []VoiceData{}
    
    for i, r := range rows {
        log.Print("reading row: ", r)
        
        if err = validateRow (r); err != nil {
            log.Printf("error: invalid string format at %d row: <%s> | %s", i, r, err)
            continue
        } 
		
        fields := strings.Split(r, ";")
        res = append(res, formatVoiceFields(fields))
    }
	
    return res
}

func formatVoiceFields (fields []string) VoiceData {
	Bandwidth,_     := strconv.Atoi(fields[1])
	ResponseTime,_  := strconv.Atoi(fields[2])
	ConStability,_  := strconv.ParseFloat(fields[3], 32)
	TTFB,_ 		    := strconv.Atoi(fields[4])
	VoicePurity,_   := strconv.Atoi(fields[5])
	MedianOfCallsTime,_ := strconv.Atoi(fields[6])

	return VoiceData{
		fields[0],
		Bandwidth,
		ResponseTime,
		fields[3],
		float32 (ConStability),
		TTFB,
		VoicePurity,
		MedianOfCallsTime,
	}
}

func validateRow(s string) error {
    // check if string has delimeter
    if strings.Count(s, ";") != (constants.VOICE_ITEM_LEN - 1) {
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
    if !functions.IsValidProvider(provider, "VOICE") {
        return fmt.Errorf("string has incorrect provider name")
    }

    return nil
}