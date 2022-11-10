package main

import "testing"

// go test --coverprofile=coverage.out
// go tool cover --func=coverage.out
// go tool cover --html=coverage.out

// go test --cpuprofile=cpu.out
// go tool pprof cpu.out

func TestSum(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := sumToTest(item.a, item.b)

		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{3, 5, 5},
	}

	for _, item := range tables {
		max := getMaxToTest(item.a, item.b)

		if max != item.n {
			t.Errorf("Max was incorrect, got %d expected %d", max, item.n)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
	}

	for _, item := range tables {
		fib := fibonacciToTest(item.a)

		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got %d expected %d", fib, item.n)
		}
	}

}
