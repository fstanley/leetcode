package leetcode

func twoSum(nums []int, target int) []int {
	for k0, v0 := range nums {
		for k1, v1 := range nums[k0+1:] {
			if v0+v1 == target {
				return []int{k0, k1 + (k0 + 1)}
			}
		}
	}
	return []int{0, 0}
}
