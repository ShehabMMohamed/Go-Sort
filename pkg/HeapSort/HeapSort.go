package HeapSort

func heapify(A []int, N int, i int) {
	largest := i 
	leftChild := (i * 2) + 1
	rightChild := (i * 2) + 2
	if leftChild < N && A[largest] < A[leftChild] {
		largest = leftChild
	}
	if rightChild < N && A[largest] < A[rightChild] {
		largest = rightChild
	}

	if largest != i {
		A[largest], A[i] = A[i], A[largest]
		heapify(A, N, largest)
	}
}

func Sort(A []int) {
	// build MaxHeap first.
	N := len(A)
	for i := (N/2) - 1; i >= 0; i-- {
		heapify(A, N, i)
	}

	for i := N - 1; i >= 0; i-- {
		A[0], A[i] = A[i], A[0]
		heapify(A, i, 0)
	}
}