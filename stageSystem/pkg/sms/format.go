package sms

import (
	"sort"
	"stageSystem/pkg/countrycodes"
)

func PrepareSmsData (sms []SMSData) [][]SMSData {
	var res [][]SMSData

	for i,v := range sms {
		c,_ := countrycodes.GetByAlpha2(v.Сountry)
		sms[i].Сountry = c.Name
	}

	sort.Slice(sms, func(i, j int) bool {
		return sms[i].Provider > sms[j].Provider
	})
	res = append(res, sms)
	
	sort.Slice(sms, func(i, j int) bool {
		return sms[i].Сountry < sms[j].Сountry
	})
	res = append(res, sms)

	return res
}