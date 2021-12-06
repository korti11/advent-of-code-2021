package day6

import (
	"advent-of-code-2021/utils"
	"strconv"
	"strings"
)

func Execute(level int) {
	input := utils.LoadData(6, level)

	initialStates := strings.Split(input[0], ",")
	fishPopulation := map[int64]int64{8: 0, 7: 0, 6: 0, 5: 0, 4: 0, 3: 0, 2: 0, 1: 0, 0: 0}

	var days int
	if level == 1 {
		days = 80
	} else {
		days = 256
	}

	for i := 0; i < len(initialStates); i++ {
		initialState, _ := strconv.ParseInt(initialStates[i], 10, 64)
		fishPopulation[initialState]++
	}

	for i := 0; i < days; i++ {
		newFishPopulation := map[int64]int64{8: 0, 7: 0, 6: 0, 5: 0, 4: 0, 3: 0, 2: 0, 1: 0, 0: 0}
		for time, population := range fishPopulation {
			newTime := time - 1
			if newTime < 0 {
				newFishPopulation[6] += population
				newFishPopulation[8] += population
			} else {
				newFishPopulation[newTime] += population
			}
		}
		fishPopulation = newFishPopulation
	}

	fishPopulationCount := int64(0)
	for _, population := range fishPopulation {
		fishPopulationCount += population
	}

	output := []string{strconv.FormatInt(int64(fishPopulationCount), 10)}
	utils.WriteData(6, level, output)
}
