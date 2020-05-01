package ArrayList

import (
	"errors"
	"fmt"
)

/*
   @author katholikos katholik@mail.ccsf.edu
   @version  01/05/2020 20:03
   @since Go1.13.5
*/


type List interface {
	Size() int //长度
	Get(index int) (interface{},error) //获取
	Set(index int, newVal interface{}) error //修改
	Insert(index int,newVal interface{}) error //插入
	Append(newVal interface{})//增加
	Clear() //清空
	Delete(index int) error //删除
	String() string //返回字符串
	checkIsFull() //检测内存
}

type ArrayList struct {
	DataStore[] interface{}
	ArraySize int
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.DataStore = make([]interface{},0,10)
	list.ArraySize = 0
	return list
}

func (list *ArrayList)Size() int {
	return list.ArraySize
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.ArraySize {
		return nil,errors.New("索引越界")
	}
	return list.DataStore[index],nil
}

func (list *ArrayList) Append(newVal interface{}) {
	list.DataStore = append(list.DataStore,newVal)
	list.ArraySize++
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.DataStore)
}

func (list *ArrayList) Set(index int, newVal interface{}) error{
	if index < 0 || index >= list.ArraySize {
		return errors.New("索引越界")
	}
	list.DataStore[index] = newVal
	return nil
}

func  (list *ArrayList) checkIsFull()  {
	if list.ArraySize == cap(list.DataStore) {
		newDataStore := make([]interface{},2*list.ArraySize,2*list.ArraySize)
		copy(newDataStore,list.DataStore)
		list.DataStore = newDataStore
	}
}

func (list *ArrayList) Insert(index int,newVal interface{}) error  {
	if index < 0 || index >= list.ArraySize {
		return errors.New("索引越界")
	}
	list.checkIsFull()
	list.DataStore = list.DataStore[:list.ArraySize+1]//内存 +1
	for i := list.ArraySize; i > index; i-- {
		list.DataStore[i] = list.DataStore[i - 1]
	}
	list.DataStore[index] = newVal
	list.ArraySize ++
	return nil

}

func (list *ArrayList) Clear()  {
	list.DataStore = make([]interface{},0,10)
	list.ArraySize = 0
}

func (list *ArrayList) Delete(index int) error  {
	if index < 0 || index >= list.ArraySize {
		return errors.New("索引越界")
	}
	list.DataStore = append(list.DataStore[:index],list.DataStore[index + 1:]...)
	list.ArraySize--
	return nil
}