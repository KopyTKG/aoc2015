package day2

import (
	"aoc2015/stream"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Star1() {
	fmt.Println("day 2 - star 1")

	data := stream.ReadBytes("day2/input.txt")
	lines := stream.BtoSa(data)

	sum := 0

	for _, item := range lines {
		el := strings.Split(item, "x")

		l, err := strconv.Atoi(el[0])
		if err != nil {
			log.Fatal(err)
		}

		w, err := strconv.Atoi(el[1])
		if err != nil {
			log.Fatal(err)
		}

		h, err := strconv.Atoi(el[2])
		if err != nil {
			log.Fatal(err)
		}

		a := 2 * l * w
		b := 2 * w * h
		c := 2 * h * l

		items := []int{a / 2, b / 2, c / 2}
		smallest := -1
		for _, e := range items {
			if smallest == -1 {
				smallest = e
			}

			if e < smallest {
				smallest = e
			}
		}

		sum += a + b + c + smallest
	}

	fmt.Printf("Sum is %d\n", sum)

}

func Star2() {
	fmt.Println("day 2 - star 2")

	data := stream.ReadBytes("day2/input.txt")
	lines := stream.BtoSa(data)

	sum := 0

	for _, item := range lines {
		el := strings.Split(item, "x")

		l, err := strconv.Atoi(el[0])
		if err != nil {
			log.Fatal(err)
		}

		w, err := strconv.Atoi(el[1])
		if err != nil {
			log.Fatal(err)
		}

		h, err := strconv.Atoi(el[2])
		if err != nil {
			log.Fatal(err)
		}

		sides := []int{l, w, h}
		small := 999
		smallest := 999
		for _, e := range sides {
			if e < smallest {
				small = smallest
				smallest = e
			} else if e < small {
				small = e
			}
		}

		sum += (l * w * h) + (2*small + 2*smallest)
	}

	fmt.Printf("Sum is %d\n", sum)

}
