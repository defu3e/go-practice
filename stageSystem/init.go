package main

import (
	"stageSystem/config"
	"stageSystem/pkg/countrycodes"
)

var (
	SmsFile,
	VoiceFile,
	EmailFile,
	BillingFile string
)

func Init () {
	countrycodes.Init()

	SmsFile = config.GoDotEnvVariable("SMS_FILE_PATH")
	VoiceFile = config.GoDotEnvVariable("VOICE_FILE_PATH")
	EmailFile = config.GoDotEnvVariable("EMAIL_FILE_PATH")
	BillingFile = config.GoDotEnvVariable("BILLING_FILE_PATH")
}