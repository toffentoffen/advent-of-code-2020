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
	var result, mode int64 = 1, 1
	for _, bus := range d.buses {
		for (result + bus.offset) % bus.bus != 0 {
			result += mode
		}
		mode *= bus.bus
	}
	fmt.Printf("result: %d\n", result)
}

type departure struct {
	offset int64
	bus    int64
}
type data struct {
	arrival int
	buses   []departure
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
	for i, v := range values {
		if v != "x" {
			d.buses = append(d.buses, departure{
				offset: int64(i),
				bus:    int64(utils.Atoi(v)),
			})
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
