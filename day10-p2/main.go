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
	sort.Ints(d.adapters)
	arrangements := d.adapterArrangements(0, 0, make(map[int]int))
	fmt.Printf("arrangements: %d\n", arrangements)
}

type data struct {
	adapters       []int
	buildInAdapter int
}

func (d data) adapterArrangements(currentJolts, adapter int, seenArrangements map[int]int) int {
	if adapter == len(d.adapters) {
		return 1
	}
	var arrangements int
	for i, a := range d.adapters[adapter:] {
		if a-currentJolts > 3 {
			break
		}
		if seenArrangements[adapter+1] == 0 {
			seenArrangements[adapter+i] = d.adapterArrangements(a, adapter+i+1, seenArrangements)
		}
		arrangements += seenArrangements[adapter+i]
	}
	return arrangements
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
		d.adapters = append(d.adapters, utils.Atoi(s.Text()))
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	d.buildInAdapter = 3
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`)
}
