package leetcode5

// https://leetcode.cn/problems/longest-palindromic-substring
//
// 题解：
// 对于一个子串而言，如果它是回文串且长度大于 2，那么将它首位的两个字母去掉后，它仍然是回文串
// 对于长度等于 1 的子串，它是天然的回文串
// 对于长度等于 2 的子串，只要它的两个字母相同，它就是一个回文串
// 根据以上规则可以构造动态规划的状态转移方程：
// - P(i, i)     == true
// - P(i, i + 1) == S[i] == S[i + 1]
// - P(i, j)     == P(i + 1, j - 1) ^ (S[i] == S[i + 1])
//
// 其中 P(i, j) 表示字符串 S 第 i 到 j 字母组成的子串是否为回文串

func longestPalindrome(s string) string {
	n := len(s)

	// dp[i][j] 表示 s[i..j] 是否是回文串
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(s))
	}
	// 构造 dp 数组
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			longestPalindromeHelper(s, dp, i, j)
		}
	}

	// 遍历 dp 数组获取最长回文子串长度
	imax, jmax, size := 0, 0, 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if dp[i][j] == 1 && j-i+1 > size {
				size = j - i + 1
				imax, jmax = i, j
			}
		}
	}

	return s[imax : jmax+1]
}

func longestPalindromeHelper(s string, dp [][]int, i, j int) int {
	switch {
	case dp[i][j] != 0:
		break
	case s[i] != s[j]:
		dp[i][j] = -1
	case j-i < 3:
		dp[i][j] = 1
	default:
		dp[i][j] = longestPalindromeHelper(s, dp, i+1, j-1)
	}
	return dp[i][j]
}
