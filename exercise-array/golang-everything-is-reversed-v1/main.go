package main

import (
	"fmt"
	"strconv"
)

func ReverseArray(arr [5]int) [5]int {
	var result [5]int
	for i := 0; i < len(arr); i++ {
		result[i] = arr[len(arr)-1-i]
	}

	return result // TODO: replace this
}

func ReverseDigit(number int) int {
	s := strconv.Itoa(number)
	rev := ""

	for i := len(s) - 1; i >= 0; i-- {
		rev = rev + string(s[i])
	}

	result, _ := strconv.Atoi(rev)

	// for _, v := range s {
	// 	rev = string(v) + rev
	// }

	// result, _ := strconv.Atoi(rev)

	return result
}

func ReverseData(arr [5]int) [5]int {
	a := ReverseArray(arr)
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		a[i] = ReverseDigit(a[i])
	}
	fmt.Println(a)
	return a
}

func main() {
	fmt.Println(ReverseData([5]int{123, 456, 11, 1, 2}))
	fmt.Println(ReverseDigit(1234))
}
