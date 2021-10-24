package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Cell uint64

const (
	DED Cell = iota
	Alive
)

func (c Cell) String() (s string) {
	switch c {
	case DED:
		s = " "
	case Alive:
		s = "*"
	}
	return
}

const (
	RowLength    uint64 = 50
	ColumnLength uint64 = RowLength + 10
)

type Cells []Cell

func (cs Cells) String() (str string) {
	for _, n := range cs {
		str += strconv.Itoa(int(n))
	}
	return
}

var rules = map[string]Cell{
	"111": 0,
	"110": 1,
	"101": 1,
	"100": 0,
	"011": 1,
	"010": 1,
	"001": 1,
	"000": 0,
}

func generateInitRow() (result Cells) {
	result = make(Cells, ColumnLength)
	rand.Seed(time.Now().UnixNano())
	var i uint64 = 0
	for ; i < ColumnLength; i++ {
		result[i] = Cell(rand.Intn(2))
	}
	return
}

func generateTriangleInit() (result Cells) {
	result = make(Cells, ColumnLength)
	result[len(result)-1] = Alive
	return result
}

func getNewGeneration(oldGen Cells) Cells {
	out := make(Cells, len(oldGen))

	var middIndex int
	for i := range oldGen {
		if i == len(oldGen)-2 {
			break
		}
		midd := rules[Cells{oldGen[i], oldGen[i+1], oldGen[i+2]}.String()]
		if i != middIndex {
			out[i] = oldGen[i]
		}
		out[i+1] = midd
		middIndex = i + 1
		out[i+2] = oldGen[i+2]
	}

	return out
}

func printGeneration(gen Cells) {
	for _, c := range gen {
		fmt.Printf("%v", c)
	}
	fmt.Println()
}

func main() {
	var iter uint64 = 0
	// r := generateInitRow()
	r := generateTriangleInit()
	for ; iter < RowLength; iter++ {
		r = getNewGeneration(r)
		printGeneration(r)
	}
}
