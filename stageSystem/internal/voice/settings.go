package voice

import (
	"stageSystem/config"
	"stageSystem/internal/functions"
)

var (
    srcFile = config.GoDotEnvVariable("VOICE_FILE_PATH")
    voiceFile = functions.GetFullFilePath(srcFile)
)