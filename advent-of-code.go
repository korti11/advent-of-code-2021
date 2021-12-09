package main

import (
	"advent-of-code-2021/day8"
	"fmt"
	"time"
)

func main() {
	day, level := 8, 1

	fmt.Printf("Start executing Day: %d - Level: %d\n", day, level)
	startTime := time.Now()
	day8.Execute(level)
	timeNeeded := time.Since(startTime)
	fmt.Printf("Finished executing in %dμs\n", timeNeeded.Microseconds())
}
