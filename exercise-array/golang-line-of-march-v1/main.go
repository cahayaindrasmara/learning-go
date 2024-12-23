package main

import "fmt"

func Sortheight(height []int) []int {

	// var result []int

	for i := 0; i < len(height)-1; i++ {
		fmt.Println(height[i])
		for j := 0; j < len(height)-i-1; j++ {
			fmt.Println("----->", height[i])
			if height[j] > height[j+1] {
				height[j], height[j+1] = height[j+1], height[j]
			}
		}
	}
	return height // TODO: replace this
}

func main() {
	fmt.Println(Sortheight([]int{172, 170, 150, 165, 144, 155, 159}))
}
