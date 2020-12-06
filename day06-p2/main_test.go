package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLettersToMask(t *testing.T) {
	var mask int64 = 0
	for _, l := range "afalhyue" {
		c := l - 97
		mask = mask | (1 << c)
	}
	fmt.Printf("%d, %s\n", mask, strconv.FormatInt(mask, 2))
}