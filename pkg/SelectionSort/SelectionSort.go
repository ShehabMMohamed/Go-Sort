package SelectionSort

func Sort(A []int) {
	var min_index int
	for i := 0; i < len(A) - 1; i++ {
		min_index = i 
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[min_index] {
				min_index = j 
			}
		}
		A[i], A[min_index] = A[min_index], A[i]
	}
}