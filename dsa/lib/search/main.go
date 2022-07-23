package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 5, 7, 9, 12, 14}
	//idx, _ := BinarySearch(arr, 20, 0, len(arr)-1)
	//fmt.Println(idx)

	val, err := QuickSelect(arr, -7)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(val)

}
