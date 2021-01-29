package leetcode

/*
来源：leetcode 13
难度：简单
答题时间：2021.01.29
思路：对输入的字符串转化map，从右往左遍历，当前的字符大于等于上一个字符，则累加； 否则累减
*/

func RomanToInt(s string) int {
	var _map = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	pre,r := 0,0
	for i := len(s) - 1;i >= 0;i--{
		cur := _map[s[i]]
		if cur >= pre{
			r += cur
		}else{
			r -= cur
		}
		pre = cur
	}
	return r
}