package constants

const (
	SMS_ITEM_LEN = 4 
	VOICE_ITEM_LEN = 8
	EMAIL_ITEM_LEN = 3
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
		"MAIL": {
			"Gmail" :{},
			"Yahoo" :{},
			"Hotmail":{},
			"MSN":{},
			"Orange":{},
			"Comcast":{},
			"AOL":{},
			"Live":{},
			"RediffMail":{},
			"GMX":{},
			"Protonmail":{},
			"Yandex":{},
			"Mail.ru":{},
		},
	}
)
 