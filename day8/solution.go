package day8

import (
	"advent-of-code-2021/utils"
	"strconv"
	"strings"
)

func Execute(level int) {
	input := utils.LoadData(8, level)

	signals := make([]*Signal, len(input))
	digitOutputs := make([]*DigitOutput, len(input))

	for i := 0; i < len(input); i++ {
		sig, out := parseInput(input[i])
		signals[i] = sig
		digitOutputs[i] = out
	}

	uniqueDigitCount := 0
	for i := 0; i < len(digitOutputs); i++ {
		uniqueDigitCount += digitOutputs[i].countUniqueDigits()
	}

	output := []string{strconv.FormatInt(int64(uniqueDigitCount), 10)}
	utils.WriteData(8, level, output)
}

type Signal struct {
	patterns []string
	digits   map[int](int)
}

type DigitOutput struct {
	digits []string
}

func parseInput(input string) (*Signal, *DigitOutput) {
	split := strings.Split(input, " | ")
	sig := new(Signal)
	out := new(DigitOutput)

	sig.patterns = strings.Split(split[0], " ")
	out.digits = strings.Split(split[1], " ")

	return sig, out
}

func (out *DigitOutput) countUniqueDigits() int {
	count := 0
	for i := 0; i < len(out.digits); i++ {
		digit := out.digits[i]
		segments := len(digit)
		if segments == 2 || segments == 4 || segments == 3 || segments == 7 {
			count++
		}
	}
	return count
}
