package main

import (
	"fmt"
)

type LessAdder interface {
	Less(b integer) bool
	Add(b integer)
}

type LessAdder2 interface {
	Add(b integer)
	Less(b integer) bool
}

type LessAdder3 interface {
	Add(b integer)
	Sum()
}

type integer int

func (a integer) Less(b integer) bool {
	return a < b
}

func (a *integer) Add(b integer) {
	*a += b
}

func main() {
	var a integer = 1
	var b LessAdder = &a
	var z interface{} = &a

	fmt.Println(b.Less(3))
	if c, ok := b.(LessAdder2); ok { //	b接口是否实现了接口LessAdder2，实现了
		fmt.Println(ok)
		fmt.Println(c)
	}

	if c, ok := b.(LessAdder3); ok { // b接口是否实现了接口LessAdder3，未实现
		fmt.Println(ok)
		fmt.Println(c)
	}

	if c, ok := b.(*integer); ok { // 接口指向的对象是否是integer类型
		fmt.Println(ok)
		fmt.Println(c)
	}

	if c, ok := z.(int); ok { // 类型查询，任何类型都实现了空接口
		fmt.Println(ok)
		fmt.Println(c)
	}

}
