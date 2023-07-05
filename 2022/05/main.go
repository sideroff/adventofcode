package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const inputFile = "input.txt"

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	solution(9001)
}

func solution(craneMoverLevel int) {
	f, err := os.Open(inputFile)
	if err != nil {
		panic("could not read input")
	}

	sc := bufio.NewScanner(f)

	sc.Split(bufio.ScanLines)

	var n int
	var stacks [][]string

	for sc.Scan() {
		line := sc.Text()

		// number of stacks
		if n == 0 {
			n = (len(line) + 1) / 4
			stacks = make([][]string, n)
			for i := 0; i < n; i++ {
				stacks[i] = make([]string, 0)
			}
		}

		isBoxLine := strings.Contains(line, "[")
		isCmdLine := strings.Contains(line, "move")
		isNewLine := len(line) < 1

		if isBoxLine {
			// parse box line
			lineAsRunes := []rune(line)
			for i := 1; i < len(lineAsRunes); i += 4 {
				stackIndex := (i - 1) / 4
				v := lineAsRunes[i]

				// if not letter, don't push to stack
				if !unicode.IsLetter(v) {
					continue
				}

				stacks[stackIndex] = append(stacks[stackIndex], string(v))
			}
		}

		if isNewLine {
			for i := 0; i < len(stacks); i++ {
				reverse(stacks[i])
			}
		}

		if isCmdLine {
			cmd := strings.Replace(line, "move ", "", -1)
			cmd = strings.Replace(cmd, " from", "", -1)
			cmd = strings.Replace(cmd, " to", "", -1)

			cmdParams := strings.Split(cmd, " ")

			n, err := strconv.Atoi(cmdParams[0])
			if err != nil {
				fmt.Println("could not parse cmd", line, err)
			}

			from, err := strconv.Atoi(cmdParams[1])
			if err != nil {
				fmt.Println("could not parse cmd", line, err)
			}

			to, err := strconv.Atoi(cmdParams[2])
			if err != nil {
				fmt.Println("could not parse cmd", line, err)
			}

			sf := stacks[from-1]
			safeFromIndex := int(math.Max(float64(len(sf)-n), 0.0))
			stacks[from-1] = sf[0:safeFromIndex]

			d := sf[safeFromIndex:]

			if craneMoverLevel < 9000 {
				reverse(d)
			}

			st := stacks[to-1]
			stacks[to-1] = append(st, d...)
		}
	}

	for _, v := range stacks {
		fmt.Print(v[len(v)-1])
	}
	fmt.Println()
}
