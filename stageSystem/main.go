package main

import (
	"stageSystem/config"
	"stageSystem/pkg/countrycodes"
	"stageSystem/pkg/mms"
	"stageSystem/pkg/smsgetter"
	"stageSystem/pkg/voice"
)

var (
	smsFile,
	voiceFile string
)

func init () {
	countrycodes.Init()
	
	smsFile = config.GoDotEnvVariable("SMS_FILE_PATH") 
	voiceFile = config.GoDotEnvVariable("VOICE_FILE_PATH")
}

func main () {
	smsgetter.GetSmsData(smsFile)
	
	mms.GetMmsStatus()

	voice.GetVoiceData(voiceFile)
}