package array_slice

func Sum(numbers []int) int {

	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(arr ...[]int) []int {

	var res []int

	for _, num := range arr {
		res = append(res, Sum(num))
	}

	return res
}

func SumTails(arr ...[]int) []int {
	var res []int

	for _, num := range arr {

		if len(num) == 0 {
			res = append(res, 0)
		} else {
			tail := num[1:]
			res = append(res, Sum(tail))
		}
	}
	return res
}
