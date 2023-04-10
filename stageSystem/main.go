package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"stageSystem/config"
	"stageSystem/pkg/billing"
	"stageSystem/pkg/countrycodes"
	"stageSystem/pkg/email"
	"stageSystem/pkg/incident"
	"stageSystem/pkg/mms"
	"stageSystem/pkg/smsgetter"
	"stageSystem/pkg/voice"

	"github.com/gorilla/mux"
)

var (
	smsFile,
	voiceFile,
	emailFile,
	port,
	billingFile string
)

func init () {
	// remove init
	// move init files to pkg
	countrycodes.Init()
	
	smsFile = config.GoDotEnvVariable("SMS_FILE_PATH") 
	voiceFile = config.GoDotEnvVariable("VOICE_FILE_PATH")
	emailFile = config.GoDotEnvVariable("EMAIL_FILE_PATH")
	billingFile = config.GoDotEnvVariable("BILLING_FILE_PATH")
	port = config.GoDotEnvVariable("STAGE_SYSTEM_PORT") 
}

func main () {
	//data := initData()
	//fmt.Println(data)
		
	getResultData()

	listenAndServeHTTP(port)
}

type ResultSetT struct {
	SMS       [][]smsgetter.SMSData              `json:"sms"`
	MMS       [][]mms.MMSData              `json:"mms"`
	VoiceCall []voice.VoiceData          `json:"voice_call"`
	Email     [][]email.EmailData `json: email"`
	Billing   billing.BillingData              `json: billing"`
	Support   []int                    `json: support"`
	Incidents []incident.IncidentData           `json:"incident"`
}

type ResultT struct {
	Status bool       `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}
func initData () ResultT {
	var res ResultT

	url := config.GoDotEnvVariable("TEST_API_URL")
	resp, err := http.Get(url)
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

func listenAndServeHTTP(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/{arg}", handleConnection)

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
	//return "OK"
}

func getResultData() ResultSetT {
	// sms := smsgetter.GetSmsData(smsFile)

	// mms := mms.GetMmsStatus()

	// voice := voice.GetVoiceData(voiceFile)

	email := email.GetEmailData(emailFile)

	prepareEmailData(email)

	// billing := billing.GetBillingData(billingFile)


	// support := support.GetSupportData()

	// incident := incident.GetIncedents()

	//fmt.Println(sms)
	
	return ResultSetT{}
	// return ResultSetT {
	// 	prepareSmsData(sms),
	// 	prepareMmsData(mms),
	// 	voice,
	// 	prepareEmailData(email),
	// 	billing,
	// 	support,
	// 	incident,
	// };
}

func prepareSmsData (sms []smsgetter.SMSData) [][]smsgetter.SMSData {
	var res [][]smsgetter.SMSData

	for i,v := range sms {
		c,_ := countrycodes.GetByAlpha2(v.Сountry)
		sms[i].Сountry = c.Name
	}

	sort.Slice(sms, func(i, j int) bool {
		return sms[i].Provider > sms[j].Provider
	})
	res = append(res, sms)
	
	sort.Slice(sms, func(i, j int) bool {
		return sms[i].Сountry < sms[j].Сountry
	})
	res = append(res, sms)

	return res
}

func prepareMmsData (sms []mms.MMSData) [][]mms.MMSData {
	var res [][]mms.MMSData

	for i,v := range sms {
		c,_ := countrycodes.GetByAlpha2(v.Country)
		sms[i].Country = c.Name
	}

	sort.Slice(sms, func(i, j int) bool {
		return sms[i].Provider > sms[j].Provider
	})
	res = append(res, sms)
	
	sort.Slice(sms, func(i, j int) bool {
		return sms[i].Country < sms[j].Country
	})
	res = append(res, sms)

	return res
}

//[][]email.EmailData
func prepareEmailData (em []email.EmailData) /**(res map[string][]email.EmailData)**/ {
	//var res map[string][]email.EmailData 

	sort.Slice(em, func(i, j int) bool {
		return em[i].DeliveryTime > em[j].DeliveryTime
	})

	m := make (map[string][][]email.EmailData)

	maxSlice := make (map[string][]email.EmailData)
	minSlice := make (map[string][]email.EmailData)
	maxFlag := make (map[string]bool)
	minAppend := make (map[string]bool)
	for _,v := range em {
		c := v.Country

		if len (maxSlice[c]) == 0 {
			maxFlag[c] = true
		}

		if len (maxSlice[c]) < 3 && maxFlag[c] {
			maxSlice[c] = append(maxSlice[c], v)
		}

		if len (maxSlice[c]) == 3 && maxFlag[c] {
			maxFlag[c] = false
			m[v.Country] = append(m[c], maxSlice[c])
		} 


		if len (minSlice[c]) < 3 && !maxFlag[c] {
			minSlice[c] = append(minSlice[c], v)
		}

		if len (minSlice[c]) == 3 && !maxFlag[c] {

			// сдвиг и запись
			
			m[c] = append(m[c], minSlice[c])
			minAppend[c] = false;
			// fmt.Println("\n\n",len(m[c][1]))
		} 

		// m[v.Country] = [
		// 	[],
		// 	[]
		// ]
	}
	

	//return map
}