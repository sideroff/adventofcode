package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1")
	partOne()
	fmt.Println()
	fmt.Println("Part 2")
	partTwo()
}

func partOne() {
	cubeCountByColor := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)

	validGamesCount := 0
	cache := map[string]bool{}
	cacheHits := 0

	for i := 0; ; i++ {
		sc.Scan()
		line := sc.Text()

		if len(line) == 0 {
			break
		}

		hasImpossibleHand := false
		gameParts := strings.Split(line, ":")
		gameId, err := strconv.Atoi(strings.Split(strings.Trim(gameParts[0], " "), " ")[1])
		if err != nil {
			panic(err)
		}

		draws := strings.Split(gameParts[1], ";")
		for _, draw := range draws {
			drawnCubes := strings.Split(draw, ", ")

			for _, drawnCube := range drawnCubes {
				drawnCube = strings.Trim(drawnCube, " ")
				possible, ok := cache[drawnCube]
				if ok {
					cacheHits++
					if possible {
						continue
					} else {
						hasImpossibleHand = true
						break
					}
				}

				cubeParts := strings.Split(drawnCube, " ")
				numberOfCubes, err := strconv.Atoi(cubeParts[0])
				if err != nil {
					panic(err)
				}

				color := cubeParts[1]
				hasImpossibleHand = numberOfCubes > cubeCountByColor[color]
				cache[drawnCube] = !hasImpossibleHand

				if hasImpossibleHand {
					break
				}
			}
		}

		if !hasImpossibleHand {
			validGamesCount += gameId
		}
	}

	fmt.Println("valid games =", validGamesCount)
	fmt.Println("cache hits  =", cacheHits)
}

func partTwo() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)
	sum := 0

	for i := 0; ; i++ {
		sc.Scan()
		line := sc.Text()

		if len(line) == 0 {
			break
		}

		minCubeCountByColor := map[string]int{}

		gameParts := strings.Split(line, ":")
		if err != nil {
			panic(err)
		}

		draws := strings.Split(gameParts[1], ";")
		for _, draw := range draws {
			drawnCubes := strings.Split(draw, ", ")

			for _, drawnCube := range drawnCubes {
				drawnCube = strings.Trim(drawnCube, " ")

				cubeParts := strings.Split(drawnCube, " ")
				numberOfCubes, err := strconv.Atoi(cubeParts[0])
				if err != nil {
					panic(err)
				}

				color := cubeParts[1]
				minCubeCountByColor[color] = max(minCubeCountByColor[color], numberOfCubes)
			}
		}

		gamePowerSum := 1
		for _, v := range minCubeCountByColor {
			gamePowerSum *= v
		}

		if gamePowerSum == 1 {
			continue
		}

		sum += gamePowerSum
	}

	fmt.Println("game power sum", sum)
}
