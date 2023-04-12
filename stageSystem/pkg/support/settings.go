package support

import "stageSystem/config"

var (
	supportApiUrl = config.GoDotEnvVariable("SUPPORT_API_URL") 
)