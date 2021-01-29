package leetcode

/*
来源：leetcode 9
难度：简单
答题时间：2021.01.29
思路：对数字进行反转并与原值比较大小
*/
func IsPalindrome(x int) bool {
	if x < 0{
		return false
	}
	origin := x
	redirect := 0
	for x != 0{
		redirect = redirect * 10 + x % 10
		x /= 10
	}
	return origin == redirect
}