# advent-of-code-2021
This are my solutions for Advent of Code 2021 written in Go: https://adventofcode.com/2021
## Prerequisite 
- Go (1.17 at the time I'm writing this)

## Implementation notes
- Each day is implemented in a separate package. The naming schema for these packages is `day{number_of_the_day}`. Example: `day1`
- In each of these day package there is one `solution.go` file. All with the same base structure. Look at the `sample/solution.go` file for the base structure.
- The `data-provider.go` file inside the folder with the package `data` has two methods. The first one for loading the input data, the naming schema for input files is `day{number_of_the_day}-{level}.txt`. The second method is is to write out the solution, the naming schema for solution files is `day{number_of_the_day}-{level}-solution.txt`.

## Usage
To run the code use `go run .`