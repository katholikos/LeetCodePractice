package SortSource

/*
   @author katholikos katholik@mail.ccsf.edu
   @version  08/05/2020 17:57
   @since Go1.13.5
*/

func HeapSortMax(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}else {
		depth := length/2 - 1
		for i := depth; i >= 0; i-- {
			top := i
			left := 2 * i + 1
			right := 2 * i + 2
			if left <= length-1 && left > top {
				top = left
			}else if right <= length-1 && right > top {
				top = right
			}else if top != i {
				arr[i],arr[top] = arr[top],arr[i]
			}
		}
		return arr
	}
}