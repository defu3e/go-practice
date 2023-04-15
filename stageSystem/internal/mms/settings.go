package mms

import "stageSystem/config"

var (
    mmsApiUrl = config.GoDotEnvVariable("MMS_API_URL") 
)
