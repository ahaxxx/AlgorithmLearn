package Algorithm

func lengthOfLongestSubstring(s string) int{
	m := map[string]int{}
	k := -1
	lens := len(s)
	jj := 0
	for i := 0;i < lens;i++{
		if i != 0{
			delete(m,string(s[i-1]))
		}
		for k+1 < lens && m[string(s[k+1])] == 0{
			m[string(s[k+1])]++
			k++
		}
		jj = max(jj,k - i + 1)
	}
	return jj
}

func max(a int,b int ) (c int) {
	if a>b {
		c = a
	} else {
		c = b
	}
	return
}