package day5

import (
	"aoc2015/stream"
	"fmt"
	"log"
	"strings"
)

func Star1() {
	fmt.Println("day 5 - star 1")
	data := stream.ReadBytes("day5/input.txt")
	lines := stream.BtoSa(data)

	count := 0

	for _, line := range lines {
		if isNiceStringV1(line) {
			count++
		}
	}

	fmt.Println(count)
}

func Star2() {
	fmt.Println("day 5 - star 2")
	data := stream.ReadBytes("day5/input.txt")
	lines := stream.BtoSa(data)

	count := 0

	for _, line := range lines {
		if isNiceStringV2(line) {
			count++
		}
	}

	fmt.Println(count)

}

func isNiceStringV1(s string) bool {
	// Rule 1: Contains at least three vowels
	vowelCount := 0
	for _, char := range s {
		if strings.ContainsRune("aeiou", char) {
			vowelCount++
		}
		if vowelCount >= 3 {
			break
		}
	}
	if vowelCount < 3 {
		return false
	}

	// Rule 2: Contains at least one letter that appears twice in a row
	hasDouble := false
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			hasDouble = true
			break
		}
	}
	if !hasDouble {
		return false
	}

	// Rule 3: Does not contain the strings "ab", "cd", "pq", or "xy"
	forbidden := []string{"ab", "cd", "pq", "xy"}
	for _, f := range forbidden {
		if strings.Contains(s, f) {
			return false
		}
	}

	return true
}

func isNiceStringV2(s string) bool {
	twice := false
	for i := 0; i < len(s)-1; i++ {
		sel := string(s[i]) + string(s[i+1])
		for k := i + 2; k < len(s)-1; k++ {
			next := string(s[k]) + string(s[k+1])
			if sel == next {
				twice = true
				break
			}
		}
		if twice {
			break
		}
	}
	if !twice {
		log.Printf("no twice for %s\n", s)
		return false
	}

	stepOver := false
	for i, c := range s {
		if i < len(s)-2 {
			if c == rune(s[i+2]) {
				stepOver = true
				break
			}
		}
	}
	if !stepOver {
		return false
	}
	return true
}
