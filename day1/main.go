package day1

import (
	"aoc2015/stream"
	"fmt"
)

func Star1() {
	fmt.Println("day 1 - star 1")

	const up = 40
	const down = 41

	level := 0

	data := stream.ReadBytes("day1/input.txt")

	for i := 0; i < len(data); i++ {
		if data[i] == up {
			level++
		} else if data[i] == down {
			level--
		}
	}

	fmt.Printf("Floor %d\n", level)
}

func Star2() {
	fmt.Println("day 1 - star 2")

	const up = 40
	const down = 41

	level := 0

	data := stream.ReadBytes("day1/input.txt")

	for i := 0; i < len(data); i++ {
		if data[i] == up {
			level++
		} else if data[i] == down {
			level--
		}

		if level < 0 {
			fmt.Printf("Basement char pos %d\n", (i + 1))
			break
		}
	}

	fmt.Printf("Floor %d\n", level)
}
