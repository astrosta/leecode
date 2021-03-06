package _8

func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}

	if n > 0 {
		copy(nums1[:m+n], nums2[:n])
	}

	return
}
