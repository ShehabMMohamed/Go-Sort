package BubbleSort


func Sort(A []int) {
	isSwapped := false
	for i := 0; i < len(A) - 1; i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				A[i], A[j] = A[j], A[i]
				isSwapped = true
			}
		}
		if !isSwapped { 	// Micro-optimization, return if already sorted from first scan.
			return
		}
	}
}