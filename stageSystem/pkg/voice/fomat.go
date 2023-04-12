package voice

import "strconv"

func FormatVoiceFields (fields []string) VoiceData {
	ResponseTime,_  := strconv.Atoi(fields[2])
	ConStability,_  := strconv.ParseFloat(fields[4], 32)
	TTFB,_ 		    := strconv.Atoi(fields[5])
	VoicePurity,_   := strconv.Atoi(fields[6])
	MedianOfCallsTime,_ := strconv.Atoi(fields[7])

	return VoiceData{
		Сountry: fields[0], 
		Bandwidth: fields[1],
		ResponseTime: ResponseTime,// int 
		Provider: fields[3],// string 
		ConnectionStability: float32(ConStability),// float32
		TTFB: TTFB,// int
		VoicePurity: VoicePurity,// int
		MedianOfCallsTime: MedianOfCallsTime,// int
	}
}