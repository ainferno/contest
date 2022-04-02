package main

import "fmt"

func copy(a []string) []string {
	var b []string
	b = append(b, a...)
	return b
}

func remove(str []string, i int) []string {
	return append(str[:i], str[i+1:]...)
}

func variants(a string, b []string, res *[]string, l int) {
	if len(a) == l {
		*res = append(*res, a)
		return
	}
	for i, elem := range b {
		b2 := copy(b)
		b2 = remove(b2, i)
		variants(a+string(elem), b2, res, l)
	}
}

func Variant(str string, l int) []string {
	var a []string
	for _, char := range str {
		a = append(a, string(char))
	}
	var res []string
	for i, elem := range a {
		a2 := copy(a)
		a2 = remove(a2, i)
		variants(string(elem), a2, &res, l)
	}
	return res
}

func main() {

	var pin string
	fmt.Scanf("%s", &pin)
	res := Variant(pin, len(pin))

	for i := 0; i < len(res); i++ {
		for j := i; j < len(res); j++ {
			if res[i] > res[j] {
				tmp := res[i]
				res[i] = res[j]
				res[j] = tmp
			}
		}
	}
	for _, elem := range res {
		fmt.Println(elem)
	}
}
