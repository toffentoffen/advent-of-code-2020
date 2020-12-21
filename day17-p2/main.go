package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const activeCube = '#'
const dimensions = 4

func main() {
	d, err := readInput("input.txt")
	//d, err := read(testInputReader())
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 6; i++ {
		d.cycle()
		fmt.Printf("cycle %d: %d\n", i, len(d.activeCoordinates))
	}
}

type coordinate struct {
	dimensions [dimensions]int
}

type data struct {
	activeCoordinates map[coordinate]bool
	neighbours        []coordinate
}

func (d *data) cycle() {
	var newState = make(map[coordinate]bool)
	for c, _ := range d.activeCoordinates {
		var active = d.activeNeighbours(c)
		var state = d.state(c)
		if active == 3 || (state && active == 2) {
			newState[c] = true
		}
		for _, n := range d.neighboursOf(c) {
			var active = d.activeNeighbours(n)
			var state = d.state(n)
			if active == 3 || (state && active == 2) {
				newState[n] = true
			}
		}

	}
	d.activeCoordinates = newState
}

func neighbours() []coordinate {
	var neighbours []coordinate
	for d3 := -1; d3 <= 1; d3++ {
		for d2 := -1; d2 <= 1; d2++ {
			for d1 := -1; d1 <= 1; d1++ {
				for d0 := -1; d0 <= 1; d0++ {
					if d0 != 0 || d1 != 0 || d2 != 0 || d3 != 0{
						neighbours = append(neighbours, coordinate{[dimensions]int{d0, d1, d2, d3}})
					}
				}
			}
		}
	}
	return neighbours
}

func (d data) activeNeighbours(n coordinate) int {
	var active int
	for _, n := range d.neighboursOf(n) {
		if _, ok := d.activeCoordinates[n]; ok {
			active++
		}
	}
	return active
}

func (d data) state(c coordinate) bool {
	if _, ok := d.activeCoordinates[c]; ok {
		return true
	}
	return false
}

func (d data) neighboursOf(c coordinate) []coordinate {
	var neighboursOf []coordinate
	for _, n := range d.neighbours {
		neighboursOf = append(neighboursOf, coordinate{
			[dimensions]int{
				c.dimensions[0] + n.dimensions[0],
				c.dimensions[1] + n.dimensions[1],
				c.dimensions[2] + n.dimensions[2],
				c.dimensions[3] + n.dimensions[3],
			},
		})
	}
	return neighboursOf
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
	var d = data{
		activeCoordinates: map[coordinate]bool{},
		neighbours:        neighbours(),
	}

	s := bufio.NewScanner(r)
	var d0 int
	for s.Scan() {
		line := s.Text()
		for d1, s := range line {
			if s == activeCube {
				c := coordinate{dimensions: [dimensions]int{d0, d1, 0, 0}}
				d.activeCoordinates[c] = true
			}
		}
		d0++
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`.#.
..#
###`)
}
