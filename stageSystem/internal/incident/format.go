package incident

import (
	"sort"
)

func PrepareIncidentsData (inct []IncidentData) []IncidentData {
	sort.Slice(inct, func(i, j int) bool {
		return inct[i].Status == "active"
	})

	return inct
}