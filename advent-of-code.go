package main

import (
	"advent-of-code-2021/day7"
	"fmt"
	"time"
)

func main() {
	day, level := 7, 2

	fmt.Printf("Start executing Day: %d - Level: %d\n", day, level)
	startTime := time.Now()
	day7.Execute(level)
	timeNeeded := time.Since(startTime)
	fmt.Printf("Finished executing in %dμs\n", timeNeeded.Microseconds())
}
