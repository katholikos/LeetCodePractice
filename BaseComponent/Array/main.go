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
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Println(list.DataStore)
	fmt.Println(list.ArraySize)
	for i := 0; i<15; i++{
		list.Insert(1,"插入")
		fmt.Println(list.DataStore)
	}
}
