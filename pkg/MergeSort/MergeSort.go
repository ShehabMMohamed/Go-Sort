package MergeSort

func Sort(A []int) {
	mergeSort(A, 0, len(A) - 1)
}

func mergeSort(A[]int, left int, right int) {
	if left >= right {
		return
	}
	middle := left + (right - left) / 2
	mergeSort(A, left, middle)
	mergeSort(A, middle+1, right)
	merge(A, left, middle, right)
}

func merge(A[]int, left int, middle int, right int) {
	var i, j, k int 
	k = 0
	n1 := middle - left + 1
	n2 := right - middle 
	L := make([]int, n1)
	R := make([]int, n2)

	for i = 0; i < n1; i++ {
		L[i] = A[left + i]
	}
	for j = 0; j < n2; j++ {
		R[j] = A[middle + 1 + j]
	}

	i = 0
	j = 0

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			A[left + k] = L[i]
			i++
		} else {
			A[left + k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		A[left + k] = L[i]
		i++
		k++
	}

	for j < n2 {
		A[left + k] = R[j]
		j++
		k++
	}
}
