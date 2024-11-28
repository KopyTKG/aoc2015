package day6

import (
	"aoc2015/stream"
	"aoc2015/vector"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type light struct {
	v          vector.Vector
	state      bool
	brightness int
}

func Star1() {
	fmt.Println("day 6 - star 1")

	var lights [][]light

	for i := 0; i < 1000; i++ {
		var row []light
		for k := 0; k < 1000; k++ {
			v := vector.New(i, k)
			row = append(row, light{v: *v, state: false})
		}
		lights = append(lights, row)
	}

	data := stream.ReadBytes("day6/input.txt")
	lines := stream.BtoSa(data)

	rOn := regexp.MustCompile("on")
	rOff := regexp.MustCompile("off")
	rToggle := regexp.MustCompile("toggle")
	rCords := regexp.MustCompile(`([0-9]{1,3})`)

	for _, l := range lines {
		line := string(l)
		on := rOn.MatchString(line)
		off := rOff.MatchString(line)
		toggle := rToggle.MatchString(line)
		Snumbers := rCords.FindAllString(line, -1)
		var numbers []int

		for _, item := range Snumbers {
			val, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, val)
		}

		for i := numbers[0]; i < numbers[2]+1; i++ {
			for k := numbers[1]; k < numbers[3]+1; k++ {
				if on {
					lights[i][k].state = true
				}

				if off {
					lights[i][k].state = false
				}

				if toggle {
					lights[i][k].state = !lights[i][k].state
				}

			}
		}

	}
	count := 0

	for i := 0; i < 1000; i++ {
		for k := 0; k < 1000; k++ {
			if lights[i][k].state {
				count++
			}
		}
	}

	fmt.Println(count)
}

func Star2() {
	fmt.Println("day 6 - star 2")
	var lights [][]light

	for i := 0; i < 1000; i++ {
		var row []light
		for k := 0; k < 1000; k++ {
			v := vector.New(i, k)
			row = append(row, light{v: *v, brightness: 0})
		}
		lights = append(lights, row)
	}

	data := stream.ReadBytes("day6/input.txt")
	lines := stream.BtoSa(data)

	rOn := regexp.MustCompile("on")
	rOff := regexp.MustCompile("off")
	rToggle := regexp.MustCompile("toggle")
	rCords := regexp.MustCompile(`([0-9]{1,3})`)

	for _, l := range lines {
		line := string(l)
		on := rOn.MatchString(line)
		off := rOff.MatchString(line)
		toggle := rToggle.MatchString(line)
		Snumbers := rCords.FindAllString(line, -1)
		var numbers []int

		for _, item := range Snumbers {
			val, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, val)
		}

		for i := numbers[0]; i < numbers[2]+1; i++ {
			for k := numbers[1]; k < numbers[3]+1; k++ {
				if on {
					lights[i][k].brightness += 1
				}

				if off {
					if lights[i][k].brightness > 0 {
						lights[i][k].brightness -= 1
					}
				}

				if toggle {
					lights[i][k].brightness += 2
				}

			}
		}

	}
	count := 0

	for i := 0; i < 1000; i++ {
		for k := 0; k < 1000; k++ {
			count += lights[i][k].brightness
		}
	}

	fmt.Println(count)
}
