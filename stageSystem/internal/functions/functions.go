package functions

import (
	"log"
	"os"
	"stageSystem/internal/constants"
	"stageSystem/internal/countrycodes"
	"strings"
)

func CheckErr (err error, mode uint8) {
    if err != nil {
        switch mode {
        case constants.ERR_INFO_MODE:
            log.Println(err)
        case constants.ERR_FATAL_MODE:
            log.Fatalln(err)
        }
    }
}

func GetFullFilePath (file string) string {
    base,_:=os.Getwd()
    return base + file
}

func RemoveFromSlice[T comparable](slice []T, s int) []T {
    return append(slice[:s], slice[s+1:]...)
}

func IsValidCountryCode (code string) bool {
    _, ok := countrycodes.GetByAlpha2(code)
    return ok
}

func IsValidProvider (provider string, tp string) bool {
    _, ok := constants.PROVIDERS[tp][provider]
    return ok
}

func IsCorrectParts (s string, n int) bool {
    return strings.Count(s, ";") == (n - 1)
}
