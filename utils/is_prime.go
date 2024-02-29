package utils

func countPrime(n int) int {
	res := 0
	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			res++
			for j := i * 2; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return res
}
