package main

import (
	"advent-of-code-2020/utils"
	"bufio"
	"bytes"
	"errors"
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
			occupied := d.layout.seesOccupied(p)
			posType := p.posType
			switch p.posType {
			case emptyType:
				if occupied == 0 {
					posType = occupiedType
					changed = true
				}
			case occupiedType:
				if occupied >= 5 {
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
func (l layout) seesOccupied(p position) int {
	var occupied int
	for _, p := range l.sees(p) {
		if p.posType == occupiedType {
			occupied++
		}
	}
	return occupied
}
func (l layout) sees(p position) []position {
	var seats []position
	for _, dir := range utils.NumpadDirections() {
		if seat, err := l.findSeatTowards(p, dir); err == nil {
			seats = append(seats, seat)
		}

	}
	return seats
}

func (l layout) findSeatTowards(p position, dir utils.Pos) (position, error) {
	var pos = p.Add(dir.X, dir.Y)
	for l.grid.Valid(pos) {
		if l.positions[pos.Y][pos.X].posType != floorType {
			return l.positions[pos.Y][pos.X], nil
		}
		pos = pos.Add(dir.X, dir.Y)
	}
	return position{}, errors.New("not seat found")
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
