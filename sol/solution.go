package sol

func countSubstrings(s string) int {
	sLen := len(s)
	count := 0
	var countPalindromeFromMid = func(startLeft, startRight int) {
		left := startLeft
		right := startRight
		for left >= 0 && right < sLen && s[right] == s[left] {
			count++
			left--
			right++
		}
	}
	for idx := 0; idx < sLen; idx++ {
		countPalindromeFromMid(idx, idx)
		countPalindromeFromMid(idx, idx+1)
	}
	return count
}
