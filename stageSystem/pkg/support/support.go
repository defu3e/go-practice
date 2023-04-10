package support

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"stageSystem/config"
)


type SupportData struct { 
	Topic string `json:"topic"` 
	ActiveTickets int `json:"active_tickets"` 
 } 

 var (
    supportApiUrl string
)

func init () {
    supportApiUrl = config.GoDotEnvVariable("SUPPORT_API_URL") 
}

func GetSupportDatra () []SupportData {
	var res []SupportData

    resp, err := http.Get(supportApiUrl)
    if err != nil {
        log.Println("error making http request:", err)
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        log.Println("non-OK response HTTP status: ", resp.StatusCode)
    }

    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
    }

    
    if err := json.Unmarshal(bytes, &res); err != nil {  
        log.Println("Can not unmarshal JSON", err)
    }

    return res
}