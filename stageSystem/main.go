package main

import (
	"stageSystem/config"
	"stageSystem/pkg/connection"
	"stageSystem/pkg/countrycodes"
)

func main () {
	countrycodes.Init()
	port := config.GoDotEnvVariable("STAGE_SYSTEM_PORT") 
	connection.ListenAndServeHTTP(port)
}