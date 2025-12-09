package day1

import (
	"advent2025/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Run() {
	const (
		DAY         = "1"
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

func parse(input string) []string {
	return strings.Split(input, "\n")
}

func part2(input []string) {
	curr := 50
	count := 0

	dec := func(curr, value int) int {
		passes := value / 100
		count += passes
		mod := value % 100
		result := curr - mod
		// fmt.Printf("   passes: %d, mod: %d, result: %d\n", passes, mod, result)
		if mod != 0 && result == 0 {
			count++
		} else if result < 0 {
			result += 100
			if curr != 0 {
				count++
			}
		}
		return result
	}

	inc := func(curr, value int) int {
		passes := value / 100
		count += passes
		mod := value % 100
		result := curr + mod
		if result > 99 {
			result -= 100
			count++
		}
		return result
	}

	for _, i := range input {
		i = strings.TrimSpace(i)
		num, err := strconv.Atoi(i[1:])
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("num: %d\n", num)
		if strings.HasPrefix(i, "L") {
			curr = dec(curr, num)
		} else if strings.HasPrefix(i, "R") {
			curr = inc(curr, num)
		}
		// fmt.Printf("num: %d, curr: %d, count: %d\n", num, curr, count)
		// if curr == 0 {
		// 	count++
		// }
	}
	fmt.Printf("result:\n%d\n", count)
}

func part1(input []string) {
	dec := func(curr, value int) int {
		result := curr - value
		for result < 0 {
			result += 100
		}
		return result
	}

	inc := func(curr, value int) int {
		result := curr + value
		for result > 99 {
			result -= 100
		}
		return result
	}

	curr := 50
	count := 0
	for _, i := range input {
		i = strings.TrimSpace(i)
		num, err := strconv.Atoi(i[1:])
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("num: %d\n", num)
		if strings.HasPrefix(i, "L") {
			curr = dec(curr, num)
		} else if strings.HasPrefix(i, "R") {
			curr = inc(curr, num)
		}
		// fmt.Printf("curr: %d\n", curr)
		if curr == 0 {
			count++
		}
	}
	fmt.Printf("result:\n%d\n", count)
}

func readInputFromUrl(inputUrl string) []string {
	httpClient := http.DefaultClient

	resp, err := httpClient.Get(inputUrl)
	if err != nil {
		log.Fatalf("error fetching input from %s: %v\n", inputUrl, err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v\n", err)
	}

	return strings.Split(string(content), "\n")
}
