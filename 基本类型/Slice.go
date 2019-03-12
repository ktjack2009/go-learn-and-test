package main

import (
	"fmt"
	"strings"
)

func SliceDemo1() {
	source := make([]string, 3, 5) // 访问长度3个元素，底层数组拥有5个元素，不允许创建容量小于长度的切片
	source = []string{"red", "green", "blue", "yellow", "black"}
	slice := source[2:3:5] // 1个元素，容量为5-2=3
	for _, v := range slice {
		fmt.Println(v)
	}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func SliceDemo2() {
	source := make([]string, 10)
	source = []string{"123", "456", "789"}
	a := strings.Join(source, ",") // 切片转字符串
	m := strings.Split(a, ",")
	fmt.Println(a)
	fmt.Println(m)
}

func SliceDemo3() {
	a := []int{1, 2, 3, 4, 5}
	a = append(a, 0)
	fmt.Println(a)
	copy(a[3:], a[2:])
	fmt.Println(a)
	a[2] = 10
	fmt.Print(a)
}

func SliceDemo4() {
	// 原地删除，公用底层数组
	s := []string{"a", "b", " ", "d", "e", " ", "f"}
	b := s[:0]
	for _, x := range s {
		if x != " " {
			b = append(b, x)
		}
	}
	fmt.Print(b)
}

func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

func main() {
	// SliceDemo1()
	// SliceDemo2()
	// SliceDemo3()
	// SliceDemo4()
	fmt.Print(Sum(1, 2, 3, 4, 5, 6, 7))
}
