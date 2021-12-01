package day1

import (
	utils "advent-of-code-2021/utils"
	"strconv"
)

func Execute(level int) {
	input := utils.LoadData(1, level)

	intInput := utils.ConvertStringArrayToIntArray(input)
	depthMesasureCount := 0

	if level == 1 {
		for i := 1; i < len(intInput); i++ {
			prevValue := intInput[i-1]
			curValue := intInput[i]
			if prevValue < curValue {
				depthMesasureCount++
			}
		}
	} else if level == 2 {
		for i := 2; i < len(intInput)-1; i++ {
			curWindow := intInput[i-2] + intInput[i-1] + intInput[i]
			nextWindow := intInput[i-1] + intInput[i] + intInput[i+1]
			if curWindow < nextWindow {
				depthMesasureCount++
			}
		}
	}

	output := []string{strconv.FormatInt(int64(depthMesasureCount), 10)}
	utils.WriteData(1, level, output)
}
