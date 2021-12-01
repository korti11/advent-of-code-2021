package main

import (
	"advent-of-code-2021/day1"
	"fmt"
)

func main() {
	day, level := 1, 2

	fmt.Printf("Start executing Day: %d - Level: %d\n", day, level)
	day1.Execute(level)
	fmt.Println("Finished executing")
}
