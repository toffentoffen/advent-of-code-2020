package main

import (
	"advent-of-code-2020/utils"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	d, err := readInput("input.txt")
	//d, err := read(testInputReader())
	if err != nil {
		log.Fatal(err)
	}
	var lastTrips []departure
	for _, bus := range d.buses {

		lastTrips = append(lastTrips, departure{
			((d.arrival / bus) * bus) + bus,
			bus,
		})
	}
	fmt.Printf("trips: %v\n", lastTrips)
	sort.Slice(lastTrips, func(i, j int) bool {
		return lastTrips[i].departed < lastTrips[j].departed
	})
	fmt.Printf("trips: %v\n", lastTrips)
	var solution = lastTrips[0].bus * (lastTrips[0].departed - d.arrival)
	fmt.Printf("solution: %d\n", solution)
}

type departure struct {
	departed int
	bus      int
}
type data struct {
	arrival int
	buses   []int
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
	s.Scan()
	d.arrival = utils.Atoi(s.Text())
	s.Scan()
	values := strings.Split(s.Text(), ",")
	for _, v := range values {
		if v != "x" {
			d.buses = append(d.buses, utils.Atoi(v))
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`939
7,13,x,x,59,x,31,19`)
}
