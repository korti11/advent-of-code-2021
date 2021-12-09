package day9

import (
	"advent-of-code-2021/utils"
	"sort"
	"strconv"
	"strings"
)

func Execute(level int) {
	input := utils.LoadData(9, level)

	hightmap := make([][]int64, len(input))

	for i := 0; i < len(input); i++ {
		row := strings.Split(input[i], "")
		hightmap[i] = utils.ConvertStringArrayToIntArray(row)
	}

	if level == 1 {
		risk := int64(0)
		for y := 0; y < len(hightmap); y++ {
			for x := 0; x < len(hightmap[y]); x++ {
				if isLowpoint(hightmap, int64(y), int64(x)) {
					risk += hightmap[y][x] + 1
				}
			}
		}
		output := []string{strconv.FormatInt(risk, 10)}
		utils.WriteData(9, level, output)
	} else {
		basinSizes := make([]int, 0)
		lookedValues := map[utils.Pos]bool{}
		for y := 0; y < len(hightmap); y++ {
			for x := 0; x < len(hightmap[y]); x++ {
				if isLowpoint(hightmap, int64(y), int64(x)) {
					basinSize := int(calculateBasinSize(lookedValues, hightmap, int64(y), int64(x)))
					basinSizes = append(basinSizes, basinSize)
				}
			}
		}
		sort.Ints(basinSizes)

		basinSizeLen := len(basinSizes)
		result := basinSizes[basinSizeLen-1] * basinSizes[basinSizeLen-2] * basinSizes[basinSizeLen-3]
		output := []string{strconv.FormatInt(int64(result), 10)}
		utils.WriteData(9, level, output)
	}
}

func isLowpoint(hightMap [][]int64, y int64, x int64) bool {
	value := hightMap[y][x]

	if y > 0 && hightMap[y-1][x] <= value {
		return false
	} else if x+1 < int64(len(hightMap[y])) && hightMap[y][x+1] <= value {
		return false
	} else if y+1 < int64(len(hightMap)) && hightMap[y+1][x] <= value {
		return false
	} else if x > 0 && hightMap[y][x-1] <= value {
		return false
	}
	return true
}

func calculateBasinSize(lookedValues map[utils.Pos]bool, hightMap [][]int64, y int64, x int64) int64 {
	value := hightMap[y][x]

	if value == 9 {
		return 0
	}

	size := int64(1)
	lookedValues[utils.Pos{X: x, Y: y}] = true

	if y > 0 && hightMap[y-1][x] > value && !lookedValues[utils.Pos{X: x, Y: y - 1}] {
		size += calculateBasinSize(lookedValues, hightMap, y-1, x)
	}
	if x+1 < int64(len(hightMap[y])) && hightMap[y][x+1] > value && !lookedValues[utils.Pos{X: x + 1, Y: y}] {
		size += calculateBasinSize(lookedValues, hightMap, y, x+1)
	}
	if y+1 < int64(len(hightMap)) && hightMap[y+1][x] > value && !lookedValues[utils.Pos{X: x, Y: y + 1}] {
		size += calculateBasinSize(lookedValues, hightMap, y+1, x)
	}
	if x > 0 && hightMap[y][x-1] > value && !lookedValues[utils.Pos{X: x - 1, Y: y}] {
		size += calculateBasinSize(lookedValues, hightMap, y, x-1)
	}
	return size
}
