package main

import "fmt"

func MapDemo1() {
	dict := make(map[string]int)
	dict["abc"] = 123
	dict["efg"] = 456
	fmt.Println(dict["abc"])
	//dict["abcd"] = 0		// 赋值为零值，exist会为true
	_, exist := dict["abcd"]
	fmt.Println(exist)               // 判断是否存在
	delete(dict, "abcd")             // 即使key不存在也不会报错
	for index, value := range dict { // map遍历是无序的
		fmt.Println(index, value)
	}
}

func MapDemo2() {
	var a interface{}
	var b uint64 = 10
	a = b
	s := make(map[string]uint64)
	s["1"] = a.(uint64)
	fmt.Println(s)
}

func MapDemo3(a map[string]string) {
	a["a"] = "a"
}

func main() {
	// MapDemo1()
	// MapDemo2()
	b := map[string]string{
		"b": "b",
	}
	MapDemo3(b)
	fmt.Println(b)
}
