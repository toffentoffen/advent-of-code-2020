package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	_, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
}

type data struct{}

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
		// do stuff
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}