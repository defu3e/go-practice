package incident

import "stageSystem/config"

var (
    incidentApiUrl  = config.GoDotEnvVariable("INCIDENT_API_URL") 
)
