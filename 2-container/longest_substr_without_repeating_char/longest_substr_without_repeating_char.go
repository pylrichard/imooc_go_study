package main

import "fmt"

func getMaxLenOfLongestSubstrWithoutRepeatingChar(str string) int {
	// map[byte]int不能处理中文
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	// []byte(str)不能处理中文
	for i, char := range []rune(str) {
		lastIndex, exist := lastOccurred[char]
		// 该字符出现过而且在start后(遇到重复字符)时
		// 需要把start移到该字符后一位使得字符不重复(消除重复)
		if exist && lastIndex >= start {
			start = lastIndex + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[char] = i
	}

	return maxLength
}

func main() {
	fmt.Println(
		getMaxLenOfLongestSubstrWithoutRepeatingChar("abcabcbb"))
	fmt.Println(
		getMaxLenOfLongestSubstrWithoutRepeatingChar("bbbbb"))
	fmt.Println(
		getMaxLenOfLongestSubstrWithoutRepeatingChar(""))
	fmt.Println(
		getMaxLenOfLongestSubstrWithoutRepeatingChar("在慕课网学习课程"))
	fmt.Println(
		getMaxLenOfLongestSubstrWithoutRepeatingChar("一二三三二一"))
}
