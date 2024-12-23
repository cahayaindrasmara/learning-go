package main

import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {
	var result []int
	result = make([]int, 0)

	// if len(date1) == 0 || len(date2) == 0 {
	// 	result = make([]int, 0)
	// }

	// for i := 1; i < 31; i++ {
	// 	isPerson1 := false
	// 	isPerson2 := false

	// 	for _, date := range date1 {
	// 		if date == i {
	// 			isPerson1 = true
	// 		}
	// 	}

	// 	for _, date := range date2 {
	// 		if date == i {
	// 			isPerson2 = true
	// 		}
	// 	}

	// 	if isPerson1 && isPerson2 {
	// 		result = append(result, i)
	// 	}
	// }
	// return result // TODO: replace this

	for _, datePerson1 := range date1 {
		for _, datePerson2 := range date2 {
			if datePerson1 == datePerson2 {
				result = append(result, datePerson1)
				break
			}
		}
	}

	return result // TODO: replace this
}

func main() {
	fmt.Println(SchedulableDays([]int{1, 2, 3, 4}, []int{3, 4, 5}))
}
