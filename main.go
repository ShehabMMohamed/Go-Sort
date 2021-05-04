package main

import (
	"Go-Sort/pkg/Algorithm"
	"Go-Sort/pkg/constants"

	"fmt"
)

func main() {
	test_cases := [][]int{ 
	{5, 4, 3, 2, 1},
	{1, 1, 1, 0, 1, 0}, 
	{10, 40, 30, 20, 50}}

	algorithms := []string{
		constants.Bubble_Sort, constants.Selection_Sort, constants.Insertion_Sort,
		constants.Heap_Sort,
	}
	for _, sorting_alg := range algorithms {
		fmt.Printf("\nRunning with Sorting Algorithm - %s\n", sorting_alg)
		for i, test_case := range test_cases {
			var A = make([]int, len(test_case))
			copy(A, test_case)
			fmt.Printf("(%d) Before -> %+v\n", i, A)
			Sort := Algorithm.New(sorting_alg)
			Sort(A)
			fmt.Printf("(%d) After -> %+v\n", i, A)	
		}
	}
}