package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stageSystem/config"
	"stageSystem/pkg/billing"
	"stageSystem/pkg/constants"
	"stageSystem/pkg/countrycodes"
	"stageSystem/pkg/email"
	"stageSystem/pkg/functions"
	"stageSystem/pkg/incident"
	"stageSystem/pkg/mms"
	"stageSystem/pkg/sms"
	"stageSystem/pkg/support"
	"stageSystem/pkg/voice"

	"github.com/gorilla/mux"
)

type ResultSetT struct {
	SMS       [][]sms.SMSData			`json:"sms"`
	MMS       [][]mms.MMSData			`json:"mms"`
	VoiceCall []voice.VoiceData			`json:"voice_call"`
	Email     map[string][][]email.EmailData	`json:"email"`
	Billing   billing.BillingData			`json:"billing"`
	Support   []int					`json:"support"`
	Incidents []incident.IncidentData		`json:"incident"`
}

type ResultT struct {
	Status bool       `json:"status"` 
	Data   ResultSetT `json:"data"`   
	Error  string     `json:"error"` 
}

func main () {
	countrycodes.Init()
	port := config.GoDotEnvVariable("STAGE_SYSTEM_PORT") 
	listenAndServeHTTP(port)
}

func listenAndServeHTTP(port string) {
	r := mux.NewRouter()
	
	r.HandleFunc("/", handleConnection)
	r.HandleFunc("/api", handleApi).Methods("GET", "OPTIONS")
	
	log.Printf("Stage system starting on %s port...\n", port)
	
	err := http.ListenAndServe(":"+port, r)
	functions.CheckErr(err, constants.ERR_INFO_MODE)
}

func handleConnection (w http.ResponseWriter, r *http.Request)  {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "OK")
}

func handleApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	data := getResultData()
	res  := prepareResponce(data)

	json.NewEncoder(w).Encode(res)
}

func getResultData() (res ResultSetT) {
	smsData := sms.GetSmsData()
	res.SMS = sms.PrepareSmsData(smsData)

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
