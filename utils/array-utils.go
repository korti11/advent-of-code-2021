package utils

import (
	"container/list"
	"strconv"
)

func ConvertStringArrayToIntArray(input []string) []int64 {
	length := len(input)
	output := make([]int64, length)

	for i := 0; i < length; i++ {
		output[i], _ = strconv.ParseInt(input[i], 10, 64)
	}
	return output
}

func ConverStringArrayToBinIntArray(input []string) []int64 {
	length := len(input)
	output := make([]int64, length)

	for i := 0; i < length; i++ {
		output[i], _ = strconv.ParseInt(input[i], 2, 64)
	}
	return output
}

func IntArrayToList(input []int64) *list.List {
	l := list.New()

	for i := 0; i < len(input); i++ {
		l.PushBack(input[i])
	}

	return l
}
