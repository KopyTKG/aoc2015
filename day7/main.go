package day7

import (
	"aoc2015/stream"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Variable struct {
	name  string
	value int
}

type Memory []Variable

func (memory *Memory) find(name string) int {
	for i, item := range *memory {
		if item.name == name {
			return i
		}
	}
	return -1
}

/*
	Je potreba dodat graf pro pomene a pak dopocitat hodnoty

	vrchol je  A a od nej udelat

	viz.: a <- lx
		   lx <- ? operace ?
*/

func Star1() {
	fmt.Println("day 7 - star 1")
	data := stream.ReadBytes("day7/test.txt")
	lines := stream.BtoSa(data)

	rAnd := regexp.MustCompile("AND")
	rOr := regexp.MustCompile("OR")
	rNot := regexp.MustCompile("NOT")
	rLeft := regexp.MustCompile(`LSHIFT`)
	rRight := regexp.MustCompile(`RSHIFT`)

	var memory Memory

	for _, l := range lines {
		items := strings.Split(l, " ")
		line := string(l)

		and := rAnd.MatchString(line)
		or := rOr.MatchString(line)
		not := rNot.MatchString(line)
		left := rLeft.MatchString(line)
		right := rRight.MatchString(line)

		if !not && !and && !or && !left && !right && len(items) == 3 {
			value, err := strconv.Atoi(items[0])
			if err != nil {
				index := memory.find(items[0])
				if index > -1 {
					value = memory[index].value
				} else {
					value = 0
				}
			}

			tmp := Variable{name: items[2], value: value}
			memory = append(memory, tmp)
		}

		if not {
			src := items[1]
			dst := items[3]

			iSRC := memory.find(src)
			iDST := memory.find(dst)
			value := 65535
			if iSRC > -1 {
				value = 65535 - memory[iSRC].value
			}

			if iDST > -1 {
				memory[iDST].value = value
			} else {

				tmp := Variable{name: dst, value: value}
				memory = append(memory, tmp)
			}

		}

		if and || or {
			iSRC1 := memory.find(items[0])
			iSRC2 := memory.find(items[2])
			dst := items[4]
			iDST := memory.find(dst)

			src1 := 0
			src2 := 0
			if iSRC1 > -1 && iSRC2 > -1 {
				src1 = memory[iSRC1].value
				src2 = memory[iSRC2].value
			}

			value := 0
			if and {
				value = src1 & src2
			} else {
				value = src1 | src2
			}

			if iDST > -1 {
				memory[iDST].value = value
			} else {

				tmp := Variable{name: dst, value: value}
				memory = append(memory, tmp)
			}

		}

		if left || right {
			iSRC := memory.find(items[0])
			size, err := strconv.Atoi(items[2])
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			dst := items[4]
			iDST := memory.find(dst)
			src := 0
			if iSRC > -1 {
				src = memory[iSRC].value
			}
			value := 0
			if left {
				value = src << size
			} else {
				value = src >> size
			}

			if iDST > -1 {
				memory[iDST].value = value
			} else {

				tmp := Variable{name: dst, value: value}
				memory = append(memory, tmp)
			}

		}

	}

	for _, i := range memory {
		fmt.Printf("%s: %d\n", i.name, i.value)
	}

}

func Star2() {
	fmt.Println("day 7 - star 2")
}
