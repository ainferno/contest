package main

import (
	"fmt"
	"strconv"
)

func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseNum(a int) int {
	b := 0
	for ; a > 0; a = a / 10 {
		b = b*10 + a%10
	}
	return b
}

func LengthNum(a int) int {
	b := 1
	for ; a > 1; a = a / 10 {
		b *= 10
	}
	return b
}
func LengthStr(str string) int {
	b := 1
	for i := 0; i < len(str); i++ {
		b *= 10
	}
	return b
}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	l := len(str)
	str_left := str[:l/2]
	str_right := str[l/2+l%2:]
	str_left_reversed := Reverse(str_left)

	length := LengthStr(str_right)

	num_left, _ := strconv.Atoi(str_left)
	num_left_reversed, _ := strconv.Atoi(str_left_reversed)
	num_right, _ := strconv.Atoi(str_right)

	if num_left_reversed >= num_right {
		fmt.Println(num_left_reversed - num_right)
	} else {
		if l%2 == 0 {
			num_left += 1
		}
		num_left_reversed = ReverseNum(num_left)
		res := length + num_left_reversed - num_right
		fmt.Println(res)
	}
}
