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

/*package main

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
*/
//solution-01.
package main

import "fmt"

func twoSum(num []int, target int) []int {
	numMap := make(map[int]int) // map will store the value
	//iterate through the array
	for i, num := range num {
		complement := target - num // the result is stored in numMap
		if j, ok := numMap[complement]; ok {
			// this'll check the current element is present in the numMap
			//as well as check make the comparison .
			// aslo produce the true/false result
			// according to this the next block is executed
			return []int{i, j}
		}
		//if numMap is enpty or not prest in num then the element is added to numMap
		numMap[num] = i

	}
	//if not found any any result return empty
	//you can give specific msg here
	return []int{}
}
func main() {
	fmt.Println("another way to solve this problem")
	num := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(num, target))
}
