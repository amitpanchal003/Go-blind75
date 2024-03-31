//leetcode problem 01:-

/*
	Given an array of integers nums and an integer target,
	return indices of the two numbers such that they add up to target.

	You may assume that each input would have exactly one solution,
	and you may not use the same element twice.

	nums[2,7,11,15]
	target=9

*/

//solution-01.

package main

import "fmt"

//lets create a function for twoSum
func twoSum(num []int, target int) []int {
	//first loop through the values

	for i := 0; i < len(num); i++ {
		//inatialise inner loop

		for j := i + 1; j < len(num); j++ {
			//check the condition
			if num[i]+num[j] == target {
				return []int{i, j}
			}
		}
	}
	//if not found any match return
	return []int{}
}

func main() {
	fmt.Println("the result is: ")
	num := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(num, target))
}
