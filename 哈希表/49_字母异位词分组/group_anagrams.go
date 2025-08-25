package leetcode49

// https://leetcode.cn/problems/group-anagrams

func groupAnagrams(strs []string) [][]string {
	cnts := make(map[[26]int][]string, len(strs))
	for _, str := range strs {
		cnt := [26]int{}
		for _, letter := range str {
			cnt[letter-'a']++
		}
		cnts[cnt] = append(cnts[cnt], str)
	}

	ans := make([][]string, 0, len(cnts))
	for _, v := range cnts {
		ans = append(ans, v)
	}
	return ans
}
