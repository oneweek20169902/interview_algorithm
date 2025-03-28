package leecode

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
