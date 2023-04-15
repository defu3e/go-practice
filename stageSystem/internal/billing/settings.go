package billing

import (
	"stageSystem/config"
	"stageSystem/internal/functions"
)

var (
	filename = config.GoDotEnvVariable("BILLING_FILE_PATH")
	billingFile = functions.GetFullFilePath(filename)
)