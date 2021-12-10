package main

import (
	"advent-of-code-2021/day10"
	"fmt"
	"time"
)

func main() {
	day, level := 10, 2

	fmt.Printf("Start executing Day: %d - Level: %d\n", day, level)
	startTime := time.Now()
	day10.Execute(level)
	timeNeeded := time.Since(startTime)
	fmt.Printf("Finished executing in %dμs\n", timeNeeded.Microseconds())
}
