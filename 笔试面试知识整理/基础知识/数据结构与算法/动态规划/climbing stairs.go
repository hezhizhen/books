package main

// dp[i] = dp[i-1] + dp[i-2]
// i=1: 1
// i=2: 2 (1+1, 2)
// i=3: 3 (1+1+1, 1+2, 2+1)
func climbStairs(n int) int {
	f0 := 1
	f1 := 1
	for i := 2; i <= n; i++ {
		f2 := f0 + f1
		f0 = f1
		f1 = f2
	}
	return f1
}
