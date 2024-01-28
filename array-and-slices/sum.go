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

func SumAllTails(slices ...[]int) []int {
	answer := []int{}
	for _, v := range slices {
		if len(v) == 0 {
			answer = append(answer, 0)
		} else {
			answer = append(answer, Sum(v[1:]))
		}
	}
	return answer
}
