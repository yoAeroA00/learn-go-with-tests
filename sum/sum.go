package sum

func Sum(nums [5]int) int {
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Sum2(nums []int) int {
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int

	for _, num := range numbersToSum {
		sum = append(sum, Sum2(num))
	}
	return sum
}

func SumAllTails(nums ...[]int) []int {
	var sum []int

	for _, num := range nums {
		if len(num) != 0 {
			sum = append(sum, Sum2(num[1:]))
		} else {
			sum = append(sum, 0)
		}
	}
	return sum
}
