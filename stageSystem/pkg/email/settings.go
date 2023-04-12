package email

import (
	"stageSystem/config"
	"stageSystem/pkg/functions"
)

var (
    file = config.GoDotEnvVariable("EMAIL_FILE_PATH")
    emailFile = functions.GetFullFilePath(file)
)