package main

import (
	"stageSystem/pkg/billing"
	"stageSystem/pkg/email"
	"stageSystem/pkg/mms"
	"stageSystem/pkg/smsgetter"
	"stageSystem/pkg/voice"
)

func main () {
	smsgetter.GetSmsData(smsFile)
	
	mms.GetMmsStatus()

	voice.GetVoiceData(voiceFile)

	email.GetEmailData(emailFile)

	billing.GetBillingData(billingFile)
}