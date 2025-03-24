package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestMaxInt(t *testing.T) {
	a, b := 2, 7

	res := MaxInt(a, b)

	if res != b {
		t.Errorf("expected %d, got %d", b, res)
	}
}

func TestRace(t *testing.T) {
	var builder strings.Builder
	counter := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	res := Race()

	for _, num := range counter {
		builder.WriteString(strconv.Itoa(num))
	}

	counterStr := builder.String()

	builder.Reset()

	for _, num := range res {
		builder.WriteString(strconv.Itoa(num))
	}

	resStr := builder.String()

	if resStr != counterStr {
		t.Errorf("expected %s, got %s", counterStr, resStr)
	}
}
