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
	diffs := d.joltDifferences()
	fmt.Printf("differences: %v\n", diffs)
	fmt.Printf("1 jolt diffs * 3 jolt diffs: %d\n", diffs[1]*diffs[3])
}

type data struct{
	adapters []int
	buildInAdapter int
}

func (d data) joltDifferences() map[int]int {
	sort.Ints(d.adapters)
	var diffs = make(map[int]int)
	var jolts int
	for i, a := range d.adapters {
		dif := a - jolts
		if dif > 3 {
			log.Fatalf("jolt difference is higher than 3 (last adapater %d, current adapter %d ) at adapter %d", jolts, a, i)
		}
		jolts = a
		diffs[dif]++
	}
	diffs[d.buildInAdapter]++
	return diffs
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