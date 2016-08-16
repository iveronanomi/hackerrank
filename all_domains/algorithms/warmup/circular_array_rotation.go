package warmup

import "fmt"

//Sherlock and Watson
func CircularArrayRotation() {
	var N, K, Q int
	fmt.Scanf("%d %d %d", &N, &K, &Q)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	X := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Scan(&X[i])
	}
	if C := K / N; C > 1 {
		K = K - N*C
	}
	A = append(A[N-K:], A[:N-K]...)
	for _, idx := range X {
		fmt.Println(A[idx])
	}
}
