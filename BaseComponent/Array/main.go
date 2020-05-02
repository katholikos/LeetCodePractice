package main

import (
	"./ArrayList"
	"fmt"
)
/*
   @author katholikos katholik@mail.ccsf.edu
   @version  01/05/2020 20:03
   @since Go1.13.5
*/

func main() {
	list := ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	for it := list.Iterator(); it.HasNext(); {
		item,_ :=it.Next()
		if item == "a1" {
			it.Remove()
			fmt.Println("成功删除")
		}
		fmt.Println(item)

	}
	fmt.Println(list)
}
