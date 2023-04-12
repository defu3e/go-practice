package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stageSystem/config"
	"stageSystem/pkg/billing"
	"stageSystem/pkg/countrycodes"
	"stageSystem/pkg/email"
	"stageSystem/pkg/incident"
	"stageSystem/pkg/mms"
	"stageSystem/pkg/smsgetter"
	"stageSystem/pkg/support"
	"stageSystem/pkg/voice"

	"github.com/gorilla/mux"
)

var (
	port,
	billingFile string
)

func init () {
	countrycodes.Init()
	port = config.GoDotEnvVariable("STAGE_SYSTEM_PORT") 
}

func main () {
	listenAndServeHTTP(port)
}

type ResultSetT struct {
	SMS       [][]smsgetter.SMSData              `json:"sms"`
	MMS       [][]mms.MMSData              `json:"mms"`
	VoiceCall []voice.VoiceData          `json:"voice_call"`
	Email     map[string][][]email.EmailData `json:"email"`
	Billing   billing.BillingData              `json:"billing"`
	Support   []int                    `json:"support"`
	Incidents []incident.IncidentData           `json:"incident"`
}

type ResultT struct {
	Status bool       `json:"status"` 
	Data   ResultSetT `json:"data"`   
	Error  string     `json:"error"` 
}
// func initData () ResultT {
// 	var res ResultT

// 	url := config.GoDotEnvVariable("TEST_API_URL")
// 	resp, err := http.Get(url)
//     if err != nil {
//         log.Println("error making http request:", err)
//     }
//     defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
//         log.Println("non-OK response HTTP status: ", resp.StatusCode)
//     }

//     bytes, err := ioutil.ReadAll(resp.Body)
//     if err != nil {
//         log.Println(err)
//     }
   
//     if err := json.Unmarshal(bytes, &res); err != nil {  
//         log.Println("Can not unmarshal JSON", err)
//     }
// 	return res
// }

func handleTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//w.Write([]byte("{\n  \"status\": true,\n  \"data\": {\n    \"sms\": [\n      [\n        {\n          \"country\": \"Canada\",\n          \"bandwidth\": \"12\",\n          \"response_time\": \"67\",\n          \"provider\": \"Rond\"\n        },\n        {\n          \"country\": \"Great Britain\",\n          \"bandwidth\": \"98\",\n          \"response_time\": \"593\",\n          \"provider\": \"Kildy\"\n        },\n        {\n          \"country\": \"Russian Federation\",\n          \"bandwidth\": \"77\",\n          \"response_time\": \"1734\",\n          \"provider\": \"Topolo\"\n        }\n      ],\n      [\n        {\n          \"country\": \"Great Britain\",\n          \"bandwidth\": \"98\",\n          \"response_time\": \"593\",\n          \"provider\": \"Kildy\"\n        },\n        {\n          \"country\": \"Canada\",\n          \"bandwidth\": \"12\",\n          \"response_time\": \"67\",\n          \"provider\": \"Rond\"\n        },\n        {\n          \"country\": \"Russian Federation\",\n          \"bandwidth\": \"77\",\n          \"response_time\": \"1734\",\n          \"provider\": \"Topolo\"\n        }\n      ]\n    ],\n    \"mms\": [\n      [\n        {\n          \"country\": \"Great Britain\",\n          \"bandwidth\": \"98\",\n          \"response_time\": \"593\",\n          \"provider\": \"Kildy\"\n        },\n        {\n          \"country\": \"Canada\",\n          \"bandwidth\": \"12\",\n          \"response_time\": \"67\",\n          \"provider\": \"Rond\"\n        },\n        {\n          \"country\": \"Russian Federation\",\n          \"bandwidth\": \"77\",\n          \"response_time\": \"1734\",\n          \"provider\": \"Topolo\"\n        }\n      ],\n      [\n        {\n          \"country\": \"Canada\",\n          \"bandwidth\": \"12\",\n          \"response_time\": \"67\",\n          \"provider\": \"Rond\"\n        },\n        {\n          \"country\": \"Great Britain\",\n          \"bandwidth\": \"98\",\n          \"response_time\": \"593\",\n          \"provider\": \"Kildy\"\n        },\n        {\n          \"country\": \"Russian Federation\",\n          \"bandwidth\": \"77\",\n          \"response_time\": \"1734\",\n          \"provider\": \"Topolo\"\n        }\n      ]\n    ],\n    \"voice_call\": [\n      {\n        \"country\": \"US\",\n        \"bandwidth\": \"53\",\n        \"response_time\": \"321\",\n        \"provider\": \"TransparentCalls\",\n        \"connection_stability\": 0.72,\n        \"ttfb\": 442,\n        \"voice_purity\": 20,\n        \"median_of_call_time\": 5\n      },\n      {\n        \"country\": \"US\",\n        \"bandwidth\": \"53\",\n        \"response_time\": \"321\",\n        \"provider\": \"TransparentCalls\",\n        \"connection_stability\": 0.72,\n        \"ttfb\": 442,\n        \"voice_purity\": 20,\n        \"median_of_call_time\": 5\n      },\n      {\n        \"country\": \"US\",\n        \"bandwidth\": \"53\",\n        \"response_time\": \"321\",\n        \"provider\": \"E-Voice\",\n        \"connection_stability\": 0.72,\n        \"ttfb\": 442,\n        \"voice_purity\": 20,\n        \"median_of_call_time\": 5\n      },\n      {\n        \"country\": \"US\",\n        \"bandwidth\": \"53\",\n        \"response_time\": \"321\",\n        \"provider\": \"E-Voice\",\n        \"connection_stability\": 0.72,\n        \"ttfb\": 442,\n        \"voice_purity\": 20,\n        \"median_of_call_time\": 5\n      }\n    ],\n    \"email\": [\n      [\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 195\n        },\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        },\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        }\n      ],\n      [\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        },\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        },\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        }\n      ]\n    ],\n    \"billing\": {\n      \"create_customer\": true,\n      \"purchase\": true,\n      \"payout\": true,\n      \"recurring\": false,\n      \"fraud_control\": true,\n      \"checkout_page\": false\n    },\n    \"support\": [\n      3,\n      62\n    ],\n    \"incident\": [\n      {\"topic\":  \"Topic 1\", \"status\": \"active\"},\n      {\"topic\":  \"Topic 2\", \"status\": \"active\"},\n      {\"topic\":  \"Topic 3\", \"status\": \"closed\"},\n      {\"topic\":  \"Topic 4\", \"status\": \"closed\"}\n    ]\n  },\n  \"error\": \"\"\n}"))

	//w.Header().Set("Content-Type", "application/json")

	
	resData := getResultData()

	res := prepareResponce(resData)

	json.NewEncoder(w).Encode(res)
    //fmt.Println(string(b))
}

func listenAndServeHTTP(port string) {
	r := mux.NewRouter()
	//r.HandleFunc("/{arg}", handleConnection)

	r.HandleFunc("/test", handleTest).Methods("GET", "OPTIONS")

	log.Printf("Stage system starting on %s port...\n", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Println(err)
	} 
}
func handleConnection (w http.ResponseWriter, r *http.Request)  {
    //vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "OK")
	//fmt.Fprintf(w, "arg: %v\n", vars["arg"])
}


func getResultData() (res ResultSetT) {
	smsData := smsgetter.GetSmsData()
	res.SMS = smsgetter.PrepareSmsData(smsData)

	mmsData := mms.GetMmsStatus()
	res.MMS = mms.PrepareMmsData(mmsData)
	
	res.VoiceCall = voice.GetVoiceData()

	e := email.GetEmailData()
	res.Email = email.PrepareEmailData(e) 

	res.Billing = billing.GetBillingData()

	sup := support.GetSupportData()
	res.Support = support.PrepareSupportData(sup)

	incd := incident.GetIncedents()
	res.Incidents = incident.PrepareIncidentsData(incd)
	
	return
}

func prepareResponce (data ResultSetT) ResultT {
	return ResultT{
		Status: true,
		Data: data,
		Error: "",
	}
}