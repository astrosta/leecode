package _1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	var tmp int
	for i, value := range nums {
		tmp = target - value
		if index, ok := m[tmp]; ok {
			return []int{index, i}
		}
		m[value] = i
	}
	return nil
}
