package leetcode

//TwoSum leetcode
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	var ret [2]int

	for index, value := range nums {
		if i, ok := m[target-value]; ok {
			ret[0] = i
			ret[1] = index
		} else {
			m[value] = index
		}
	}

	return ret[:]

}
