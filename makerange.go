package piscine

func MakeRange(min, max int) []int {
	if min < max {
		nums := make([]int, max-min)
		for i := 0; i < max-min; i++ {
			nums[i] = min + i
		}
		return nums
	}
	return nil
}
