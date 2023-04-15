package result

import (
	"stageSystem/internal/billing"
	"stageSystem/internal/email"
	"stageSystem/internal/incident"
	"stageSystem/internal/mms"
	"stageSystem/internal/sms"
	"stageSystem/internal/support"
	"stageSystem/internal/voice"
)

type ResultSetT struct {
	SMS       [][]sms.SMSData					`json:"sms"`
	MMS       [][]mms.MMSData					`json:"mms"`
	VoiceCall []voice.VoiceData					`json:"voice_call"`
	Email     map[string][][]email.EmailData	`json:"email"`
	Billing   billing.BillingData				`json:"billing"`
	Support   []int								`json:"support"`
	Incidents []incident.IncidentData			`json:"incident"`
}

type ResultT struct {
	Status bool       `json:"status"` 
	Data   ResultSetT `json:"data"`   
	Error  string     `json:"error"` 
}

func GetResultData() (res ResultSetT) {
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

func PrepareResponce (data ResultSetT) ResultT {
	return ResultT{
		Status: true,
		Data: data,
		Error: "",
	}
}