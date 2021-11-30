package main

import (
	data "advent-of-code-2021/utils"
	"fmt"
)

func main() {
	day, level := 1, 1

	fmt.Printf("Start executing Day: %d - Level: %d\n", day, level)
	fmt.Println("Reading input file")
	input := data.LoadData(day, level)

	output := input
	fmt.Println("Writting solution file")
	data.WriteData(1, 1, output)
}
