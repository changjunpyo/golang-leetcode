package problems

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	first := 0
	second := 1
	for i := 2; i <= n; i++ {
		next := first + second
		first, second = second, next
	}
	return second
}
