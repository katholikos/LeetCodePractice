package main

import (
	"./SortSource"
	"fmt"
)

/*
   @author katholikos katholik@mail.ccsf.edu
   @version  07/05/2020 23:04
   @since Go1.13.5
*/

func main() {
	arr := []int{6,2,2,10,45}
	fmt.Println(arr)
	fmt.Println(SortSource.QuickSort(arr))
}