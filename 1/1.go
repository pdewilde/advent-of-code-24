package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// We have two lists.
// 1, 3, 3, 5
// 2, 3, 3, 5
// We want
// 1  0  0  0

// Step 1: Sort lists.
// Naieve solution would be a quicksort.
// I think I can improve.
// N log n is for loosers
// Constant Sort O(max(n)) (ignore the O(n) components, they are inconvienet to this statement):
func constSort(unsorted []int) (sorted []int) {
	sortedCh := make(chan int, len(unsorted))
	// Get min value to handle negatives as well as optimize for cases where all numbers are large.
	min := math.MaxInt64
	for _, v := range unsorted {
		if v < min {
			min = v
		}
	}
	for _, v := range unsorted {
		v := v // if you are running old golang
		go func(v int) {
			time.Sleep(time.Duration(v-min) * time.Millisecond)
			sortedCh <- v
		}(v)
	}
	for i := 0; i < len(unsorted); i++ {
		sorted = append(sorted, <-sortedCh)
	}
	return sorted
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("it's advent of code!")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var list1 []int
	var list2 []int

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) != 2 {
			panic("no effort was used in this project")
		}
		parsed, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(":(")
		}
		list1 = append(list1, parsed)
		parsed, err = strconv.Atoi(parts[1])
		if err != nil {
			panic(":(")
		}
		list2 = append(list2, parsed)
	}

	sorted1 := constSort(list1)
	sorted2 := constSort(list2)
	cumulative := 0
	for i := 0; i < len(sorted1); i++ {
		var diff int
		// Its dumb there is no abs for int
		if sorted1[i] > sorted2[i] {
			diff = sorted1[i] - sorted2[i]
		} else {
			diff = sorted2[i] - sorted1[i]
		}
		cumulative = cumulative + diff
	}
	fmt.Println(cumulative)
}
