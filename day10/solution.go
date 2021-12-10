package day10

import (
	"advent-of-code-2021/utils"
	"bytes"
	"container/list"
	"io"
	"sort"
	"strconv"
)

func Execute(level int) {
	input := utils.LoadData(10, level)
	closingBrackets := map[byte](byte){40: 41, 60: 62, 91: 93, 123: 125}   // (: ), <: >, [: ], {: }
	errorValues := map[byte]int{0: 0, 41: 3, 93: 57, 125: 1197, 62: 25137} // 0: no error found, ): 3, ]: 57, }: 1197, >: 25137
	remainingValues := map[byte]int{41: 1, 93: 2, 125: 3, 62: 4}           // ): 1, ]: 2, }: 3, >: 4

	result := 0
	results := make([]int, 0)
	for i := 0; i < len(input); i++ {
		charBuffer := bytes.NewBufferString(input[i])
		char := handleChar(charBuffer, closingBrackets)
		if level == 1 && len(char) == 2 && char[0] == 126 {
			value := errorValues[char[1]]
			result += value
		} else if level == 2 && len(char) > 1 && char[0] != 126 {
			remainingBuffer := bytes.NewBuffer(char)
			result := 0
			for cur, err := remainingBuffer.ReadByte(); err != io.EOF; cur, err = remainingBuffer.ReadByte() {
				result = (result * 5) + remainingValues[cur]
			}
			results = append(results, result)
		}
	}

	if level == 2 {
		sort.Ints(results)
		result = results[(len(results)-1)/2]
	}

	output := []string{strconv.FormatInt(int64(result), 10)}
	utils.WriteData(10, level, output)
}

func handleChar(buffer *bytes.Buffer, brackets map[byte](byte)) []byte {
	bracketStack := list.New()
	for cur, err := buffer.ReadByte(); err != io.EOF; cur, err = buffer.ReadByte() {
		closing, ok := brackets[cur]

		if ok {
			bracketStack.PushBack(closing)
		} else {
			closing = bracketStack.Back().Value.(byte)

			if cur != closing {
				// fmt.Printf("Expected: %s, Found: %s, ", string(closing), string(cur))
				return []byte{126, cur}
			} else {
				bracketStack.Remove(bracketStack.Back())
			}
		}
	}

	remaining := make([]byte, 0)
	for e := bracketStack.Back(); e != nil; e = e.Prev() {
		v := e.Value.(byte)
		remaining = append(remaining, v)
	}
	return remaining
}
