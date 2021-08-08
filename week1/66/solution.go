package _6

func plusOne(digits []int) []int {
	if len(digits) < 1 {
		return nil
	}

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}

	return append([]int{1}, digits...)
}
