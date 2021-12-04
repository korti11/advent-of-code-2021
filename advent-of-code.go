package main

import (
	"advent-of-code-2021/day4"
	"fmt"
)

func main() {
	day, level := 4, 2

	fmt.Printf("Start executing Day: %d - Level: %d\n", day, level)
	day4.Execute(level)
	fmt.Println("Finished executing")
}
