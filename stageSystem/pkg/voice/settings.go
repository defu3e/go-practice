package voice

import (
	"stageSystem/config"
	"stageSystem/pkg/functions"
)

var (
    srcFile = config.GoDotEnvVariable("VOICE_FILE_PATH")
    voiceFile = functions.GetFullFilePath(srcFile)
)