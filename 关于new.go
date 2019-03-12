/*
	1. make只能用来分配及初始化类型为slice、map、chan的数据，new可以分配任意类型的数据
	2. make返回引用，new返回指针
	3. new分配的空间被清零，make分配后会进行初始化
*/

package main

import "fmt"

func main() {
	var a int = 1
	b := new(int)
	b = &a
	fmt.Println(b)
}
