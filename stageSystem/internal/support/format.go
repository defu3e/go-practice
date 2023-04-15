package support

import "stageSystem/internal/constants"

func PrepareSupportData(data []SupportData) []int {
	var (
		totalActiveTickets int;
		loadLevel = constants.SUPPORT_LEVEL_LOW;
		waitingTime float32
	)
	
	for _,v := range data {
		totalActiveTickets += v.ActiveTickets
	}

	if totalActiveTickets >= constants.SUPPORT_LOW_LOAD_LIMIT {
		loadLevel = constants.SUPPORT_LEVEL_MIDDLE
		if totalActiveTickets >= constants.SUPPORT_MIDDLE_LOAD_LIMIT {
			loadLevel = constants.SUPPORT_LEVEL_HIGH
		}
	}

	waitingTime = constants.AVERAGE_TASK_PERFORM * float32(totalActiveTickets)	
	
	return []int{
		loadLevel,
		int (waitingTime),
	}
}