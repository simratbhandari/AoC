package day01


// Time complexity of this function if O(n).
func SumPairs(A []int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		next := (i + 1) % len(A)
		if A[i] == A[next] {
			sum += A[i]
		}
	}
	return sum
}


// Time complexity of this function if O(n).
func SumSplitPairs(A []int) int {
	if len(A)%2 != 0 {
		panic("array length is not even")
	}
	sum := 0
	for i := 0; i < len(A); i++ {
		next := (i + len(A)/2) % len(A)
		if A[i] == A[next] {
			sum += A[i]
		}
	}
	return sum
}