package main

import (
	"advent-of-code-2020/utils"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	d, err := readInput("input.txt")
	//d, err := read(testInputReader())
	if err != nil {
		log.Fatal(err)
	}
	var rounds int
	var changed bool
	for {
		rounds++
		d, changed = d.round()
		if !changed {
			occupied := d.occupied()
			fmt.Printf("occupied seats: %d\n", occupied)
			break
		}
	}
}

const (
	emptyType    = "L"
	floorType    = "."
	occupiedType = "#"
)

type position struct {
	utils.Pos
	posType string
}

type layout struct {
	grid      utils.Grid
	positions [][]position // rows x columns
}

type data struct {
	layout
}

func (d data) round() (*data, bool) {
	var changed bool
	var layout = layout{
		grid:      d.grid,
		positions: make([][]position, len(d.positions)),
	}
	for y, row := range d.positions {
		layout.positions[y] = make([]position, len(d.positions[y]))
		for x, p := range row {
			occupied := d.layout.adjacentOccupied(p)
			posType := p.posType
			switch p.posType {
			case emptyType:
				if occupied == 0 {
					posType = occupiedType
					changed = true
				}
			case occupiedType:
				if occupied >= 4 {
					posType = emptyType
					changed = true
				}
			}
			layout.positions[y][x] = position{
				Pos: utils.Pos{
					X: p.X,
					Y: p.Y,
				},
				posType: posType,
			}
		}
	}
	d.layout = layout
	return &d, changed
}

func (d data) print() string{
	var b bytes.Buffer
	for _, row := range d.positions {
		for _, p := range row {
			b.WriteString(p.posType)
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (d data) occupied() int {
	var occupied int
	for _, row := range d.positions {
		for _, p := range row {
			if p.posType == occupiedType {
				occupied++
			}
		}
	}
	return occupied
}
func (l layout) adjacentOccupied(p position) int {
	var occupied int
	for _, p := range l.adjacency(p) {
		if p.posType == occupiedType {
			occupied++
		}
	}
	return occupied
}
func (l layout) adjacency(p position) []position {
	var moves []position
	for _, m := range p.Numpad() {
		if l.grid.Valid(m) {
			move := l.positions[m.Y][m.X]
			if move.posType != floorType {
				moves = append(moves, move)
			}
		}
	}
	return moves
}

func readInput(path string) (*data, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()
	return read(f)
}

func read(r io.Reader) (*data, error) {
	var d data

	s := bufio.NewScanner(r)
	for s.Scan() {
		t := s.Text()
		rows := make([]position, 0, len(t))
		for i, r := range t {
			rows = append(rows, position{
				Pos: utils.Pos{
					X: i,
					Y: len(d.positions),
				},
				posType: string(r),
			})
		}
		d.positions = append(d.positions, rows)
	}
	d.grid = utils.Grid{
		X: len(d.positions[0]) - 1,
		Y: len(d.positions) - 1,
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)
}
