package day6

import (
	"advent2025/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Run() {
	const (
		DAY         = "6"
		INPUT_URL   = "https://adventofcode.com/2025/day/" + DAY + "/input"
		SAMPLE_FILE = "day" + DAY + "/sample.txt"
		INPUT_FILE  = "day" + DAY + "/input.txt"
	)
	fmt.Printf("================= DAY %s =================\n", DAY)
	fmt.Println("   Sample:")
	content := utils.ReadInputFromFile(SAMPLE_FILE)
	input1 := parse1(content)
	part1(input1)
	input2 := parse2(content)
	part2(input2)
	fmt.Println("   Actual:")
	content = utils.ReadInputFromFile(INPUT_FILE)
	input1 = parse1(content)
	part1(input1)
	input2 = parse2(content)
	part2(input2)
}

func filterEmpty(input []string) []string {
	i := 0
	for i < len(input) {
		if input[i] == "" {
			input = append(input[:i], input[i+1:]...)
		} else {
			i++
		}
	}
	return input
}

func parse1(input string) [][]string {
	lines := strings.Split(input, "\n")
	results := make([][]string, len(lines))

	//// keep the whitespace to be processed by the problem
	//// find the demarcater of each problem by using the symbol
	//// as a new symbol marks the start of a new problem so then
	//// we know how many chars to take for each problem
	for i := range len(lines) {
		line := strings.TrimSpace(lines[i])
		lineSplit := strings.Split(line, " ")
		// filter out white spaces
		numbers := filterEmpty(lineSplit)
		results[i] = numbers
	}

	// validate that each array has the same number of items
	for i := range len(results) {
		if len(results[i]) != len(results[0]) {
			log.Panicf("row %d does not has the same number of items as row 0; expected %d, got %d", i, len(results[0]), len(results[i]))
		}
	}

	return results
}

func part1(input [][]string) {
	noOfProblems := len(input[0])
	noOfRows := len(input)
	opIdx := noOfRows - 1
	results := make([]int, noOfProblems)

	// loop through each vertical problem
	for i := 0; i < noOfProblems; i++ {
		// convert all numbers to ints
		numbers := make([]int, opIdx)
		for j := 0; j < opIdx; j++ {
			numbers[j], _ = strconv.Atoi(input[j][i])
		}
		op := strings.TrimSpace(input[opIdx][i])

		results[i] = numbers[0]
		for k := 1; k < opIdx; k++ {
			if op == "*" {
				results[i] *= numbers[k]
			} else if op == "+" {
				results[i] += numbers[k]
			}
		}
	}

	sum := 0
	for _, x := range results {
		sum += x
	}

	fmt.Printf("result:\n%d\n", sum)
}

func parse2(input string) [][]rune {
	lines := strings.Split(input, "\n")
	results := make([][]rune, len(lines))
	for i := 0; i < len(lines); i++ {
		results[i] = make([]rune, len(lines[i]))
	}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			results[i][j] = rune(lines[i][j])
		}
	}

	return results
}

func part2(input [][]rune) {
	// find starting index of each problem
	// starting index is demarcated by -2 from the nearest symbol on the right
	// -2 to account for whitespace between each problem
	OFFSET := 2
	noOfRows := len(input)
	opIdx := noOfRows - 1
	problemIndices := []int{len(input[opIdx]) - 1}
	// dont loop until 0 bcos start cannot be less than 0
	for i := len(input[opIdx]) - 2; i > 0; i-- {
		if input[opIdx][i] == '*' || input[opIdx][i] == '+' {
			problemIndices = append(problemIndices, i-OFFSET)
		}
	}

	noOfProblems := len(problemIndices)
	results := make([]int, noOfProblems)

	// jump through each problem indices
	for i := 0; i < noOfProblems; i++ {
		startIdx := problemIndices[i]
		// fmt.Printf("startIdx, i: %d, %d\n", startIdx, i)
		endIdx := 0
		if i != noOfProblems-1 {
			endIdx = problemIndices[i+1] + OFFSET
		}
		inputIdx := startIdx
		numbers := []int{}
		for inputIdx >= endIdx {
			numberStr := ""
			for row := 0; row < opIdx; row++ {
				numberStr += string(input[row][inputIdx])
			}
			numberStr = strings.TrimSpace(numberStr)
			number, _ := strconv.Atoi(numberStr)
			// fmt.Printf("%d ", number)
			numbers = append(numbers, number)
			inputIdx--
		}
		// fmt.Println()
		op := input[opIdx][endIdx]

		results[i] = numbers[0]
		for k := 1; k < len(numbers); k++ {
			if op == '*' {
				results[i] *= numbers[k]
			} else if op == '+' {
				results[i] += numbers[k]
			}
		}
	}

	sum := 0
	for _, x := range results {
		sum += x
	}

	fmt.Printf("result:\n%d\n", sum)
}
