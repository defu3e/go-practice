package constants

const (
	SMS_ITEM_LEN = 4 
	VOICE_ITEM_LEN = 8
	EMAIL_ITEM_LEN = 3

	/** SUPPORT SYSTEM LEVELS **/
	SUPPORT_MIDDLE_LOAD_LIMIT = 16
	SUPPORT_LOW_LOAD_LIMIT = 9
	SUPPORT_LEVEL_LOW = 1
	SUPPORT_LEVEL_MIDDLE = 2
	SUPPORT_LEVEL_HIGH = 3
	AVERAGE_TASK_PERFORM = 3.33 

	ERR_INFO_MODE = 0
	ERR_FATAL_MODE = 1
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
		"EMAIL": {
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
 