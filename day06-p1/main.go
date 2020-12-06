package main

import (
	"advent-of-code-2020/utils"
	"bufio"
	"fmt"
	"io"
	"log"
	"math/bits"
	"os"
	"strings"
)

func main() {
	d, err := readInput("input.txt")
	//d, err := read(testInputReader())
	if err != nil {
		log.Fatal(err)
	}

	sum := utils.Sum(len(d.groups), func(i int) int {
		var mask uint32
		for _, a := range d.groups[i] {
			mask = mask | uint32(a)
		}
		ones := bits.OnesCount32(mask)
		fmt.Printf("g %d, ones %d\n", i, ones)
		return ones
	})
	fmt.Println("sum: ", sum)
}

type answers []int32
type data struct{
	groups []answers
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
	var as answers
	for s.Scan() {
		if s.Text() == ""{
			d.groups = append(d.groups, as)
			as = answers{}
			continue
		}
		m := stringToMask(s.Text())
		as = append(as, m)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`abc

a
b
c

ab
ac

a
a
a
a

b

`)
}

func stringToMask(s string) int32 {
	var mask int32 = 0
	for _, l := range s {
		c := l - 97
		mask = mask | (1 << c)
	}
	return mask
}