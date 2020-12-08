package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	p, err := readInput("input.txt")
	//p, err := read(testInputReader())
	if err != nil {
		log.Fatal(err)
	}
	acc := p.run()
	fmt.Printf("acc: %d\n", acc)
}

type instruction struct {
	operation string
	value     int64
	executed  bool
}
type program struct {
	instructions []instruction
	accumulator  int64
}

func (p program) run() int64 {
	var ic int64
	for !p.instructions[ic].executed {
		p.instructions[ic].executed = true
		switch p.instructions[ic].operation {
		case "nop":
			ic++
		case "acc":
			 p.accumulator += p.instructions[ic].value
			ic++
		case "jmp":
			ic += p.instructions[ic].value
		}

	}
	return p.accumulator
}

func readInput(path string) (*program, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()
	return read(f)
}

func read(r io.Reader) (*program, error) {
	var p program

	s := bufio.NewScanner(r)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		val, err := strconv.ParseInt(fields[1], 10, 16)
		if err != nil {
			log.Fatalf("could not parse int %s: %v", fields[1], err)
		}
		p.instructions = append(p.instructions, instruction{
			operation: fields[0],
			value:     val,
		})
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &p, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)
}
