package QuickSort

func Sort(A[] int) {
	quickSort(A, 0, len(A) - 1)
}

func quickSort(A[]int, p int, r int) {
	if p >= r {
		return
	}
	q := partition(A, p, r)
	quickSort(A, p, q-1)
	quickSort(A, q+1, r)
}

func partition(A[]int, p int, r int) int {
	x := A[r]
	i := p - 1 
	for j := p; j <= r - 1; j++ {
		if A[j] < x {
			i++ 
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}
