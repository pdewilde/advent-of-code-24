package main

import (
	"fmt"
	"slices"
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
	for _, v := range unsorted {
		v := v // if you are running old golang
		go func (v int)  {
			time.Sleep(time.Duration(v * 1) * time.Millisecond)
			sortedCh <- v
		}(v)
	}
	for i := 0; i < len(unsorted); i++ {
		sorted = append(sorted, <-sortedCh)
	}
	return sorted
}

func main() {
	// test sorted
	unsorted := []int{5, -4, 4, 4, 0, 1, 100, 900}
	sorted1 := constSort(unsorted)
    fmt.Println(sorted1)
	slices.Sort(unsorted)
    fmt.Println(unsorted)
}
