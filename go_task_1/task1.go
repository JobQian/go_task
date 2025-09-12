package main

import (
	"fmt"
	"math"
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

func main() {
	array := []int{1, 2, 3, 4, 1, 2, 3, 4, 5}
	fmt.Println("singleNumber_1:", singleNumber_1(array))
	fmt.Println("singleNumber_2:", singleNumber_2(array))

	num := 121
	fmt.Println(isPalindrome_1(num))
	fmt.Println(isPalindrome_2(num))

	str := "([)]"
	fmt.Println(isValid_1(str))
	fmt.Println(isValid_2(str))
}
