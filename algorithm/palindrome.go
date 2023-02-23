package main

import (
	"fmt"
)

//回文，英文palindrome，指一个顺着读和反过来读都一样的字符串，比如madam、我爱我，这样的短句在智力性、趣味性和艺术性上都颇有特色，中国历史上还有很多有趣的回文诗。
//那么，我们的第一个问题就是：判断一个字串是否是回文？

//判断是否为回文
func palindrome(s string) {
	first, last := 0, len(s)-1
	for i := 0; i <= len(s)/2; i++ {
		if s[first] != s[last] {
			fmt.Println("false")
			return
		}
		first++
		last--
	}
	fmt.Println("true")
}

func main() {
	s := "abcdedcba"
	palindrome(s)

	s = "abcdefdcba"
	palindrome(s)

	s = "abcdefedcba"
	palindrome(s)
}