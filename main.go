package main

import (
	"Go-Sort/pkg/BubbleSort"
	"Go-Sort/pkg/InsertionSort"
	"Go-Sort/pkg/SelectionSort"

	"fmt"
)

func main() {
	A := []int{5, 4, 2, 3, 1}
	B := []int{5, 4, 2, 3, 1}
	C := []int{5, 4, 2, 3, 1}


	fmt.Printf("Unsorted -> %+v\n", A)
	BubbleSort.Sort(A)
	fmt.Printf("Bubble Sort -> %+v\n\n", A)

	fmt.Printf("Unsorted -> %+v\n", B)
	SelectionSort.Sort(B)
	fmt.Printf("Selection Sort -> %+v\n\n", B)

	fmt.Printf("Unsorted -> %+v\n", C)
	InsertionSort.Sort(C)
	fmt.Printf("Insertion Sort -> %+v\n\n", C)
}