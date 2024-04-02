//Given an integer array nums, return true if any value appears
//at least twice in the array, and return false if every element is distinct.

package main

import "fmt"

func duplicate(nums []int) bool {
	myMap := make(map[int]bool)

	for _, num := range nums {
		if myMap[num] {
			return true
		}
		myMap[num] = true
	}

	return false
}
func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(duplicate(nums))
	nums = []int{1, 2, 3, 4, 5, 1}
	fmt.Println(duplicate(nums))

}
