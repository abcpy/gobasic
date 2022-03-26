package main

import "fmt"

// lastOccurred[x] 不存在 或者 < start ->无需操作
// lastOccurred[x] >= start -> start
// 更新lastOccurred[x], 更新maxLength

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxlength := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxlength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))

}
