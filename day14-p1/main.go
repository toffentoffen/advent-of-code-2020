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
	var memory = map[int]uint64{}
	for _, mo := range d.maskOps {
		for _, wm := range mo.memoryWrites {
			memory[wm.address] = applyMask(mo.mask, wm.value)
		}
	}

	var sum uint64
	for _, v := range memory {
		sum += v
	}

	fmt.Printf("sum: %d\n", sum)
}

type memoryWrite struct {
	address int
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
	return strings.NewReader(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`)
}

func applyMask(mask string, value uint64) uint64 {
	var l = len(mask)
	for i := l - 1; i >= 0; i-- {
		if mask[i] == '0' {
			value = clearBit(value, uint64(l-1-i))
		}
		if mask[i] == '1' {
			value = setBit(value, uint64(l-1-i))
		}
	}
	return value
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
