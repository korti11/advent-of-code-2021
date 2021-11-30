package day

import data "advent-of-code-2021/utils"

func Execute(level int) {
	input := data.LoadData(1, level)

	output := input
	data.WriteData(1, level, output)
}
