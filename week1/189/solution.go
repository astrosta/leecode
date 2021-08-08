package _89

func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}

	k = k % len(nums)

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, beg, end int) {
	for beg < end {
		nums[beg], nums[end] = nums[end], nums[beg]
		beg++
		end--
	}
}
