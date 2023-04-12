package incident

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetIncedents () []IncidentData {
	log.Println("\n=== Getting incidents data ===")
	var res []IncidentData

    resp, err := http.Get(incidentApiUrl)
    if err != nil {
        log.Println("error making http request:", err)
		return res
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