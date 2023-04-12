package billing

import (
	"stageSystem/config"
	"stageSystem/pkg/functions"
)

var (
	filename = config.GoDotEnvVariable("BILLING_FILE_PATH")
	billingFile = functions.GetFullFilePath(filename)
)