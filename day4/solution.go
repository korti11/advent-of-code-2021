package day4

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func Execute(level int) {
	input := utils.LoadData(4, level)

	numbers := strings.Split(input[0], ",")

	boards := make([]Board, (len(input)-1)/6) // (Input lines - the first line) / (One board lines + the empty line between the boards)
	for i := 2; i < len(input); i += 6 {
		lines := make([]string, 5)

		lines[0] = input[i]
		lines[1] = input[i+1]
		lines[2] = input[i+2]
		lines[3] = input[i+3]
		lines[4] = input[i+4]

		boards[(i-2)/6] = *toBoard(lines)
	}

	var foundBoard *Board = nil
	lastNumber := ""

	if level == 1 {
	out1:
		for i := 0; i < len(numbers); i++ {
			lastNumber = numbers[i]
			for j := 0; j < len(boards); j++ {
				board := boards[j]
				board.checkNumber(lastNumber)
				if i > 4 {
					if board.solved() {
						foundBoard = &board
						break out1
					}
				}
			}
		}
	} else {
		solvedBoards := make([]bool, len(boards))
		solved := 0
	out2:
		for i := 0; i < len(numbers); i++ {
			lastNumber = numbers[i]
			fmt.Printf("Current number: %s\n", lastNumber)
			for j := 0; j < len(boards); j++ {
				if solvedBoards[j] {
					continue
				}

				board := boards[j]
				board.checkNumber(lastNumber)

				if i > 4 && board.solved() {
					solvedBoards[j] = true
					solved++

					if solved == len(boards) {
						foundBoard = &board
						break out2
					}
				}
			}
		}
	}

	lastN, _ := strconv.ParseInt(lastNumber, 10, 64)
	output := []string{strconv.FormatInt(foundBoard.sumUnmarkedNumbers()*lastN, 10)}
	utils.WriteData(4, level, output)
}

func toBoard(input []string) *Board {
	board := make([][]string, 5)

	for i := 0; i < 5; i++ {
		split := strings.Split(input[i], " ")
		split = utils.RemoveEmptyEntries(split)
		board[i] = split
	}

	return newBoard(board)
}

type Board struct {
	board [][]string
	found [][]bool
}

func newBoard(board [][]string) *Board {
	b := new(Board)
	b.board = board

	b.found = make([][]bool, 5)
	for i := 0; i < 5; i++ {
		b.found[i] = make([]bool, 5)
	}

	return b
}

func (b *Board) checkNumber(number string) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if b.board[y][x] == number {
				b.found[y][x] = true
			}
		}
	}
}

func (b *Board) sumUnmarkedNumbers() int64 {
	sum := int64(0)
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !b.found[y][x] {
				number, _ := strconv.ParseInt(b.board[y][x], 10, 64)
				sum += number
			}
		}
	}
	return sum
}

func (b *Board) solved() bool {
	// Check rows
	for y := 0; y < 5; y++ {
		trueCount := 0
		for x := 0; x < 5; x++ {
			if b.found[y][x] {
				trueCount++
			}
		}
		if trueCount == 5 {
			return true
		}
	}

	// Check columns
	for x := 0; x < 5; x++ {
		trueCount := 0
		for y := 0; y < 5; y++ {
			if b.found[y][x] {
				trueCount++
			}
		}
		if trueCount == 5 {
			return true
		}
	}

	// Check diagonal
	trueCountR := 0
	trueCountL := 0
	for i := 0; i < 5; i++ {
		if b.found[i][i] {
			trueCountR++
		}
		if b.found[4-i][4-i] {
			trueCountL++
		}
	}
	return trueCountR == 5 || trueCountL == 5
}
