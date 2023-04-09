package main

import (
	"stageSystem/config"
	"stageSystem/pkg/countrycodes"
	"stageSystem/pkg/email"
	"stageSystem/pkg/mms"
	"stageSystem/pkg/smsgetter"
	"stageSystem/pkg/voice"
)

var (
	smsFile,
	voiceFile,
	emailFile string
)

func init () {
	countrycodes.Init()
	
	smsFile = config.GoDotEnvVariable("SMS_FILE_PATH") 
	voiceFile = config.GoDotEnvVariable("VOICE_FILE_PATH")
	emailFile = config.GoDotEnvVariable("EMAIL_FILE_PATH")
}

func main () {
	smsgetter.GetSmsData(smsFile)
	
	mms.GetMmsStatus()

	voice.GetVoiceData(voiceFile)

	email.GetEmailData(emailFile)
}