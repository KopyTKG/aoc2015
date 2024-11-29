package day7

import (
	"aoc2015/stream"
	"fmt"
	"strconv"
	"strings"
)

type Wire struct {
	Name     string
	Value    uint16
	Op       string
	Inputs   []string
	Computed bool
}

type Circuit map[string]*Wire

func parseInstruction(instruction string) *Wire {
	parts := strings.Split(instruction, " -> ")
	wire := &Wire{Name: parts[1]}

	inputs := strings.Split(parts[0], " ")
	if len(inputs) == 1 {
		wire.Op = "ASSIGN"
		wire.Inputs = inputs
	} else if len(inputs) == 2 {
		wire.Op = inputs[0]
		wire.Inputs = inputs[1:]
	} else {
		wire.Op = inputs[1]
		wire.Inputs = []string{inputs[0], inputs[2]}
	}

	return wire
}

func (c Circuit) evaluate(wireName string) uint16 {
	wire, exists := c[wireName]
	if !exists {
		val, err := strconv.Atoi(wireName)
		if err != nil {
			panic(err)
		}
		return uint16(val)
	}

	if wire.Computed {
		return wire.Value
	}

	wire.Computed = true
	wire.Value = 0

	var result uint16
	switch wire.Op {
	case "ASSIGN":
		if val, err := strconv.Atoi(wire.Inputs[0]); err == nil {
			result = uint16(val)
		} else {
			result = c.evaluate(wire.Inputs[0])
		}
	case "AND":
		result = c.evaluate(wire.Inputs[0]) & c.evaluate(wire.Inputs[1])
	case "OR":
		result = c.evaluate(wire.Inputs[0]) | c.evaluate(wire.Inputs[1])
	case "LSHIFT":
		shift, _ := strconv.Atoi(wire.Inputs[1])
		result = c.evaluate(wire.Inputs[0]) << uint(shift)
	case "RSHIFT":
		shift, _ := strconv.Atoi(wire.Inputs[1])
		result = c.evaluate(wire.Inputs[0]) >> uint(shift)
	case "NOT":
		result = ^c.evaluate(wire.Inputs[0])
	default:
		panic(fmt.Sprintf("Unknown operation: %s", wire.Op))
	}

	wire.Value = result
	return result
}

func Star1() {
	fmt.Println("day 7 - star 1")
	data := stream.ReadBytes("day7/input.txt")
	lines := stream.BtoSa(data)
	circuit := make(Circuit)

	for _, instruction := range lines {
		wire := parseInstruction(instruction)
		circuit[wire.Name] = wire
	}

	fmt.Println("Value of wire a:", circuit.evaluate("a"))

}

func Star2() {
	fmt.Println("day 7 - star 2")

	data := stream.ReadBytes("day7/input.txt")
	lines := stream.BtoSa(data)
	circuit := make(Circuit)

	for _, instruction := range lines {
		wire := parseInstruction(instruction)
		circuit[wire.Name] = wire
	}

	circuit["b"].Inputs[0] = "16076"
	a2 := circuit.evaluate("a")
	fmt.Println("Value of new wire a:", a2)
}
