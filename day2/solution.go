package day2

import (
	data "advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
)

type submarine struct {
	horizontalPos int64
	depth         int64
	aim           int64
}

func Execute(level int) {
	input := data.LoadData(2, level)
	sub := submarine{0, 0, 0}

	for i := 0; i < len(input); i++ {
		instruction := strings.Split(input[i], " ")
		value, _ := strconv.ParseInt(instruction[1], 10, 32)

		if level == 1 {
			if instruction[0] == "forward" {
				sub.horizontalPos += value
			} else if instruction[0] == "down" {
				sub.depth += value
			} else if instruction[0] == "up" {
				sub.depth -= value
			} else {
				fmt.Errorf("Unknown instruction %s\n", instruction[0])
			}
		} else if level == 2 {
			if instruction[0] == "forward" {
				sub.horizontalPos += value
				sub.depth += sub.aim * value
			} else if instruction[0] == "down" {
				sub.aim += value
			} else if instruction[0] == "up" {
				sub.aim -= value
			} else {
				fmt.Errorf("Unknown instruction %s\n", instruction[0])
			}
		}
	}

	output := []string{strconv.FormatInt(sub.horizontalPos*sub.depth, 10)}
	data.WriteData(2, level, output)
}
