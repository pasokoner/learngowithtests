package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {

	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, nums := range numbersToSum {

		if len(nums) == 0 {
			sums[i] = 0
		} else {
			sums[i] = Sum(nums[1:])
		}
	}

	return sums
}
