package _83

func moveZeroes(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}

	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}
