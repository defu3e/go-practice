package mms

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"stageSystem/config"
	"stageSystem/pkg/functions"
)

type MMSData struct {  
    Country string `json:"country"` 
    Provider string `json:"provider"` 
    Bandwidth string `json:"bandwidth"` 
    ResponseTime string `json:"response_time"` 
} 

var (
    mmsApiUrl string
)

func init () {
    mmsApiUrl = config.GoDotEnvVariable("MMS_API_URL") 
}

func GetMmsStatus () []MMSData {
    resp, err := http.Get(mmsApiUrl)
    if err != nil {
        log.Fatalln("error making http request:", err)
    }
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

