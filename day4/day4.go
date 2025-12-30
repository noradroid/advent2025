package day4

import (
	"advent2025/utils"
	"fmt"
	"strings"
)

func Run() {
	const (
		DAY         = "4"
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

func parse(input string) [][]rune {
	lines := strings.Split(input, "\n")
	for i := range len(lines) {
		lines[i] = strings.TrimSpace(lines[i])
	}

	i := 0
	for i < len(lines) {
		if lines[i] == "" {
			lines = append(lines[:i], lines[i+1:]...)
		} else {
			i++
		}
	}

	runes := [][]rune{}
	for _, line := range lines {
		runes = append(runes, []rune(line))
	}

	return runes
}

func part2(input [][]rune) {
	total := 0
	for {
		toRemove := [][]int{}
		sum := 0
		for i := 0; i < len(input); i++ {
			for j := 0; j < len(input[i]); j++ {
				if input[i][j] == '@' {
					if isAccessible(input, i, j) {
						sum += 1
						toRemove = append(toRemove, []int{i, j})
					}
				}
			}
		}
		if sum == 0 {
			break
		}
		total += sum
		for _, coord := range toRemove {
			i := coord[0]
			j := coord[1]
			input[i][j] = '.'
		}
	}
	fmt.Printf("result:\n%d\n", total)
}

func part1(input [][]rune) {
	sum := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '@' {
				if isAccessible(input, i, j) {
					sum += 1
				}
			}
		}
	}
	fmt.Printf("result:\n%d\n", sum)
}

// isAccessible if < 4 rolls in adjacent positions
func isAccessible(grid [][]rune, i int, j int) bool {
	maxI := len(grid) - 1
	maxJ := len(grid[0]) - 1
	coords := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	count := 0
	for _, coord := range coords {
		I := i + coord[0]
		J := j + coord[1]
		if I < 0 || I > maxI || J < 0 || J > maxJ {
			continue
		}
		if grid[I][J] == '@' {
			count++
		}
	}

	return count < 4
}
