package Algorithm

import (
	"Go-Sort/pkg/BubbleSort"
	"Go-Sort/pkg/HeapSort"
	"Go-Sort/pkg/InsertionSort"
	"Go-Sort/pkg/SelectionSort"
	"Go-Sort/pkg/constants"
	"fmt"
)

func New(algorithm string) func(A[]int) {
	switch algorithm {
	case constants.Bubble_Sort:
		return BubbleSort.Sort
	case constants.Insertion_Sort:
		return InsertionSort.Sort
	case constants.Selection_Sort:
		return SelectionSort.Sort
	case constants.Heap_Sort:
		return HeapSort.Sort
	default:
		fmt.Printf("Could not find Algorithm: %s\n", algorithm)
		return func(A []int) { 
			fmt.Println("Could not call Sort.")
		}
	}
}