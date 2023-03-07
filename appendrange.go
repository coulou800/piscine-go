package piscine

func AppendRange(min, max int) []int {
	nums := []int{}
	if min < max {
		for i := min; i < max; i++ {
			nums = append(nums, i)
		}
		return nums
	} else {
		return nil
	}
}
