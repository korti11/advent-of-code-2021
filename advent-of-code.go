package main

import (
	"advent-of-code-2021/day9"
	"fmt"
	"time"
)

func main() {
	day, level := 9, 2

	fmt.Printf("Start executing Day: %d - Level: %d\n", day, level)
	startTime := time.Now()
	day9.Execute(level)
	timeNeeded := time.Since(startTime)
	fmt.Printf("Finished executing in %dμs\n", timeNeeded.Microseconds())
}
