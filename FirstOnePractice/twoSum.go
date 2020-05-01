/*
@Time : 4/28/2020 1:52 AM
@Author : ka
@File : twoSum
@Software: GoLand
*/
package main

import (
	"fmt"
)


func twoSumWithSlice(nums []int, target int) []int {
	for firstkey, valueOne := range nums {
		//O(n^2) || O(1)
		for secondKey:= firstkey + 1; firstkey < len(nums); secondKey++ {
			if valueOne + nums[secondKey] == target {
				return []int{firstkey,secondKey}
			}

		}
	}
	return nil
}

func twoSumWithMap(nums []int, target int) []int {
	var result []int
	resultMap := make(map[int]int)
	for k, v := range nums {
		if value, isValueInMap := resultMap[target-v];isValueInMap {
			result = append(result,value,k)
			return result
		}else {
			resultMap[v] = k
		}
	}
	return nil
}

func main()  {
	num := []int {2,7,11,15}
	target := 9
	fmt.Println(twoSumWithSlice(num,target))
	fmt.Println(twoSumWithMap(num,target))

	a := 36
	ptr := &a
	pptr := &ptr
	fmt.Println(&a)
	fmt.Println(pptr)// ptr地址
	fmt.Println(*pptr)//a地址 *pptr == ptr
	fmt.Println(**pptr)//a **pptr == *ptr == a
}
