package main

import (
	"advent-of-code-2021/day2"
	"fmt"
)

func main() {
	day, level := 2, 2

	fmt.Printf("Start executing Day: %d - Level: %d\n", day, level)
	day2.Execute(level)
	fmt.Println("Finished executing")
}
