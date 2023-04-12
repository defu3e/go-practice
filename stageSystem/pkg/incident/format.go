package incident

import (
	"fmt"
	"sort"
)

func PrepareIncidentsData (inct []IncidentData) []IncidentData {
	fmt.Println(inct)
	sort.Slice(inct, func(i, j int) bool {
		return inct[i].Status == "active"
	})
	fmt.Println("\n\nsorted:", inct)

	return inct
}