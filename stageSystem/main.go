package main

import (
	"stageSystem/config"
	"stageSystem/internal/connection"
	"stageSystem/internal/countrycodes"
)

func main () {
	countrycodes.Init()
	port := config.GoDotEnvVariable("STAGE_SYSTEM_PORT") 
	connection.ListenAndServeHTTP(port)
}