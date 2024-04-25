package main

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sum := make([]int, len(numbersToSum))
	for i, v := range numbersToSum {
		sum[i] = Sum(v)
	}
	return sum
}
