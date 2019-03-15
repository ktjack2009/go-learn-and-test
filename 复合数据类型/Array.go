package main

import "fmt"

const (
	USD int = iota
	EUR
	GBP
	RMB
)

func ArrayDemo1() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])

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
	ArrayDemo1()
	// ArrayDemo2()
}
