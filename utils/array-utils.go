package utils

import (
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
