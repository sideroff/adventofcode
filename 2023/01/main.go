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

func isAsciiCodeANumber(asciiCode byte) bool {
	return (asciiCode >= 48 || asciiCode <= 57)
}

func partOne(inputFile string) {
	f, err := os.OpenFile(inputFile, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	if err != nil {
		panic(err)
	}

	sum := 0

	scanner.Scan()
	line := []rune(scanner.Text())
	lineLen := len(line)

	for lineLen != 0 {
		li := -1
		ri := -1

		for i := 0; i < lineLen/2; i++ {
			l := line[i]
			r := line[lineLen-i-1]
			if li == -1 && unicode.IsNumber(l) {
				li = i
			}

			if ri == -1 && unicode.IsNumber(r) {
				ri = lineLen - i - 1
			}

			if li != -1 && ri != -1 {
				str := string(line[li]) + string(line[ri])
				num, err := strconv.Atoi(str)
				if err != nil {
					panic(err)
				}

				sum += num
				break
			}
		}

		scanner.Scan()
		line = []rune(scanner.Text())
		lineLen = len(line)
	}

	fmt.Println("sum", sum)
}

func partTwo(inputFile string) {
	f, err := os.OpenFile(inputFile, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	if err != nil {
		panic(err)
	}

	digitWords := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	sum := 0

	scanner.Scan()
	line := []rune(scanner.Text())
	lineLen := len(line)

	for lineLen != 0 {
		fmt.Println("line", string(line))
		// left digit/word index
		ldi := -1
		lwi := -1

		// left digitWords index
		ldwi := -1

		// right digit/word index
		rdi := -1
		rwi := -1

		// right digitWords index
		rdwi := -1

		for i := 0; i < lineLen; i++ {
			l := line[i]
			r := line[lineLen-i-1]
			if ldi == -1 && unicode.IsNumber(l) {
				ldi = i
				fmt.Println("-- found l digit", string(l), "i", ldi)
			}

			if rdi == -1 && unicode.IsNumber(r) {
				rdi = lineLen - i - 1
				fmt.Println("-- found r digit", string(r), "i", rdi)
			}

			if lwi == -1 {
				for dwi, word := range digitWords {
					if strings.HasSuffix(string(line[:i]), word) {
						ldwi = dwi
						lwi = i - len(word)
						fmt.Println("-- found l word", word, "i", lwi)
					}
				}
			}

			if rwi == -1 {
				for dwi, word := range digitWords {
					if strings.HasPrefix(string(line[lineLen-i:]), word) {
						rdwi = dwi
						rwi = lineLen - i
						fmt.Println("-- found r word", word, "i", lineLen-i)

					}
				}
			}
		}

		firstDigit := ""
		lastDigit := ""

		if ldi < lwi && ldi != -1 || lwi == -1 {
			firstDigit = string(line[ldi])
		} else {
			firstDigit = strconv.Itoa(ldwi)
		}

		if rdi > rwi && rdi != -1 || rwi == -1 {
			lastDigit = string(line[rdi])
		} else {
			lastDigit = strconv.Itoa(rdwi)
		}

		fmt.Println("-- found digits for line", string(line), "", firstDigit, lastDigit)
		str := firstDigit + lastDigit
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		sum += num

		math.Max(1, 2.3)
		scanner.Scan()
		line = []rune(scanner.Text())
		lineLen = len(line)
	}

	fmt.Println("sum", sum)
}

func main() {
	inputFile := "./input.txt"

	partTwo(inputFile)
}
