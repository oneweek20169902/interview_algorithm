package leecode

import "sort"

// 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	set := map[byte]bool{}
	right, left, maxLen := 0, 0, 0
	for right < len(s) {
		if !set[s[right]] {
			set[s[right]] = true
			right++
			if len(set) > maxLen {
				maxLen = len(set)
			}
		} else {
			delete(set, s[left])
			left++
		}
	}
	return maxLen
}

// 三数之和
func threeSum(nums []int) (result [][]int) {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[r] + nums[l]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[r], nums[l]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					l++
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return
}
