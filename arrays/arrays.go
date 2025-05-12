package arrays

func Sum(nos []int) int {
	sum := 0
	for _, val := range nos {
		sum += val
	}
	return sum
}
