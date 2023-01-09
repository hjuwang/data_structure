package main

//生成斐波那契数列数组

var F []int

func CreateFibonacci(n int) {

	for i := 0; i < n; i++ {
		F = append(F, Fibonacci(i))
	}
}

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
