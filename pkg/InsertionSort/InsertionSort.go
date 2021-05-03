package InsertionSort

func Sort(A []int) {
	var key, j int
	for i := 1; i < len(A); i++ {
		key = A[i]
		j = i - 1 
		for j >= 0 && A[j] > key {
			A[j+1] = A[j]
			j--
		}
		A[j + 1] = key
	}
}