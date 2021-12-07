package day7

import (
	"advent-of-code-2021/utils"
	"math"
	"strconv"
	"strings"
)

func Execute(level int) {
	input := utils.LoadData(7, level)
	intInput := utils.ConvertStringArrayToIntArray(strings.Split(input[0], ","))

	minInput := utils.FindMinInArray(intInput)
	maxInput := utils.FindMaxInArray(intInput)

	fuelUsages := make([]int64, maxInput-minInput)

	for distance := minInput; distance < maxInput; distance++ {
		fuelUsage := int64(0)
		for i := 0; i < len(intInput); i++ {
			if level == 1 {
				delta := float64(intInput[i]) - float64(distance)
				fuelUsage += int64(math.Abs(delta))
			} else {
				delta := math.Abs(float64(intInput[i] - distance))
				fuelUsage += int64(delta * (float64(1+delta) / float64(2)))
			}
		}
		fuelUsages[distance-minInput] = fuelUsage
	}

	lowestFuelUsage := utils.FindMinInArray(fuelUsages)
	output := []string{strconv.FormatInt(lowestFuelUsage, 10)}
	utils.WriteData(7, level, output)
}
