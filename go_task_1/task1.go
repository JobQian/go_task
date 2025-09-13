package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// 1---------singleNumber
func singleNumber_1(array []int) int {
	kvmap := make(map[int]int)
	for i := 0; i < len(array); i++ {
		x := array[i]
		_, exist := kvmap[x]
		if exist {
			kvmap[x]++
		} else {
			kvmap[x] = 1
		}
	}
	for key, value := range kvmap {
		if value == 1 {
			return key
		}
	}
	return 0
}

// 1---------singleNumber
func singleNumber_2(array []int) int {
	result := 0
	for _, num := range array {
		result ^= num
	}
	return result
}

// 2---------palindromeNumber-use string
func isPalindrome_1(num int) bool {
	if num < 0 {
		return false
	}
	runes := []rune(strconv.Itoa(num))
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	result, _ := strconv.Atoi(string(runes))
	fmt.Println(result)
	return num == result
}

// 2---------palindromeNumber- do not use string
func isPalindrome_2(num int) bool {
	if num < 0 {
		return false
	}
	tmp := num
	rev := 0
	for tmp != 0 {
		pop := tmp % 10
		tmp /= 10
		if rev > math.MaxInt32/10 || rev < math.MinInt32/10 {
			return false
		}
		rev = rev*10 + pop
	}
	return rev == num
}

// 3---------string isValid "()[]{}""
func isValid_1(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	if strings.Contains(s, "(") {
		if strings.Contains(s, ")") {
			inxl := strings.Index(s, "(")
			inxr := strings.Index(s, ")")
			if inxl > inxr {
				return false
			} else {
				if (inxr-inxl)%2 == 0 {
					return false
				}
			}
		} else {
			return false
		}
	}
	if strings.Contains(s, "[") {
		if strings.Contains(s, "]") {
			inxl := strings.Index(s, "[")
			inxr := strings.Index(s, "]")
			if inxl > inxr {
				return false
			} else {
				if (inxr-inxl)%2 == 0 {
					return false
				}
			}
		} else {
			return false
		}
	}
	if strings.Contains(s, "{") {
		if strings.Contains(s, "}") {
			inxl := strings.Index(s, "{")
			inxr := strings.Index(s, "}")
			if inxl > inxr {
				return false
			} else {
				if (inxr-inxl)%2 == 0 {
					return false
				}
			}
		} else {
			return false
		}
	}
	return true
}

// 3---------string isValid "()[]{}""
func isValid_2(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// 4---------longestCommonPrefix
func longestCommonPrefix_1(strs []string) string {
	if len(strs) != 0 {
		res := []rune{}
		runestrs := []rune(strs[0])
		for i := 0; i < len(runestrs); i++ {
			flag := true
			for j := 1; j < len(strs); j++ {
				runestrss := []rune(strs[j])
				if i >= len(runestrss) || runestrs[i] != runestrss[i] {
					flag = false
				}
			}
			if flag {
				res = append(res, runestrs[i])
			}
		}
		return string(res)
	}
	return ""
}

// 4---------longestCommonPrefix
func longestCommonPrefix_2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	first := []rune(strs[0])
	for i := 0; i < len(first); i++ {
		c := first[i]
		for j := 1; j < len(strs); j++ {
			runes := []rune(strs[j])
			if i >= len(runes) || runes[i] != c {
				return string(first[:i])
			}
		}
	}
	return string(first)
}

// 5---------plusOne
func plusOne(digits []int) []int {
	for i := 0; i < len(digits); i++ {
		res := digits[len(digits)-1-i] + 1
		if res == 10 {
			digits[len(digits)-1-i] = 0
		} else {
			digits[len(digits)-1-i] += 1
			return digits
		}
	}
	return append([]int{1}, digits...)
}

// 6---------removeDuplicates
func removeDuplicates(nums []int) int {
	for i := len(nums) - 2; i >= 0; i-- {
		x := nums[i+1]
		if nums[i] == x {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}

// 7---------merge
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		curr := intervals[i]
		if curr[0] <= last[1] {
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
			res[len(res)-1] = last
		} else {
			res = append(res, curr)
		}
	}
	return res
}

// 8---------twoSum
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	// array := []int{1, 2, 3, 4, 1, 2, 3, 4, 5}
	// fmt.Println("singleNumber_1:", singleNumber_1(array))
	// fmt.Println("singleNumber_2:", singleNumber_2(array))

	// num := 121
	// fmt.Println(isPalindrome_1(num))
	// fmt.Println(isPalindrome_2(num))

	// str := "([)]"
	// fmt.Println(isValid_1(str))
	// fmt.Println(isValid_2(str))

	// strs := []string{"flower", "flow", "flight"}
	// fmt.Println(longestCommonPrefix_1(strs))
	// fmt.Println(longestCommonPrefix_2(strs))

	// digits := []int{9, 9, 9}
	// fmt.Println(plusOne(digits))

	// digits := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// fmt.Println(removeDuplicates(digits))

	// digits := [][]int{{1, 4}, {0, 2}, {3, 5}}
	// fmt.Println(merge(digits))

	// digits := []int{3, 3}
	// target := 6
	// fmt.Println(twoSum(digits, target))
}
