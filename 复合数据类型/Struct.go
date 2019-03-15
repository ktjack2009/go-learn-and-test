/*
	一个命名为S的结构体类型将不能再包含S类型的成员：因为一个聚合的值不能包含它自身。
	但是S类型的结构体可以包含*S指针类型的成员，这可以让我们创建递归的数据结构，比如链表和树结构等
*/

package main

import "fmt"

// 递归数据结构
type tree struct {
	value int
	left  *tree
	right *tree
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

func main() {
	var c Circle
	c.X = 1
	c.Y = 2
	c.Radius = 3
	fmt.Println(c)
}
