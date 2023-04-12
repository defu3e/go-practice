package voice

type VoiceData struct { 
    Сountry 			string 	`json:"country"`
	Bandwidth 			string 	`json:"bandwidth"`
    ResponseTime 		int 	`json:"response_time"`
    Provider 			string 	`json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB 				int 	`json:"ttfb"`
	VoicePurity 		int 	`json:"voice_purity"`
	MedianOfCallsTime 	int  	`json:"median_of_call_time"`
}

