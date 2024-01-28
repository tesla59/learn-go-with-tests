package arrayandslices

func Sum(num []int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	answer := []int{}
	for _, v := range slices {
		answer = append(answer, Sum(v))
	}
	return answer
}
