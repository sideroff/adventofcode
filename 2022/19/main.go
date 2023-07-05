package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type State struct {
	resources [4]int
	robots    [4]int
	path      string
}

const inputFile = "./input_small.txt"
const turns = 10
const mostWantedResource = 3

var cache = make(map[string]State)

var calculations = 0
var cacheHits = 0

func parseInt(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return res
}

func main() {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	str := string(bytes)
	re := regexp.MustCompile("[0-9]+")

	for _, line := range strings.Split(str, "\n") {
		// simple parsing, but works for task
		found := re.FindAllString(line, -1)

		resources := [4]int{0, 0, 0, 0}
		robots := [4]int{1, 0, 0, 0}
		costs := [4][4]int{
			{parseInt(found[1]), 0, 0, 0},
			{parseInt(found[2]), 0, 0, 0},
			{parseInt(found[3]), parseInt(found[4]), 0, 0},
			{parseInt(found[5]), 0, parseInt(found[6]), 0},
		}

		endResources, _, path := solve(turns, resources, robots, costs, "")

		fmt.Printf("- end resources    %v %s \n", endResources, path)
		fmt.Printf("- end calculations %d \n", calculations)
		fmt.Printf("- end cache hits   %d \n", cacheHits)
	}
}

func solve(turnsLeft int, resources [4]int, robots [4]int, costs [4][4]int, path string) (newResources [4]int, newRobots [4]int, outputPath string) {
	calculations += 1
	hit, cResources, cRobots, cPath := getCache(turnsLeft, resources, robots)
	if hit {
		return cResources, cRobots, cPath
	}

	lresources := resources
	lrobots := robots

	maxResources := resources
	maxRobots := resources
	bestPath := path

	for i := 1; i <= turnsLeft; i++ {
		lresources = getResources(lresources, lrobots)

		lrobots = robots
		// fmt.Printf("%s %d:%d -income resources %v robots %v\n", path, turnsLeft, i, lresources, lrobots)

		// try waiting path
		// fmt.Printf("%s %d:%d waiting\n", path, turnsLeft, i)
		wResources, wRobots, theirPath := solve(turnsLeft-i, lresources, lrobots, costs, path+"w, ") // waiting
		// fmt.Printf("%s %d:%d resource %v robots %v gotten by waiting %d turns with %v robots\n", path, turnsLeft, i, wResources, wRobots, turnsLeft-i, robots)

		if maxResources[mostWantedResource] < wResources[mostWantedResource] {
			maxResources = wResources
			maxRobots = wRobots
			bestPath = theirPath
		}

		// for type of ore
		for j := 0; j < 4; j++ {
			// fmt.Println("testing ores")
			canBuildRobot, newResources, newRobots := tryBuyingRobot(j, lresources, lrobots, costs)
			if !canBuildRobot {
				continue
			}

			// fmt.Printf("%d:%d can build r %d with resources %v \n", turnsLeft, i, j, lresources)

			// solve for case
			endResources, endRobots, theirPath := solve(turnsLeft-i, newResources, newRobots, costs, fmt.Sprintf("%sb%d, ", path, j))

			// see if solution is better than current max
			if maxResources[mostWantedResource] < endResources[mostWantedResource] {
				maxResources = endResources
				maxRobots = endRobots
				bestPath = theirPath
			}
		}
	}

	// setCache(turnsLeft, maxResources, maxRobots, bestPath)
	return maxResources, maxRobots, bestPath
}

func getResources(resources [4]int, robots [4]int) (newResources [4]int) {
	for i, output := range robots {
		resources[i] += output
	}

	return resources
}

func getCache(turnsLeft int, resources [4]int, robots [4]int) (hit bool, cResources [4]int, cRobots [4]int, cPath string) {
	key := fmt.Sprintf("%d,%v,%v", turnsLeft, resources, robots)
	val, ok := cache[key]
	if ok {
		cacheHits += 1
		return true, val.resources, val.robots, val.path
	}
	return false, [4]int{}, [4]int{}, ""
}

func setCache(turnsLeft int, resources [4]int, robots [4]int, path string) bool {
	key := fmt.Sprintf("%d,%v,%v", turnsLeft, resources, robots)
	_, ok := cache[key]
	if ok {
		return true
	}

	cache[key] = State{
		resources,
		robots,
		path,
	}
	return true
}

func tryBuyingRobot(index int, resources [4]int, robots [4]int, costs [4][4]int) (success bool, newResources [4]int, newRobots [4]int) {
	cost := costs[index]
	for i, v := range cost {
		resources[i] = resources[i] - v
		if resources[i] < 0 {
			d := [4]int{}
			return false, d, d
		}
	}

	robots[index] += 1
	return true, resources, robots
}
