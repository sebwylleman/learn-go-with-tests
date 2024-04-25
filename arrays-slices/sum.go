package main

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int
	for _, v := range numbersToSum {
		sum = append(sum, Sum(v))
	}
	return sum
}
