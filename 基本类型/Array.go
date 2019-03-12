package main

import "fmt"

func ArrayDemo1() {
	array := [5]int{10, 20, 30, 40, 50}
	array2 := array
	array2[0] = 50 // 只是值赋值，并不会改变array中的值，与python的list的区别
	for _, v := range array {
		fmt.Println(v)
	}
	for _, v := range array2 {
		fmt.Println(v)
	}
}

func ArrayDemo2() {
	var array3 [5]*int
	for i := range array3 {
		array3[i] = new(int)
	}
	*array3[0] = 1
	*array3[1] = 2
	*array3[2] = 3
	*array3[3] = 4
	*array3[4] = 5
	array4 := array3
	for _, v := range array4 {
		fmt.Println(*v)
	}
}

func main() {
	ArrayDemo2()
}
