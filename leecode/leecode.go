package leecode

// 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	left, right, maxlength := 0, 0, 0
	set := make(map[byte]bool)
	for right < len(s) {
		if !set[s[right]] {
			set[s[right]] = true
			right++
			if len(set) > maxlength {
				maxlength = len(set)
			}
		} else {
			delete(set, s[left])
			left++
		}
	}
	return maxlength
}
