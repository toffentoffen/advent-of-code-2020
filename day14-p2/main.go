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
	var memory = map[uint64]uint64{}
	for _, mo := range d.maskOps {
		for _, wm := range mo.memoryWrites {
			wm.address = applyMask(mo.mask, wm.address)
			for _, m := range getFloatMasks(mo.mask) {
				addr := wm.address | uint64(m)
				memory[addr] = wm.value
			}

		}
	}

	var sum uint64
	for _, v := range memory {
		sum += v
	}

	fmt.Printf("sum: %d\n", sum)
}

type memoryWrite struct {
	address uint64
	value   uint64
}
type maskOp struct {
	mask         string
	memoryWrites []memoryWrite
}
type data struct {
	maskOps []maskOp
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
	var mop maskOp
	var i int = -1
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "mask = ") {
			mop = maskOp{
				mask: strings.TrimPrefix(line, "mask = "),
			}
			d.maskOps = append(d.maskOps, mop)
			i++
			continue
		}
		var mw memoryWrite
		utils.Sscanf(line, "mem[%d] = %d", &mw.address, &mw.value)
		d.maskOps[i].memoryWrites = append(d.maskOps[i].memoryWrites, mw)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`)
}

func applyMask(mask string, value uint64) uint64 {
	var l = len(mask)
	for i := l - 1; i >= 0; i-- {
		if mask[i] == '1' {
			value = setBit(value, uint64(l-1-i))
		}
		if mask[i] == 'X' {
			value = clearBit(value, uint64(l-1-i))
		}
	}
	return value
}

// 1X0X => [0000, 0001, 0100, 0101] => [0, 1, 4, 5]
func getFloatMasks(s string) []int {
	a := []int{}
	for idx, char := range s {
		if char == 'X' {
			a = append(a, 1<<(len(s)-idx-1))
		}
	}

	b := []int{0}
	for _, i := range a {
		for _, j := range b {
			b = append(b, i|j)
		}
	}

	sort.IntSlice(b).Sort()
	return b
}

// Sets the bit at pos in the integer n.
func setBit(n uint64, pos uint64) uint64 {
	n |= (uint64(1) << pos)
	return n
}

// Clears the bit at pos in n.
func clearBit(n uint64, pos uint64) uint64 {
	mask := ^(uint64(1) << pos)
	n &= mask
	return n
}
