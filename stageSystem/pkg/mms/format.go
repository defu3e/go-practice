package mms

import (
	"sort"
	"stageSystem/pkg/countrycodes"
)

func PrepareMmsData (sms []MMSData) [][]MMSData {
	var res [][]MMSData

	for i,v := range sms {
		c,_ := countrycodes.GetByAlpha2(v.Country)
		sms[i].Country = c.Name
	}

	sort.Slice(sms, func(i, j int) bool {
		return sms[i].Provider > sms[j].Provider
	})
	res = append(res, sms)
	
	sort.Slice(sms, func(i, j int) bool {
		return sms[i].Country < sms[j].Country
	})
	res = append(res, sms)

	return res
}