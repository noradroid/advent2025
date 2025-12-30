package day5

import (
	"advent2025/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	const (
		DAY         = "5"
		INPUT_URL   = "https://adventofcode.com/2025/day/" + DAY + "/input"
		SAMPLE_FILE = "day" + DAY + "/sample.txt"
		INPUT_FILE  = "day" + DAY + "/input.txt"
	)
	fmt.Printf("================= DAY %s =================\n", DAY)
	fmt.Println("   Sample:")
	content := utils.ReadInputFromFile(SAMPLE_FILE)
	freshRanges, ingredients := parse(content)
	part1(freshRanges, ingredients)
	part2(freshRanges)
	fmt.Println("   Actual:")
	content = utils.ReadInputFromFile(INPUT_FILE)
	freshRanges, ingredients = parse(content)
	part1(freshRanges, ingredients)
	part2(freshRanges)
}

type Range struct {
	Start int
	End   int
}

func parse(input string) ([]Range, []int) {
	lines := strings.Split(input, "\n")
	emptyLineIdx := -1
	for i := range len(lines) {
		lines[i] = strings.TrimSpace(lines[i])
		if lines[i] == "" {
			emptyLineIdx = i
		}
	}

	freshRanges := make([]Range, len(lines[:emptyLineIdx]))
	for i, line := range lines[:emptyLineIdx] {
		startEnd := strings.Split(line, "-")
		start, _ := strconv.Atoi(startEnd[0])
		end, _ := strconv.Atoi(startEnd[1])
		// fmt.Printf("start: %d\n", start)
		// fmt.Printf("end: %d\n", end)
		freshRanges[i] = Range{
			Start: start,
			End:   end,
		}
	}

	ingredients := make([]int, len(lines[emptyLineIdx+1:]))
	for i, line := range lines[emptyLineIdx+1:] {
		ingredient, _ := strconv.Atoi(line)
		// fmt.Printf("ingredient: %d\n", ingredient)
		ingredients[i] = ingredient
	}

	return freshRanges, ingredients
}

func part2(freshRanges []Range) {
	slices.SortStableFunc(freshRanges, func(a, b Range) int {
		// if a.Start-b.Start == 0 {
		// 	return a.End - b.End
		// }
		return a.Start - b.Start
	})

	rangeIdx := 1
	fmt.Printf("length before: %d\n", len(freshRanges))
	for rangeIdx < len(freshRanges) {
		// fmt.Printf("%d\n", rangeIdx)
		currRange := freshRanges[rangeIdx]
		prevRange := freshRanges[rangeIdx-1]
		if currRange.Start <= prevRange.End && currRange.Start >= prevRange.Start {
			freshRanges[rangeIdx-1] = Range{
				Start: prevRange.Start,
				End:   max(currRange.End, prevRange.End),
			}
			fmt.Printf("prevRange: %d - %d\n", prevRange.Start, prevRange.End)
			fmt.Printf("currRange: %d - %d\n", currRange.Start, currRange.End)
			fmt.Printf("newRange: %d - %d\n", freshRanges[rangeIdx-1].Start, freshRanges[rangeIdx-1].End)

			freshRanges = append(freshRanges[:rangeIdx], freshRanges[rangeIdx+1:]...)
		} else {
			rangeIdx++
		}
	}
	fmt.Printf("length after: %d\n", len(freshRanges))

	count := 0
	for _, freshRange := range freshRanges {
		// fmt.Printf("freshRange: %d - %d\n", freshRange.Start, freshRange.End)
		count += freshRange.End - freshRange.Start + 1
	}
	fmt.Printf("result:\n%d\n", count)
}

func part1(freshRanges []Range, ingredients []int) {
	// sort freshRanges
	// iterate through freshRanges per ingredient and once the item is > ingredient
	// can stop checking for that ingredient

	slices.SortStableFunc(freshRanges, func(a, b Range) int {
		return a.Start - b.Start
	})

	slices.SortStableFunc(ingredients, func(a, b int) int {
		return a - b
	})

	count := 0
	rangeIdx := 0
	for _, ingredient := range ingredients {
		// fmt.Printf("ingredient: %d\n", ingredient)
		for rangeIdx < len(freshRanges) {
			currRange := freshRanges[rangeIdx]
			// fmt.Printf("currRange: %d - %d\n", currRange.Start, currRange.End)

			if ingredient < currRange.Start {
				break
			}

			if ingredient >= currRange.Start && ingredient <= currRange.End {
				count++
				break
			}

			rangeIdx++
		}
		if rangeIdx >= len(freshRanges) {
			break
		}
	}
	fmt.Printf("result:\n%d\n", count)
}
