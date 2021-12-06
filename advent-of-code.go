package main

import (
	"advent-of-code-2021/day6"
	"fmt"
	"time"
)

func main() {
	day, level := 6, 2

	fmt.Printf("Start executing Day: %d - Level: %d\n", day, level)
	startTime := time.Now()
	day6.Execute(level)
	timeNeeded := time.Since(startTime)
	fmt.Printf("Finished executing in %dμs\n", timeNeeded.Microseconds())
}
