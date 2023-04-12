package email

import (
	"fmt"
	"sort"
	"strconv"
)

func formatEmailFields (fields []string) EmailData {
	DeliveryTime,_     := strconv.Atoi(fields[2])

	return EmailData{
		fields[0],
		fields[1],
		DeliveryTime,
	}
}

func ShiftAppend (em []EmailData, newEl EmailData) []EmailData {
	return []EmailData{
		em[1], em[2], newEl,
	}
}

func PrepareEmailData (em []EmailData) (m map[string][][]EmailData) {
	sort.Slice(em, func(i, j int) bool {
		return em[i].DeliveryTime > em[j].DeliveryTime
	})

	m = make (map[string][][]EmailData)
	OUTER:
	for _,v := range em {
		c := v.Country

		switch p := len(m[c]); p {
		case 0:
			add := []EmailData{v}
			m[c] = append (m[c], add)
		case 1:
			if len(m[c][0]) < 3 {
				m[c][0] = append (m[c][0], v)
			} else {
				add := []EmailData{v}
				m[c] = append (m[c], add)
				continue OUTER
			}
		case 2:
			if (len(m[c][1]) < 3) {
				m[c][1] = append (m[c][1], v)
				fmt.Println(m[c])
			} else {
				m[c][1] = ShiftAppend (m[c][1], v)
			}
		}
	}
	
	return 
}