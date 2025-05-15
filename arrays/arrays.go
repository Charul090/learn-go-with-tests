package arrays

func Sum(nos []int) int {
	sum := 0
	for _, val := range nos {
		sum += val
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sumOfNos := make([]int, lengthOfNumbers)
	for i, numbers := range numbersToSum {
		sumOfNos[i] = Sum(numbers)
	}

	return sumOfNos
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, nos := range numbersToSum {
		if len(nos) == 0 {
			sums = append(sums, 0)
		} else {
			slice := nos[1:]
			sums = append(sums, Sum(slice))
		}
	}

	return sums
}
