package main

import (
	"stageSystem/config"
	"stageSystem/pkg/countrycodes"
)

var (
	smsFile,
	voiceFile,
	emailFile,
	billingFile string
)

func init () {
	countrycodes.Init()
	
	smsFile = config.GoDotEnvVariable("SMS_FILE_PATH") 
	voiceFile = config.GoDotEnvVariable("VOICE_FILE_PATH")
	emailFile = config.GoDotEnvVariable("EMAIL_FILE_PATH")
	billingFile = config.GoDotEnvVariable("BILLING_FILE_PATH")
}