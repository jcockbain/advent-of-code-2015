package main

import (
	input "aoc2015/inpututils"
	"fmt"
)

var (
	neighbours = []loc{loc{-1, -1}, loc{-1, 0}, loc{-1, 1}, loc{0, -1}, loc{0, 1}, loc{1, -1}, loc{1, 0}, loc{1, 1}}
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt", 100))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt", 100))
}

func Part1(filename string, turns int) int {
	lines := input.ReadLines(filename)
	board := newBoard(lines)
	for i := 0; i < turns; i++ {
		board.cycle()
	}
	return board.countLive()
}

func Part2(filename string, turns int) int {
	lines := input.ReadLines(filename)
	board := newBoard(lines)
	board.turnOnCorners()
	for i := 0; i < turns; i++ {
		board.cycle2()
	}
	return board.countLive()
}

type loc struct {
	r, c int
}

type board struct {
	cells map[loc]int
	h, w  int
}

func newBoard(inp []string) board {
	h := len(inp)
	w := len(inp[0])
	cells := map[loc]int{}
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			newLoc := loc{r, c}
			cell := inp[r][c]
			if string(cell) == "#" {
				cells[newLoc] = 1
			} else if string(cell) == "." {
				cells[newLoc] = 0
			}
		}
	}
	return board{
		cells,
		h,
		w,
	}
}

func (b *board) cycle() {
	cells := b.cells
	newCells := map[loc]int{}
	for l, cell := range cells {
		liveNeighbours := 0
		for _, n := range neighbours {
			newLoc := loc{l.r + n.r, l.c + n.c}
			if (0 <= newLoc.r && newLoc.r <= b.h) && (0 <= newLoc.c && newLoc.c <= b.w) {
				liveNeighbours += cells[newLoc]
			}
		}
		if cell == 1 {
			if 2 <= liveNeighbours && liveNeighbours <= 3 {
				newCells[l] = 1
			} else {
				newCells[l] = 0
			}
		} else if cell == 0 {
			if liveNeighbours == 3 {
				newCells[l] = 1
			} else {
				newCells[l] = 0
			}
		}
	}
	b.cells = newCells
}

func (b *board) cycle2() {
	cells := b.cells
	newCells := map[loc]int{}
	for l, cell := range cells {
		if b.isCornerCell(l) {
			newCells[l] = 1
		} else {
			liveNeighbours := 0
			for _, n := range neighbours {
				newLoc := loc{l.r + n.r, l.c + n.c}
				if (0 <= newLoc.r && newLoc.r <= b.h) && (0 <= newLoc.c && newLoc.c <= b.w) {
					liveNeighbours += cells[newLoc]
				}
			}
			if cell == 1 {
				if 2 <= liveNeighbours && liveNeighbours <= 3 {
					newCells[l] = 1
				} else {
					newCells[l] = 0
				}
			} else if cell == 0 {
				if liveNeighbours == 3 {
					newCells[l] = 1
				} else {
					newCells[l] = 0
				}
			}
		}
	}
	b.cells = newCells
}

func (b *board) isCornerCell(l loc) bool {
	for _, cornerCell := range []loc{loc{0, 0}, loc{0, b.w - 1}, loc{b.h - 1, 0}, loc{b.h - 1, b.w - 1}} {
		if l == cornerCell {
			return true
		}
	}
	return false
}

func (b *board) turnOnCorners() {
	b.cells[loc{0, 0}] = 1
	b.cells[loc{0, b.w - 1}] = 1
	b.cells[loc{b.h - 1, 0}] = 1
	b.cells[loc{b.h - 1, b.w - 1}] = 1
}

func (b *board) countLive() int {
	s := 0
	for r := 0; r < b.h; r++ {
		for c := 0; c < b.w; c++ {
			s += (b.cells)[loc{r, c}]
		}
	}
	return s
}
