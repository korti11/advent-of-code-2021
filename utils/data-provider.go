package data

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadData(day int, level int) []string {
	filePath := fmt.Sprintf("./data/day%d-%d.txt", day, level)
	data, err := os.ReadFile(filePath)
	check(err)

	return strings.Split(string(data), "\n")
}

func WriteData(day int, level int, data []string) {
	filePath := fmt.Sprintf("./data/day%d-%d-solution.txt", day, level)
	f, err := os.Create(filePath)
	check(err)

	_, err = f.WriteString(strings.Join(data, "\n"))
	check(err)
}
