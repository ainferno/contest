package main

import "fmt"

func sort_min(str []string) []string {
	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {
			if str[i] > str[j] {
				tmp := str[i]
				str[i] = str[j]
				str[j] = tmp
			}
		}
	}
	return str
}

func sort_max(str []string) []string {
	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {
			if str[i] < str[j] {
				tmp := str[i]
				str[i] = str[j]
				str[j] = tmp
			}
		}
	}
	return str
}

func assemble(str []string) string {
	var res string
	for _, elem := range str {
		res += elem
	}
	return res
}

func Variant_min(str string) string {
	var a []string
	for _, char := range str {
		a = append(a, string(char))
	}
	return assemble(sort_min(a))
}

func Variant_max(str string) string {
	var a []string
	for _, char := range str {
		a = append(a, string(char))
	}
	return assemble(sort_max(a))
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		var l int
		fmt.Scanf("%d", &l)

		var str1 string
		fmt.Scanf("%s", &str1)
		res1 := Variant_min(str1)

		var str2 string
		fmt.Scanf("%s", &str2)
		res2 := Variant_max(str2)

		if res1 < res2 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
