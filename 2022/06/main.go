package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputFile = "input.txt"

func main() {
	solution()
}

type possibleVals interface {
	rune
}

func indexOf[T possibleVals](s []T, v T) int {
	for i, r := range s {
		if r == v {
			return i
		}
	}

	return -1
}

func solution() {
	f, err := os.Open(inputFile)
	if err != nil {
		panic("could not read input")
	}

	sc := bufio.NewScanner(f)

	sc.Split(bufio.ScanLines)
	sc.Scan()

	text := sc.Text()

	prev := []rune{}
	max := 14

	for index, current := range text {
		i := indexOf(prev, current)
		fmt.Println(index, string(current), string(prev))
		prev = append(prev, current)

		if i < 0 {

			if len(prev) != max {
				continue
			}

			fmt.Println(string(prev), index+1)
			return
		}

		prev = prev[i+1:]
	}
}
