# golang_palindromic_substrings

Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.

## Examples

**Example 1:**

```
Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".

```

**Example 2:**

```
Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

```

**Constraints:**

- `1 <= s.length <= 1000`
- `s` consists of lowercase English letters.

## 解析

類似 [5. Longest Palindromic Substring](https://www.notion.so/5-Longest-Palindromic-Substring-53a75a24c3c348f1afb5fc9e4cf17718) 但這題是要計算所有 palindrome substring 個數

給定一個字串 s, 要找尋所有是 Palindrome substring 的個數

所以只要一樣使用從中間開始往兩邊去伸展

當發現是 palindrome substring 時就把 count 加一

![](https://i.imgur.com/V15tsZn.png)

## 程式碼
```go
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

```
## 困難點

1. 要看出 palindrome substring 的特性

## Solve Point

- [x]  透過每個字元從左與右伸展當發現是 palindrome 時 就累加
- [x]  要對 odd 與 even 個別做考慮