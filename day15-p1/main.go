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
	//d, err := readInput("input.txt")
	d, err := read(testInputReader())
	if err != nil {
		log.Fatal(err)
	}
	var lastSpoken int
	var round = 1
	for _, s := range d.starting {
		d.spokenAt[s] = []int{round}
		lastSpoken = s
		round++
	}


	for round <= 2020 {
		if len(d.spokenAt[lastSpoken]) <= 1 {
			d.spokenAt[0] = append([]int{round}, d.spokenAt[0]...)
			lastSpoken = 0
		} else {
			age := d.spokenAt[lastSpoken][0] - d.spokenAt[lastSpoken][1]
			d.spokenAt[age] = append([]int{round}, d.spokenAt[age]...)
			lastSpoken = age
		}
		round++
	}
	fmt.Printf("lastspoken: %d\n", lastSpoken)
}

type data struct {
	starting []int
	spokenAt map[int][]int // key is number, value is last round seen
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
		spokenAt: make(map[int][]int),
	}

	s := bufio.NewScanner(r)
	for s.Scan() {
		parts := utils.Split(s.Text(), ",")
		for _, p := range parts {
			n := utils.Atoi(p)
			d.starting = append(d.starting, n)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`13,16,0,12,15,1`)
}
