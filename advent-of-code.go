package main

import (
	"advent-of-code-2021/day3"
	"fmt"
)

func main() {
	day, level := 3, 2

	fmt.Printf("Start executing Day: %d - Level: %d\n", day, level)
	day3.Execute(level)
	fmt.Println("Finished executing")
}
