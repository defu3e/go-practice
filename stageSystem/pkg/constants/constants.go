package constants

const (
	SMS_ITEM_LEN = 4 
	VOICE_ITEM_LEN = 8
)

var (
	PROVIDERS = map[string]map[string]struct{}{
		"SMS": {
			"Topolo": {},
			"Rond": {},
			"Kildy": {},
		},
		"MMS": {
			"Topolo": {},
			"Rond": {},
			"Kildy": {},
		},
		"VOICE": {
			"TransparentCalls": {},
			"E-Voice": {},
			"JustPhone": {},
		},
	}
)
 