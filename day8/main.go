package day8

import (
	"aoc2015/stream"
	"fmt"
	"regexp"
)

func Star1() {
	fmt.Println("day 8 - star 1")
	data := stream.ReadBytes("day8/input.txt")
	lines := stream.BtoSa(data)

	rSlash := regexp.MustCompile(`\\\\`)
	rHex := regexp.MustCompile(`\\x`)
	rQuote := regexp.MustCompile(`\\"`)
	totalChar := 0
	totalString := 0
	for _, i := range lines {
		line := []byte(i)
		slash := len(rSlash.FindAll(line, -1))
		hex := len(rHex.FindAll(line, -1)) * 3
		quote := len(rQuote.FindAll(line, -1))
		charL := len(i)
		strL := charL - 2 - slash - hex - quote

		totalChar += charL
		totalString += strL

	}
	fmt.Printf("%d - %d = %d\n", totalChar, totalString, totalChar-totalString)

}

func Star2() {
	fmt.Println("day 8 - star 2")
	// data := stream.ReadBytes("day8/test.txt")
}
