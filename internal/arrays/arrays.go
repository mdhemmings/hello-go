package array

// Sum : Adds numbers together passed into it as an array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll : Adds multiple slices together and returns result
func SumAll(numbers1 []int, numbers2 []int) int {
	sum := 0
	for _, number := range numbers1 {
		sum += number
	}
	for _, number := range numbers2 {
		sum += number
	}
	return sum
}

// SumReflect : Shows the totals of the slices passed into it
func SumReflect(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
