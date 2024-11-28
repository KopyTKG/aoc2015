package day3

import (
	"aoc2015/stream"
	"aoc2015/vector"
	"fmt"
)

func Star1() {
	fmt.Println("day 3 - star 1")

	data := stream.ReadBytes("day3/input.txt")
	lines := stream.BtoSa(data)

	position := vector.New(0, 0)

	var houses []vector.Vector

	houses = append(houses, *position.Copy())
	sum := 1

	for _, c := range lines[0] {
		if c == '<' {
			position.X--
		}

		if c == '^' {
			position.Y--
		}

		if c == '>' {
			position.X++
		}

		if c == 'v' {
			position.Y++
		}
		has := false
		for _, house := range houses {
			if position.IsEqual(&house) {
				has = true
				break
			}
		}

		if !has {
			sum++
			houses = append(houses, *position.Copy())
		}

	}

	fmt.Printf("Total house count: %d\n", sum)

}

func Star2() {
	fmt.Println("day 3 - star 2")

	data := stream.ReadBytes("day3/input.txt")
	lines := stream.BtoSa(data)

	positionS := vector.New(0, 0)
	positionR := vector.New(0, 0)

	var houses []vector.Vector

	houses = append(houses, *positionS.Copy())
	sum := 1

	santa := true

	for _, c := range lines[0] {
		position := positionR
		if santa {
			position = positionS
		}

		if c == '<' {
			position.X--
		}

		if c == '^' {
			position.Y--
		}

		if c == '>' {
			position.X++
		}

		if c == 'v' {
			position.Y++
		}
		has := false
		for _, house := range houses {
			if position.IsEqual(&house) {
				has = true
				break
			}
		}

		if !has {
			sum++
			houses = append(houses, *position.Copy())
		}

		santa = !santa
	}

	fmt.Printf("Total house count: %d\n", sum)
}
