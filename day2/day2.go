package day2

import (
	"advent2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	const (
		DAY         = "2"
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

func parse(input string) []Range {
	ranges := []Range{}
	rangeStrs := strings.Split(input, ",")
	for _, rangeStr := range rangeStrs {
		if rangeStr == "" {
			continue
		}
		r := strings.Split(rangeStr, "-")
		start, _ := strconv.Atoi(r[0])
		end, _ := strconv.Atoi(r[1])
		ranges = append(ranges, Range{
			Start: start,
			End:   end,
		})
	}

	return ranges
}

func part2(input []Range) {
	isRepeating := func(x string) bool {
		length := len(x)
		if length == 1 {
			return false
		}
		if length == 2 {
			return x[0] == x[1]
		}
		for i := 1; i <= length/2; i++ {
			substr := x[0:i]
			count := strings.Count(x, substr)
			if count*i == length {
				return true
			}
		}
		return false
	}

	sum := 0
	for _, r := range input {
		for i := r.Start; i <= r.End; i++ {
			if isRepeating(strconv.Itoa(i)) {
				sum += i
			}
		}
	}
	fmt.Printf("result:\n%d\n", sum)
}

func part1(input []Range) {
	isRepeating := func(x string) bool {
		length := len(x)
		if length%2 == 1 {
			return false
		}
		for i := 0; i < length/2; i++ {
			if x[i] != x[length/2+i] {
				return false
			}
		}
		return true
	}

	sum := 0
	for _, r := range input {
		for i := r.Start; i <= r.End; i++ {
			if isRepeating(strconv.Itoa(i)) {
				sum += i
			}
		}
	}
	fmt.Printf("result:\n%d\n", sum)
}
