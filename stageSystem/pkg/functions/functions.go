package functions

import (
	"stageSystem/pkg/constants"
	"stageSystem/pkg/countrycodes"
)

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
