package arrayandslices

func Sum(num []int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}
