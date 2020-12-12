package main

import (
	"advent-of-code-2020/utils"
	"bufio"
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
	var s = ship{
		facing:    E,
		distances: make(map[direction]int),
	}
	for _, a := range d.actions {
		s.apply(a)
	}
	fmt.Printf("distance: %d\n", s.distance())
}

type action struct {
	code string
	val  int
}
type data struct {
	actions []action
}
type direction string

const N direction = "N"
const S direction = "S"
const E direction = "E"
const W direction = "W"

type ship struct {
	facing    direction
	distances map[direction]int
}

func (s *ship) apply(a action) {
	switch a.code {
	case "N", "S", "E", "W":
		s.distances[direction(a.code)] += a.val
	case "F":
		s.distances[s.facing] += a.val
	case "R":
		s.turnRight(a.val)
	case "L":
		s.turnLeft(a.val)
	}
}

func (s *ship) distance() int {
	return utils.Abs(s.distances[N]-s.distances[S]) + utils.Abs(s.distances[E]-s.distances[W])
}

func (s *ship) String() string {
	return fmt.Sprintf("facing %s distances N=%d S=%d E= %d W=%d", s.facing,
		s.distances[N], s.distances[S], s.distances[E], s.distances[W])
}

func (s *ship) turnRight(degrees int) {
	for i := 0; i < degrees; i = i + 90 {
		switch s.facing {
		case "N":
			s.facing = E
		case "S":
			s.facing = W
		case "E":
			s.facing = S
		case "W":
			s.facing = N
		}
	}
}

func (s *ship) turnLeft(degrees int) {
	for i := 0; i < degrees; i = i + 90 {
		switch s.facing {
		case "N":
			s.facing = W
		case "S":
			s.facing = E
		case "E":
			s.facing = N
		case "W":
			s.facing = S
		}
	}
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
		d.actions = append(d.actions, action{
			code: t[0:1],
			val:  utils.Atoi(t[1:]),
		})
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`F10
N3
F7
R90
F11
`)
}
