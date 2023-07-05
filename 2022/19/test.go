package main

import "fmt"

func main2() {
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := arr1

	arr1[0] = 5

	fmt.Println(arr1)
	fmt.Println(arr2)
}
