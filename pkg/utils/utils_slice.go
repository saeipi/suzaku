package utils

func Max(vals []int) int {
	var max int
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func Min(vals []int) int {
	var min int
	for _, val := range vals {
		if min == 0 || val <= min {
			min = val
		}
	}
	return min
}
