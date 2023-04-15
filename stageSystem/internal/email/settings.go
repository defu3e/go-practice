package email

import (
	"stageSystem/config"
	"stageSystem/internal/functions"
)

var (
    file = config.GoDotEnvVariable("EMAIL_FILE_PATH")
    emailFile = functions.GetFullFilePath(file)
)