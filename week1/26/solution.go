package _6

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var index int
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[index] {
			index++
			if i != index {
				nums[index] = nums[i]
			}
		}
	}

	return index + 1
}
