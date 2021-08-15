/*
	字母异位词分组
	解题思路：将字符串中的字母排序或者用数组记录字母作为哈希表的key, 哈希表的value为原字符串切片，最后输出结果
*/

package _9

func groupAnagrams(strs []string) [][]string {
	if strs == nil {
		return nil
	}

	m := make(map[[26]int][]string)
	for i := 0; i < len(strs); i++ {
		//用容量为26的数组记录字符串的数组，数组相同的字符串为字母异位词
		cnt := [26]int{}
		for _, val := range strs[i] {
			cnt[val-'a']++
		}
		m[cnt] = append(m[cnt], strs[i])
	}

	var result [][]string
	for _, value := range m {
		result = append(result, value)
	}

	return result
}
