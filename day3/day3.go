package day3

import (
	"advent2025/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	const (
		DAY         = "3"
		INPUT_URL   = "https://adventofcode.com/2025/day/" + DAY + "/input"
		SAMPLE_FILE = "day" + DAY + "/sample.txt"
		INPUT_FILE  = "day" + DAY + "/input.txt"
	)
	fmt.Printf("================= DAY %s =================\n", DAY)
	fmt.Println("   Sample:")
	content := utils.ReadInputFromFile(SAMPLE_FILE)
	input := parse(content)
	part1(input)
	part2(input)
	fmt.Println("   Actual:")
	content = utils.ReadInputFromFile(INPUT_FILE)
	input = parse(content)
	part1(input)
	part2(input)
}

type Range struct {
	Start int
	End   int
}

func parse(input string) []string {
	lines := strings.Split(input, "\n")
	for i := range len(lines) {
		lines[i] = strings.TrimSpace(lines[i])
	}

	return lines
}

func part2(input []string) {
	sum := 0
	for _, bank := range input {
		sum += getLargestJoltage(bank, 12)
	}
	fmt.Printf("result:\n%d\n", sum)
}

func part1(input []string) {
	sum := 0
	for _, bank := range input {
		sum += getLargestJoltage(bank, 2)
	}
	fmt.Printf("result:\n%d\n", sum)
}

func getLargestJoltage(x string, maxLength int) int {
	largestAtIndex := map[string]string{}
	// keys are index underscore number of batteries
	for i := len(x) - 1; i >= 0; i-- {
		if i == len(x)-1 {
			largestAtIndex[fmt.Sprintf("%d_1", i)] = x[i:]
			continue
		}
		currMaxLength := len(x) - i
		for length := 1; length <= maxLength && length <= currMaxLength; length++ {
			curr := fmt.Sprintf("%d_%d", i, length)
			// check if there's a number of this length on the previous index
			items := []string{}
			if largestAtNextIndex, ok := largestAtIndex[fmt.Sprintf("%d_%d", i+1, length)]; ok {
				// see if i can replace the first digit or join to front of the number (ejecting the back)
				// to form a larger number
				digitStr := string(x[i])
				if length == 1 {
					if largestAtNextIndex > digitStr {
						largestAtIndex[curr] = largestAtNextIndex
					} else {
						largestAtIndex[curr] = digitStr
					}
					continue
				}

				replaceFirstDigit := digitStr + largestAtNextIndex[1:]
				ejectLastDigit := digitStr + largestAtNextIndex[:len(largestAtNextIndex)-1]
				items = append(items, replaceFirstDigit, largestAtNextIndex, ejectLastDigit)
			}

			// find the number of a -1 length and prepend digit to the number
			oneSmaller := largestAtIndex[fmt.Sprintf("%d_%d", i+1, length-1)]
			items = append(items, string(x[i])+oneSmaller)

			slices.Sort(items)
			largestAtIndex[curr] = items[len(items)-1]
		}
	}
	// find largest possible number starting at each index
	// then sort and return highest value
	largest := largestAtIndex[fmt.Sprintf("0_%d", maxLength)]
	largestInt, _ := strconv.Atoi(largest)
	// fmt.Printf("input: %s\n", x)
	// fmt.Printf("largest: %d\n", largestInt)
	return largestInt
}
