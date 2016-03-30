package main

import "fmt"

const (
	MaxInt64 = 1<<63 - 1
	MinInt64 = -1 << 63
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) int {
	max := MinInt64
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	min := MaxInt64
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min
}

func main() {
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
	fmt.Println(max(values...)) // "4"
	fmt.Println(min(values...)) // "1"
	fmt.Println(max())          // "MinInt"
	fmt.Println(min())          // "MaxInt"
}
