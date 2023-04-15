package mms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"stageSystem/internal/constants"
	"stageSystem/internal/functions"
)

func GetMmsStatus () []MMSData {
    fmt.Println("=== Getting mms data ===")

    resp, err := http.Get(mmsApiUrl)
    functions.CheckErr(err, constants.ERR_FATAL_MODE)
    
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusOK {
        log.Fatal("non-OK response HTTP status: ", resp.StatusCode)
    }

    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    var res []MMSData
    if err := json.Unmarshal(bytes, &res); err != nil {  
        log.Fatal("Can not unmarshal JSON", err)
    }

    filterData(&res)

    return res
}

func filterData (data *[]MMSData) {
    for i,v := range *data {
        // validate alpha field
        if !functions.IsValidCountryCode(v.Country) {
            log.Println("filter data (ivalid country code):", v)
            functions.RemoveFromSlice(*data, i)
            continue
        }
        // validate provider
        if !functions.IsValidProvider(v.Provider, "MMS") {
            log.Println("filter data (invalid provider):", v)
            functions.RemoveFromSlice(*data, i)
        }
    }
}
