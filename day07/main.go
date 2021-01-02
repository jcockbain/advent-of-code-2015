package main

import (
	input "aoc2015/inpututils"

	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt", "a"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt", "a"))
}

var (
	notRe    = regexp.MustCompile(`NOT ([a-z]+|\d+)`)
	andRe    = regexp.MustCompile(`([a-z]+|\d+) AND ([a-z]+|\d+)`)
	orRe     = regexp.MustCompile(`([a-z]+|\d+) OR ([a-z]+|\d+)`)
	rShiftRe = regexp.MustCompile(`([a-z]+|\d+) RSHIFT (\d+)`)
	lShiftRe = regexp.MustCompile(`([a-z]+|\d+) LSHIFT (\d+)`)
)

type Wire struct {
	gate   string
	val    int
	solved bool
}

type Wires map[string]*Wire

func Part1(filename string, targetGate string) int {
	lines := input.ReadLines(filename)
	wires := Wires{}

	re := regexp.MustCompile(`(.*) -> ([a-z]+)`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		gate, wire_id := match[1], match[2]
		wire := &Wire{
			gate,
			0,
			false,
		}
		wires[wire_id] = wire
	}

	return wires.processWires(targetGate)
}

func Part2(filename string, targetGate string) int {
	lines := input.ReadLines(filename)
	wires := Wires{}

	re := regexp.MustCompile(`(.*) -> ([a-z]+)`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		gate, wire_id := match[1], match[2]
		wire := &Wire{
			gate,
			0,
			false,
		}
		wires[wire_id] = wire
	}

	part1 := Part1(filename, targetGate)
	bWire := wires["b"]
	bWire.gate = fmt.Sprintf("%v", part1)

	return wires.processWires(targetGate)
}

func (w *Wires) processWires(wire_id string) int {
	if isNumeric(wire_id) {
		return toInt(wire_id)
	}

	wire := (*w)[wire_id]
	gate := wire.gate

	if wire.solved {
		return wire.val
	}

	if isNumeric(gate) {
		return toInt(gate)
	} else if len(strings.Split(gate, " ")) == 1 {
		wire.val = w.processWires(gate)
	} else if notRe.MatchString(gate) {
		x := notRe.FindStringSubmatch(gate)[1]
		wire.val = 65535 - w.processWires(x)
	} else if andRe.MatchString(gate) {
		parts := andRe.FindStringSubmatch(gate)
		wire.val = w.processWires(parts[1]) & w.processWires(parts[2])
	} else if orRe.MatchString(gate) {
		parts := orRe.FindStringSubmatch(gate)
		wire.val = (w.processWires(parts[1])) | (w.processWires(parts[2]))
	} else if lShiftRe.MatchString(gate) {
		parts := lShiftRe.FindStringSubmatch(gate)
		wire.val = w.processWires(parts[1]) << toInt(parts[2])
	} else if rShiftRe.MatchString(gate) {
		parts := rShiftRe.FindStringSubmatch(gate)
		wire.val = w.processWires(parts[1]) >> toInt(parts[2])
	} else {
		panic(fmt.Sprintf("Invalid Gate %v for %v", gate, wire_id))
	}

	wire.solved = true
	return wire.val
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
